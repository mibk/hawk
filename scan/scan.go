package scan

import (
	"io"
	"log"
	"strings"
)

type Scanner struct {
	Writer io.Writer
	line   string
	fields []string
}

func NewScanner(w io.Writer) *Scanner {
	return &Scanner{Writer: w}
}

func (sc *Scanner) SplitLine(line string) {
	sc.line = strings.TrimRight(line, "\r\n")
	sc.fields = strings.Fields(sc.line)
}

func (sc *Scanner) Field(i int) string {
	switch {
	case i < 0:
		log.Fatal("attempt to access field -1")
	case i == 0:
		return sc.line
	case i <= len(sc.fields):
		return sc.fields[i-1]
	}
	return ""
}

// NF returns number of fields of the current row.
func (sc *Scanner) NF() int {
	return len(sc.fields)
}
