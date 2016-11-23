package compiler

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

// Found by go-fuzz
var crashers = []struct {
	prog string
}{
	0: {`{m[$0]=0; for k,v in m{print v<k}}`},
}

func TestCrashers(t *testing.T) {
	for i, tt := range crashers {
		prog, err := Compile("fuzz", strings.NewReader(tt.prog))
		if err != nil {
			t.Errorf("test[%d]: unexpected compile err: %v", i, err)
		}
		src := withDummyName{strings.NewReader(input)}
		if err := prog.Run(ioutil.Discard, src); err != nil {
			t.Errorf("test[%d]: unexpected runtime err: %v", i, err)
		}
	}
}

const input = `one two three
four five six seven
eight 999999 10
eleven

`

type withDummyName struct {
	io.Reader
}

func (withDummyName) Name() string { return "dummy" }
