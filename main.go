package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/mibk/hawk/compiler"
	"github.com/mibk/hawk/parse"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`usage:
	hawk <program>

It is possible to write the program in a separate file and then call:
	xargs -0 -a <program-file> hawk`)
		os.Exit(1)
	}
	p := parse.NewParser(os.Stdout)
	prog := compiler.Compile(strings.NewReader(os.Args[1]), p)

	prog.Begin()
	if prog.AnyPatternActions() {
		in := bufio.NewReader(os.Stdin)
		for {
			line, err := in.ReadBytes('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("ReadBytes: %s", err)
			}
			p.SetFields(strings.Fields(string(line)))
			prog.Exec()
		}
		prog.End()
	}
}
