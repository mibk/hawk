package parse

import (
	"io"
	"log"
	"strings"
)

type Parser struct {
	Writer io.Writer
	line   string
	fields []string
}

func NewParser(w io.Writer) *Parser {
	return &Parser{Writer: w}
}

func (p *Parser) SplitLine(line string) {
	p.line = strings.TrimRight(line, "\r\n")
	p.fields = strings.Fields(p.line)
}

func (p *Parser) Field(i int) string {
	switch {
	case i < 0:
		log.Fatal("attempt to access field -1")
	case i == 0:
		return p.line
	case i <= len(p.fields):
		return p.fields[i-1]
	}
	return ""
}

// NF returns number of fields of the current row.
func (p *Parser) NF() int {
	return len(p.fields)
}
