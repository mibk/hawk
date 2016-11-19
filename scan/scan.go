package scan

import (
	"bufio"
	"errors"
	"io"
	"log"
	"regexp"
	"strings"
)

// endOfSource is returned when io.EOF is reached in one of
// the sources and there are still more sources to process.
var endOfSource = errors.New("end of source")

// Source is the interface thats wraps io.Reader and provides
// the Name method.
type Source interface {
	io.Reader

	// Name returns the name of the source.
	Name() string
}

// MultiSource returns a Source that's the logical concatenation
// of the provided input sources.
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
	lr       lineReader
	rowsRx   *regexp.Regexp
	fieldsRx *regexp.Regexp
	err      error // sticky err

	recNumber     int
	fileRecNumber int
	rec           string
	fields        []string
}

// SetReader sets an io.Reader for scanner to read from.
func (sc *Scanner) SetSource(src Source) {
	if sc.rowsRx != nil {
		sc.lr = newRxLineReader(src, sc.rowsRx)
	} else {
		sc.lr = newSimpleLineReader(src)
	}
	sc.recNumber = 0
}

// SetRowSep sets regexp rx that will be used to separate
// input into rows.
func (sc *Scanner) SetRowSep(rx string) {
	if sc.err != nil {
		return
	}
	sc.rowsRx, sc.err = regexp.Compile(rx)
	if sc.err == nil && sc.lr != nil {
		sc.lr = newRxLineReader(sc.lr, sc.rowsRx)
	}
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
	if sc.lr == nil {
		sc.err = errors.New("scan: nil reader")
		return false
	}

readRecord:
	line, err := sc.lr.ReadLine()
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
	sc.rec = string(rec)
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

// Filename returns the name of the currently processed source.
func (sc *Scanner) Filename() string {
	return sc.lr.Name()
}

// FileRecordNumber returns the current record number in the currently
// processed file.
func (sc *Scanner) FileRecordNumber() int {
	return sc.fileRecNumber
}

type lineReader interface {
	Source // to be able to read buffered data
	ReadLine() ([]byte, error)
}

type simpleLineReader struct {
	src  Source
	name string // name of the current source
	br   *bufio.Reader
}

func newSimpleLineReader(src Source) *simpleLineReader {
	return &simpleLineReader{
		src:  src,
		name: src.Name(),
		br:   bufio.NewReader(src),
	}
}

func (sr *simpleLineReader) Read(p []byte) (n int, err error) {
	n, err = sr.br.Read(p)
	if err == endOfSource {
		sr.name = sr.src.Name()
	}
	return
}

func (sr *simpleLineReader) Name() string { return sr.name }

func (sr *simpleLineReader) ReadLine() ([]byte, error) {
	line, err := sr.br.ReadBytes('\n')
	if len(line) > 0 {
		line = line[:len(line)-1] // remove '\n'
	}
	return line, err
}

const bufSize = 4096

var _bufSize = bufSize // for testing purposes

type rxLineReader struct {
	buf      [bufSize]byte
	ptr      []byte
	src      Source
	name     string // name of the current source
	rx       *regexp.Regexp
	eos      bool
	finished bool
}

func newRxLineReader(src Source, sepRx *regexp.Regexp) *rxLineReader {
	return &rxLineReader{
		src:  src,
		name: src.Name(),
		rx:   sepRx,
	}
}

func (rr *rxLineReader) Read(p []byte) (n int, err error) {
	if len(rr.ptr) > 0 {
		n := copy(p, rr.ptr)
		rr.ptr = rr.ptr[n:]
		return n, nil
	}
	return rr.src.Read(p)
}

func (rr *rxLineReader) Name() string {
	if len(rr.ptr) > 0 {
		return rr.name
	}
	return rr.src.Name()
}

func (rr *rxLineReader) ReadLine() (line []byte, err error) {
	var loc []int
	for {
		if len(rr.ptr) == 0 {
		End:
			if rr.finished {
				if len(line) > 0 {
					if loc != nil {
						line = line[:loc[0]]
					}
					return line, nil
				}
				return nil, io.EOF
			}
			if err := rr.loadBuf(); err != nil {
				return nil, err
			}
			if rr.finished {
				goto End
			}
		}
		line = append(line, rr.ptr...)
		loc = rr.rx.FindIndex(line)

		if loc == nil || loc[1] == len(line) {
			rr.ptr = nil

			if rr.eos {
				rr.eos = false
				if loc != nil {
					line = line[:loc[0]]
				}
				return line, nil
			}
			continue
		}

		rr.ptr = line[loc[1]:]
		return line[:loc[0]], nil
	}
}

func (rr *rxLineReader) loadBuf() error { return rr.loadBufN(_bufSize) }

func (rr *rxLineReader) loadBufN(n int) error {
	m, err := rr.src.Read(rr.buf[:n])
	rr.ptr = rr.buf[:m]
	if err == io.EOF {
		rr.finished = true
	} else if err == endOfSource {
		rr.name = rr.src.Name()
		rr.eos = true
	} else {
		return err
	}
	return nil
}
