package scan

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"strings"
)

// A Scanner is used for splitting input into rows and
// splitting rows into fields.
type Scanner struct {
	br     *bufio.Reader
	line   string
	fields []string
	err    error // sticky err
}

// SetReader sets an io.Reader for scanner to read from.
func (sc *Scanner) SetReader(r io.Reader) {
	sc.br = bufio.NewReader(r)
}

// Scan scans another row and parses it into fields. It there
// is an error or EOF is reached, Scan returns false. Otherwise
// it returns true.
func (sc *Scanner) Scan() bool {
	if sc.err != nil {
		return false
	}
	if sc.br == nil {
		sc.err = errors.New("scan: nil reader")
		return false
	}

	line, err := sc.br.ReadBytes('\n')
	if err == io.EOF {
		return false
	} else if err != nil {
		sc.err = err
		return false
	}
	sc.splitLine(line)
	return true
}

func (sc *Scanner) splitLine(line []byte) {
	sc.line = string(bytes.TrimRight(line, "\r\n"))
	sc.fields = strings.Fields(sc.line)
}

func (sc *Scanner) Err() error {
	return sc.err
}

// Field returns ith field from the current row.
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
