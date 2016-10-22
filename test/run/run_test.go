package test

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/mibk/hawk/compiler"
	"github.com/mibk/hawk/scan"
)

func TestRun(t *testing.T) {
	fis, _ := ioutil.ReadDir(".")
	for _, fi := range fis {
		if !fi.IsDir() && strings.HasSuffix(fi.Name(), ".hawk") {
			progname := strings.TrimSuffix(fi.Name(), ".hawk")
			t.Run(progname, func(t *testing.T) {
				var out bytes.Buffer
				sc := scan.NewScanner(&out)
				f, err := os.Open(progname + ".hawk")
				if err != nil {
					panic(err)
				}
				prog, err := compiler.Compile(f, sc)
				if err != nil {
					t.Errorf("%s.hawk:%v", progname, err)
					return
				}

				fin, _ := os.Open(progname + ".in")
				prog.Run(fin)

				var want []byte
				fout, _ := os.Open(progname + ".out")
				if fout != nil {
					want, _ = ioutil.ReadAll(fout)
				}
				if got := out.Bytes(); bytes.Compare(got, want) != 0 {
					t.Errorf("%s.out:\n got:\n%s\nwant:\n%s", progname, string(got), string(want))
				}
			})
		}
	}
}
