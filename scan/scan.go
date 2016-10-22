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
	br  *bufio.Reader
	err error // sticky err

	recNumber int
	rec       string
	fields    []string
}

// SetReader sets an io.Reader for scanner to read from.
func (sc *Scanner) SetReader(r io.Reader) {
	sc.br = bufio.NewReader(r)
	sc.recNumber = 0
}

// Scan scans another record and parses it into fields. It there
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
	sc.splitRecord(line)
	sc.recNumber++
	return true
}

func (sc *Scanner) splitRecord(rec []byte) {
	sc.rec = string(bytes.TrimRight(rec, "\r\n"))
	sc.fields = strings.Fields(sc.rec)
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
		return sc.rec
	case i <= len(sc.fields):
		return sc.fields[i-1]
	}
	return ""
}

// NR returns the current record number.
func (sc *Scanner) NR() int {
	return sc.recNumber
}

// NF returns number of fields of the current row.
func (sc *Scanner) NF() int {
	return len(sc.fields)
}
