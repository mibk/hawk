package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	"github.com/mibk/hawk/compiler"
	"github.com/mibk/hawk/parse"
)

func main() {
	p := parse.NewParser(os.Stdout)
	source := []byte(os.Args[1])
	prog := compiler.Compile(source, p)

	in := bufio.NewReader(os.Stdin)
	for {
		line, err := in.ReadBytes('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("ReadBytes: %s", err)
		}
		p.SetFields(strings.Fields(string(line)))
		prog.Exec()
	}
}
