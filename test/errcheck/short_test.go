package test

import (
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
