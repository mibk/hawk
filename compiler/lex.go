package compiler

//go:generate -command yacc go tool yacc
//go:generate yacc -o hawk.go hawk.y

import (
	"log"
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

func (l *yyLex) Lex(yylval *yySymType) int {
	for {
		c := l.next()
		if unicode.IsDigit(c) {
			return l.lexNum(yylval)
		}
		switch c {
		case eof:
			l.emit()
			return eof
		case '$':
			l.emit()
			return int(c)
		case '=':
			if l.peek() == '=' {
				l.next()
				return EQ
			}
			return int(c)
		case '!':
			if l.peek() == '=' {
				l.next()
				return NE
			}
			return int(c)
		case '<':
			if l.peek() == '=' {
				l.next()
				return LE
			}
			return LT
		case '>':
			if l.peek() == '=' {
				l.next()
				return GE
			}
			return GT
		case ' ', '\t', '\n', '\r':
			// ignore whitespace
			l.emit()
		default:
			log.Printf("unrecognized character %q", c)
		}
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

func (l *yyLex) next() rune {
	if len(l.src) == l.start {
		return eof
	}
	r, w := utf8.DecodeRune(l.src[l.pos:])
	l.width = w
	l.pos += l.width
	return r
}

func (l *yyLex) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *yyLex) backup() {
	l.pos -= l.width
}

func (l *yyLex) emit() []byte {
	tok := l.src[l.start:l.pos]
	l.start = l.pos
	return tok
}

func (l *yyLex) Error(s string) {
	log.Printf("parse error: %s", s)
}
