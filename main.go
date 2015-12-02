package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
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
	name, src := getSource()
	prog, err := compiler.Compile(src, p)
	src.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "hawk: %s:%v", name, err)
		os.Exit(1)
	}
	prog.Run(os.Stdin)
}

func getSource() (name string, src io.ReadCloser) {
	name = "line"
	if *file != "" {
		f, err := os.Open(*file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "hawk: %v\n", err)
		}
		src = f
		name = *file
	} else if len(os.Args) < 2 {
		fmt.Println(`usage:
	hawk <program>

It is possible to write the program in a separate file and then call:
	xargs -0 -a <program-file> hawk`)
		os.Exit(2)
	} else {
		src = ioutil.NopCloser(strings.NewReader(os.Args[1]))
	}
	return
}
