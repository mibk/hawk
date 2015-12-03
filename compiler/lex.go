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

const eof = 0

var lexlineno = 1

func init() {
	yyErrorVerbose = true
}

func (l *yyLex) Lex(yylval *yySymType) int {
	for {
		c := l.next()
		if unicode.IsDigit(c) {
			return l.lexNum(yylval)
		} else if unicode.IsLetter(c) {
			return l.lexIdent(yylval)
		}
		switch c {
		case eof:
			return eof
		case ';', '{', '}', ',', '(', ')', '$':
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
			}
		case '-':
			if l.peek() == '-' {
				l.next()
				return DEC
			}
		case '*':
		case '/':
			if l.peek() == '/' {
				l.next()
				for l.next() != '\n' {
				}
				l.backup()
				continue // ignore oneline comment
			} else if l.peek() == '*' {
				l.next()
				for !(l.next() == '*' && l.peek() == '/') {
				}
				l.next()
				continue // ignore block comment
			}
		case '%':
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
		return eof
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
