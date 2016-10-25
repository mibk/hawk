package test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/mibk/hawk/compiler"
)

var valid = []struct {
	prog string
}{
	{`{}`},
	{`x`},
	{`x > 3`},
	{`$1 > 3`},
	{`{ print $0 }`},
	{`{} // `},
	{`{ "\a\b\f\n\r\t\v\\\"'" }`},
	{`{ '\a\b\f\n\r\t\v\\"\'' }`},
}

func TestValid(t *testing.T) {
	for i, tt := range valid {
		b := strings.NewReader(tt.prog)
		if _, err := compiler.Compile(b); err != nil {
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
	{`/* `, "1: eof in block comment"},
	{`" `, "1: eof in string literal"},
	{`' `, "1: eof in string literal"},
	{`"
		"`, "2: newline in string literal"},
	{`'
		'`, "2: newline in string literal"},
	{`"\e"`, `1: unknown escape character \e`},
	{`"\i"`, `1: unknown escape character \i`},
}

func TestErrors(t *testing.T) {
	for i, tt := range testProgs {
		b := strings.NewReader(tt.prog)
		_, err := compiler.Compile(b)
		if err == nil {
			t.Errorf("%d: test unexpectedly succeded", i+1)
			continue
		}
		if err.Error() != tt.err {
			t.Errorf("test %d:\n got: %v\nwant: %v", i+1, err, tt.err)
		}
	}
}

// All programs are wrapped in 'BEGIN { }' before executing.
var runtimeInvalid = []struct {
	prog string
	err  string
}{
	0:  {`x = 0; x[0] = 2`, "assigning to a scalar value using index expression"},
	1:  {`a = []; if a {}`, "non-scalar value used as a condition"},
	2:  {`doesntexist()`, "unknown function: doesntexist"},
	3:  {`a = []; print $a`, "attempting to access a field using a non-scalar value"},
	4:  {`sin(a, b)`, "sin: 1 != 2: argument count mismatch"},
	5:  {`a = []; cos(a)`, "cos: all arguments must be scalar values"},
	6:  {`a = "scalar"; for x in a {}`, "attempting to range over a scalar value"},
	7:  {`[] < ""`, "cannot compare array and string using <, >, <=, or >="},
	8:  {`[] < x`, "cannot compare array and array using <, >, <=, or >="},
	9:  {`[] < 50`, "cannot compare array and number using <, >, <=, or >="},
	10: {`[] < []`, "cannot compare array and array using <, >, <=, or >="},
}

func TestRuntimeErrors(t *testing.T) {
	for i, tt := range runtimeInvalid {
		b := strings.NewReader("BEGIN { " + tt.prog + " }")
		prog, err := compiler.Compile(b)
		if err != nil {
			t.Errorf("test %d: unexpected err: %v", i, err)
			continue
		}
		err = prog.Run(ioutil.Discard, nil)
		if err == nil {
			t.Errorf("%d: test unexpectedly succeded", i)
			continue
		}
		if err.Error() != tt.err {
			t.Errorf("test %d:\n got: %v\nwant: %v", i, err, tt.err)
		}
	}
}
