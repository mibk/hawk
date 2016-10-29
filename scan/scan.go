package scan

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"regexp"
	"strings"
)

// endOfSource is returned when io.EOF is reached in one of
// the sources and there are still more sources to process.
var endOfSource = errors.New("end of source")

type Source interface {
	io.Reader

	// Name returns the name of the source.
	Name() string
}

func MultiSource(sources ...Source) Source {
	return &multiSource{sources}
}

type multiSource struct {
	sources []Source
}

func (ms *multiSource) Read(p []byte) (n int, err error) {
	for len(ms.sources) > 0 {
		n, err = ms.sources[0].Read(p)
		if err == io.EOF {
			ms.sources = ms.sources[1:]
			if len(ms.sources) > 0 {
				err = endOfSource
				return
			}
		}
		if n > 0 || err != nil {
			return
		}
	}
	return 0, io.EOF
}

func (ms *multiSource) Name() string {
	if len(ms.sources) == 0 {
		return ""
	}
	return ms.sources[0].Name()
}

// A Scanner is used for splitting input into rows and
// splitting rows into fields.
type Scanner struct {
	src      Source // To retrieve filename only.
	br       *bufio.Reader
	fieldsRx *regexp.Regexp
	err      error // sticky err

	recNumber     int
	fileRecNumber int
	rec           string
	fields        []string
}

// SetReader sets an io.Reader for scanner to read from.
func (sc *Scanner) SetSource(src Source) {
	sc.src = src
	sc.br = bufio.NewReader(src)
	sc.recNumber = 0
}

// SetFieldSep sets regexp rx that will be used to separate
// row into fields.
func (sc *Scanner) SetFieldSep(rx string) {
	if sc.err != nil {
		return
	}
	sc.fieldsRx, sc.err = regexp.Compile(rx)
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

readRecord:
	line, err := sc.br.ReadBytes('\n')
	if err == io.EOF {
		return false
	} else if err == endOfSource {
		sc.fileRecNumber = 0
		goto readRecord
	} else if err != nil {
		sc.err = err
		return false
	}
	sc.splitRecord(line)
	sc.recNumber++
	sc.fileRecNumber++
	return true
}

func (sc *Scanner) splitRecord(rec []byte) {
	sc.rec = string(bytes.TrimRight(rec, "\r\n"))
	if sc.fieldsRx != nil {
		sc.fields = sc.fieldsRx.Split(sc.rec, -1)
		if len(sc.fields) > 0 && sc.fields[0] == "" {
			sc.fields = sc.fields[1:]
		}
		if len(sc.fields) > 0 && sc.fields[len(sc.fields)-1] == "" {
			sc.fields = sc.fields[:len(sc.fields)-1]
		}
	} else {
		sc.fields = strings.Fields(sc.rec)
	}
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

// RecordNumber returns the current record number.
func (sc *Scanner) RecordNumber() int {
	return sc.recNumber
}

// FieldCount returns number of fields of the current row.
func (sc *Scanner) FieldCount() int {
	return len(sc.fields)
}

// Filename returns the name of the currently processed file.
func (sc *Scanner) Filename() string {
	return sc.src.Name()
}

// FileRecordNumber returns the current record number in the currently
// processed file.
func (sc *Scanner) FileRecordNumber() int {
	return sc.fileRecNumber
}
