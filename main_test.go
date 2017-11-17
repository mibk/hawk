package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mibk/hawk/compiler"
)

func TestRun(t *testing.T) {
	files, err := filepath.Glob("testdata/*.hawk")
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		name := strings.TrimSuffix(filepath.Base(file), ".hawk")
		t.Run(name, func(t *testing.T) {
			f, err := os.Open(file)
			if err != nil {
				t.Fatal(err)
			}
			prog, err := compiler.Compile(file, f)
			if err != nil {
				t.Errorf("%s.hawk:%v", name, err)
				return
			}

			var out bytes.Buffer
			fin, err := os.Open(strings.TrimSuffix(file, ".hawk") + ".in")
			if err != nil && !os.IsNotExist(err) {
				t.Fatal(err)
			}
			if err := prog.Run(&out, fin); err != nil {
				t.Errorf("unexpected runtime err: %v", err)
				return
			}

			var want []byte
			fout, err := os.Open(strings.TrimSuffix(file, ".hawk") + ".out")
			if err != nil && !os.IsNotExist(err) {
				t.Fatal(err)
			}
			if fout != nil {
				want, err = ioutil.ReadAll(fout)
				if err != nil {
					t.Fatal(err)
				}
			}
			if got := out.Bytes(); !bytes.Equal(got, want) {
				t.Errorf("%s.out:\n got:\n`%s`\nwant:\n`%s`", name, string(got), string(want))
			}
		})
	}
}
