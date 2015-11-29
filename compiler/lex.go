package compiler

//go:generate -command yacc go tool yacc
//go:generate yacc -o hawk.go hawk.y

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

type yyLex struct {
	src   []byte
	start int
	pos   int
	width int
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
			l.emit()
			return eof
		case ';', '{', '}', ',', '(', ')', '$':
		case '=':
			if l.peek() == '=' {
				l.next()
				l.emit()
				return EQ
			}
		case '!':
			if l.peek() == '=' {
				l.next()
				l.emit()
				return NE
			}
		case '<':
			if l.peek() == '=' {
				l.next()
				l.emit()
				return LE
			}
			l.emit()
			return LT
		case '>':
			if l.peek() == '=' {
				l.next()
				l.emit()
				return GE
			}
			l.emit()
			return GT
		case '+':
			if l.peek() == '+' {
				l.next()
				l.emit()
				return INC
			}
		case '-':
			if l.peek() == '-' {
				l.next()
				l.emit()
				return DEC
			}
		case '*':
		case '/':
			if l.peek() == '/' {
				l.next()
				for l.next() != '\n' {
				}
				l.backup()
				l.emit() // ignore oneline comment
				continue
			} else if l.peek() == '*' {
				l.next()
				for !(l.next() == '*' && l.peek() == '/') {
				}
				l.next()
				l.emit()
				continue // ignore block comment
			}
		case '%':
		case '"':
			return l.lexString(yylval)
		case ' ', '\t', '\n', '\r':
			l.emit() // ignore whitespace
			continue
		default:
			if c == '&' && l.peek() == '&' {
				l.next()
				l.emit()
				return LAND
			} else if c == '|' && l.peek() == '|' {
				l.next()
				l.emit()
				return LOR
			}
			l.Error(fmt.Sprintf("unrecognized character %q", c))
		}

		l.emit()
		return int(c)
	}
}

func (l *yyLex) lexNum(yylval *yySymType) int {
	for unicode.IsDigit(l.next()) {
	}
	l.backup()
	num := l.emit()
	yylval.num, _ = strconv.Atoi(string(num))
	return NUM
}

func (l *yyLex) lexIdent(yylval *yySymType) int {
	for unicode.IsLetter(l.next()) {
	}
	l.backup()
	sym := string(l.emit())
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
	for l.next() != '"' {
	}
	s := l.emit()
	yylval.sym = string(s[1 : len(s)-1]) // trim ""
	return STRING
}

func (l *yyLex) next() rune {
	if len(l.src) == l.start {
		return eof
	}
	r, w := utf8.DecodeRune(l.src[l.pos:])
	l.width = w
	l.pos += l.width
	if r == '\n' {
		lexlineno++
	}
	return r
}

func (l *yyLex) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *yyLex) backup() {
	l.pos -= l.width
	r, _ := utf8.DecodeRune(l.src[l.pos : l.pos+l.width])
	if r == '\n' {
		lexlineno--
	}
}

func (l *yyLex) emit() []byte {
	tok := l.src[l.start:l.pos]
	l.start = l.pos
	return tok
}

func (l *yyLex) Error(s string) {
	fmt.Printf("hawk: line:%d: %s\n", lexlineno, s)
	os.Exit(1)
}
