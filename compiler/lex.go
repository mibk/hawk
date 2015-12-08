package compiler

//go:generate -command yacc go tool yacc
//go:generate yacc -o hawk.go hawk.y

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"unicode"
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
	lexlineno = 1
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
		c := l.next()
		if unicode.IsDigit(c) {
			return l.lexNum(yylval)
		} else if unicode.IsLetter(c) {
			return l.lexIdent(yylval)
		}
		switch c {
		case eof:
			if nlsemi {
				// Treat EOF as \n.
				nlsemi = false
				return ';'
			}
			return 0
		case ';', '{', '}', ',', '(', ')', '$', '|':
		case '?', ':':
		case '=':
			if l.peek() == '=' {
				l.next()
				return EQ
			}
		case '!':
			if l.peek() == '=' {
				l.next()
				return NE
			}
		case '<':
			if l.peek() == '=' {
				l.next()
				return LE
			}
		case '>':
			if l.peek() == '=' {
				l.next()
				return GE
			}
		case '+':
			if l.peek() == '+' {
				l.next()
				return INC
			} else if l.peek() == '=' {
				l.next()
				return ADDEQ
			}
		case '-':
			if l.peek() == '-' {
				l.next()
				return DEC
			} else if l.peek() == '=' {
				l.next()
				return SUBEQ
			}
		case '*':
			if l.peek() == '=' {
				l.next()
				return MULEQ
			}
		case '/':
			switch l.peek() {
			case '=':
				l.next()
				return DIVEQ
			case '/':
				l.next()
				for l.next() != '\n' {
				}
				l.backup()
				continue // ignore oneline comment
			case '*':
				nl := false
				l.next()
				for {
					r := l.next()
					if r == '*' && l.peek() == '/' {
						break
					} else if nl == false && r == '\n' {
						lexlineno--
						nl = true
					}
				}
				l.next()
				if nl {
					l.peeked = '\n'
				}
				continue // ignore block comment
			}
		case '%':
			if l.peek() == '=' {
				l.next()
				return MODEQ
			}
		case '"':
			return l.lexString(yylval)
		case ' ', '\t', '\n', '\r':
			continue // ignore whitespace
		default:
			if c == '&' && l.peek() == '&' {
				l.next()
				return ANDAND
			} else if c == '|' && l.peek() == '|' {
				l.next()
				return OROR
			}
			l.Error(fmt.Sprintf("unrecognized character %q", c))
		}
		return int(c)
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
	for unicode.IsLetter(l.next()) {
		l.buf.WriteRune(l.last)
	}
	l.backup()
	sym := l.buf.String()
	switch sym {
	case "BEGIN":
		return BEGIN
	case "END":
		return END
	case "if":
		return IF
	case "else":
		return ELSE
	case "for":
		return FOR
	case "break":
		return BREAK
	case "continue":
		return CONTINUE
	case "func":
		return FUNC
	case "return":
		return RETURN
	}
	yylval.sym = sym
	switch sym {
	case "print":
		return PRINT
	}
	return IDENT
}

func (l *yyLex) lexString(yylval *yySymType) int {
	l.buf.Reset()
	for l.next() != '"' {
		l.buf.WriteRune(l.last)
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
		l.err = fmt.Errorf("%d: %s\n", lexlineno, s)
	}
}
