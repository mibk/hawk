package parse

import (
	"io"
	"log"
	"strings"
)

type Parser struct {
	Writer io.Writer
	fields []string
}

func NewParser(w io.Writer) *Parser {
	return &Parser{Writer: w}
}

func (p *Parser) SetFields(fields []string) {
	p.fields = fields
}

func (p *Parser) Field(i int) string {
	switch {
	case i < 0:
		log.Fatal("attempt to access field -1")
	case i == 0:
		return strings.Join(p.fields, " ")
	case i <= len(p.fields):
		return p.fields[i-1]
	}
	return ""
}
