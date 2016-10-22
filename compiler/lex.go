package compiler

//go:generate -command yacc go tool yacc
//go:generate yacc -o hawk.go hawk.y

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

type yyLex struct {
	err    error
	reader *bufio.Reader
	last   rune
	peeked rune
	buf    *bytes.Buffer
}

const eof = -1

var (
	lexlineno int
	nlsemi    bool
)

func init() {
	yyErrorVerbose = true
}

func (l *yyLex) Lex(yylval *yySymType) (tok int) {
	defer func() {
		switch tok {
		case IDENT, NUM, STRING, BREAK, CONTINUE, INC, DEC, ')', '}':
			nlsemi = true
		default:
			nlsemi = false
		}
	}()
	for {
		if nlsemi && l.peek() == '\n' {
			nlsemi = false
			return ';'
		}
		r := l.next()
		if unicode.IsDigit(r) {
			return l.lexNum(yylval)
		} else if r >= utf8.RuneSelf || isLetter(r) {
			return l.lexIdent(yylval)
		}
		switch r {
		case eof:
			if nlsemi {
				// Treat EOF as \n.
				nlsemi = false
				return ';'
			}
			return 0
		case '_':
			return l.lexIdent(yylval)
		case ';', '{', '}', ',', '(', ')', '$', '|':
		case '?', ':':
		case '=':
			if l.accept('=') {
				return EQ
			}
		case '!':
			if l.accept('=') {
				return NE
			}
		case '<':
			if l.accept('=') {
				return LE
			}
		case '>':
			if l.accept('=') {
				return GE
			}
		case '+':
			if l.accept('+') {
				return INC
			} else if l.accept('=') {
				return ADDEQ
			}
		case '-':
			if l.accept('-') {
				return DEC
			} else if l.accept('=') {
				return SUBEQ
			}
		case '*':
			if l.accept('=') {
				return MULEQ
			}
		case '/':
			switch l.next() {
			case '=':
				return DIVEQ
			case '/':
				for {
					r := l.next()
					if r == '\n' || r == eof {
						break
					}
				}
				l.backup()
				continue // ignore oneline comment
			case '*':
				nl := false
				for {
					r := l.next()
					if r == eof {
						l.Error("eof in block comment")
						return eof
					} else if r == '*' && l.accept('/') {
						break
					} else if nl == false && r == '\n' {
						lexlineno--
						nl = true
					}
				}
				if nl {
					l.peeked = '\n'
				}
				continue // ignore block comment
			default:
				l.backup()
			}
		case '%':
			if l.accept('=') {
				return MODEQ
			}
		case '"', '\'':
			return l.lexString(r, yylval)
		case ' ', '\t', '\n', '\r':
			continue // ignore whitespace
		default:
			if r == '&' && l.accept('&') {
				return ANDAND
			} else if r == '|' && l.accept('|') {
				return OROR
			}
			l.Errorf("unrecognized character %q", r)
		}
		return int(r)
	}
}

func (l *yyLex) lexNum(yylval *yySymType) int {
	l.buf.Reset()
	l.buf.WriteRune(l.last)
	for unicode.IsDigit(l.next()) {
		l.buf.WriteRune(l.last)
	}
	l.backup()
	yylval.num, _ = strconv.Atoi(l.buf.String())
	return NUM
}

func (l *yyLex) lexIdent(yylval *yySymType) int {
	l.buf.Reset()
	l.buf.WriteRune(l.last)
	for {
		r := l.next()
		if r < utf8.RuneSelf && !isLetter(r) && !isDigit(r) && r != '_' {
			break
		}
		l.buf.WriteRune(l.last)
	}
	l.backup()
	name := l.buf.String()
	if tok, ok := isSymbol(name); ok {
		return tok
	}
	yylval.sym = name
	switch name {
	case "print", "printf":
		return PRINT
	}
	return IDENT
}

func isSymbol(name string) (int, bool) {
	for _, sym := range symbols {
		if sym.name == name {
			return sym.tok, true
		}
	}
	return 0, false
}

var symbols = []struct {
	name string
	tok  int
}{
	{"BEGIN", BEGIN},
	{"END", END},
	{"if", IF},
	{"else", ELSE},
	{"for", FOR},
	{"break", BREAK},
	{"continue", CONTINUE},
	{"func", FUNC},
	{"return", RETURN},
}

func (l *yyLex) lexString(quote rune, yylval *yySymType) int {
	l.buf.Reset()
loop:
	for {
		r := l.next()
		switch r {
		case eof:
			l.Error("eof in string literal")
			return eof
		case '\n':
			l.Error("newline in string literal")
		case '\\':
			switch l.next() {
			case 'a': // alert or bell
				r = '\a'
			case 'b': // backspace
				r = '\b'
			case 'f': // form feed
				r = '\f'
			case 'n': // line feed or newline
				r = '\n'
			case 'r': // carriage return
				r = '\r'
			case 't': // horizontal tab
				r = '\t'
			case 'v': // vertical tab
				r = '\v'
			case '\\': // backslash
				r = '\\'
			case quote: // " or '
				r = quote
			default:
				l.Errorf("unknown escape character \\%c", l.last)
			}
		case quote:
			break loop
		}
		l.buf.WriteRune(r)
	}
	yylval.sym = l.buf.String()
	return STRING
}

func (l *yyLex) next() (r rune) {
	defer func() {
		if r == '\n' {
			lexlineno++
		}
	}()
	if l.peeked != 0 {
		r := l.peeked
		l.peeked = 0
		return r
	}
	r, _, err := l.reader.ReadRune()
	if err != nil {
		r = eof
	}
	l.last = r
	return r
}

func (l *yyLex) accept(r rune) bool {
	if l.next() == r {
		return true
	}
	l.backup()
	return false
}

func (l *yyLex) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *yyLex) backup() {
	l.peeked = l.last
	if l.last == '\n' {
		lexlineno--
	}
}

func (l *yyLex) Error(s string) {
	if l.err == nil {
		l.err = fmt.Errorf("%d: %s", lexlineno, s)
	}
}

func (l *yyLex) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

func isLetter(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
