package test

import (
	"io"
	"strings"
	"testing"

	"github.com/mibk/hawk/compiler"
	"github.com/mibk/hawk/parse"
)

var parser = parse.NewParser(io.Writer(nil))

var valid = []struct {
	prog string
}{
	{`{}`},
	{`x`},
	{`x > 3`},
	{`$1 > 3`},
	{`{ print $0 }`},
}

func TestValid(t *testing.T) {
	for i, tt := range valid {
		b := strings.NewReader(tt.prog)
		if _, err := compiler.Compile(b, parser); err != nil {
			t.Errorf("test %d: unexpected err: %v", i+1, err)
		}
	}
}

var testProgs = []struct {
	prog string
	err  string
}{
	{`BEGIN {
	} BEGIN`, "2: syntax error: unexpected BEGIN, expecting ';'"},
	{`BEGIN { 00 = 20 }`, "1: syntax error: unexpected '=', expecting '}'"},
}

func TestErrors(t *testing.T) {
	for i, tt := range testProgs {
		b := strings.NewReader(tt.prog)
		_, err := compiler.Compile(b, parser)
		if err == nil {
			t.Errorf("%d: test unexpectedly succeded", i+1)
			continue
		}
		if err.Error() != tt.err {
			t.Errorf("test %d:\n got: %v\nwant: %v", i+1, err, tt.err)
		}
	}
}
