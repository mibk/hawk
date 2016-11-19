package scan

import (
	"io"
	"regexp"
	"strings"
	"testing"
)

func TestRxReader(t *testing.T) {
	// Reduce the default buffer size to make it easier to test
	// the edge cases.
	_bufSize = 8

	tests := []struct {
		rx    string
		src   Source
		lines []string
	}{
		0: {`-+`, stringSrcs("All----work and no play makes Jack-a dull boy"),
			[]string{"All", "work and no play makes Jack", "a dull boy"}},
		1: {`\.+`, stringSrcs("123456.." + "...45678"),
			[]string{"123456", "45678"}},
		2: {`\.+`, stringSrcs("123456.." + "123"),
			[]string{"123456", "123"}},
		3: {`#+`, stringSrcs("abcdefghijklmnop####"),
			[]string{"abcdefghijklmnop"}},

		4: {`-+`, stringSrcs("AA--BB", "--DDD-"),
			[]string{"AA", "BB", "", "DDD"}},
		5: {`\.+`, stringSrcs("123456..", "...45678"),
			[]string{"123456", "", "45678"}},

		6: {`-+`, dummySource{earlyEOFReader{"A-A-BB"}},
			[]string{"A", "A", "BB"}},
	}

	for j, tt := range tests {
		rx := regexp.MustCompile(tt.rx)
		rr := newRxLineReader(tt.src, rx)
		for i := 0; ; i++ {
			line, err := rr.ReadLine()
			if err == io.EOF {
				if i != len(tt.lines) {
					t.Errorf("test[%d]: not enough lines: got %d, want %d",
						j, i, len(tt.lines))
				}
				break
			} else if err != nil {
				t.Errorf("test[%d]: unexpected err: %v", j, err)
				break
			}
			if i >= len(tt.lines) {
				t.Errorf("test[%d]: unexpected %dth line %q; expected %d", j, i+1, line, len(tt.lines))
				break
			}
			if string(line) != tt.lines[i] {
				t.Errorf("test[%d]: %d: got %q, want %q", j, i, line, tt.lines[i])
			}
		}
	}
}

func stringSrcs(s ...string) Source {
	var srcs []Source
	for _, s := range s {
		srcs = append(srcs, dummySource{strings.NewReader(s)})
	}
	if len(srcs) == 1 {
		return srcs[0]
	}
	return MultiSource(srcs...)
}

type dummySource struct {
	io.Reader
}

func (ds dummySource) Name() string { return "<anonymous>" }

type earlyEOFReader struct {
	s string
}

func (r earlyEOFReader) Read(p []byte) (n int, err error) {
	if len(p) < len(r.s) {
		panic("p must be longer than s to make the test work")
	}
	n = copy(p, r.s)
	return n, io.EOF
}
