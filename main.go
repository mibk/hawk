package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/mibk/hawk/compiler"
	"github.com/mibk/hawk/parse"
)

var (
	file = flag.String("f", "", "read program from `file`")
)

func main() {
	flag.Parse()
	p := parse.NewParser(os.Stdout)
	prog := compile(p)

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

func compile(p *parse.Parser) *compiler.Tree {
	var srcr io.Reader
	if *file != "" {
		f, err := os.Open(*file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "hawk: %v\n", err)
		}
		defer f.Close()
		srcr = f
	} else if len(os.Args) < 2 {
		fmt.Println(`usage:
	hawk <program>

It is possible to write the program in a separate file and then call:
	xargs -0 -a <program-file> hawk`)
		os.Exit(1)
	} else {
		srcr = strings.NewReader(os.Args[1])
	}
	return compiler.Compile(srcr, p)
}
