package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/mibk/hawk/compiler"
)

var (
	file = flag.String("f", "", "read program from `file`")
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("hawk: ")
	flag.Usage = usage
	flag.Parse()

	var srcCode io.Reader
	args := flag.Args()
	name := "line"
	if *file != "" {
		f, err := os.Open(*file)
		if err != nil {
			log.Fatal(err)
		}
		name = *file
		srcCode = f
		defer f.Close()
	} else {
		if len(args) == 0 {
			flag.Usage()
			os.Exit(2)
		}
		srcCode = strings.NewReader(args[0])
		args = args[1:]
	}

	var input io.Reader = os.Stdin
	if len(args) > 0 {
		rds := make([]io.Reader, 0, len(args))
		for _, file := range args {
			f, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			rds = append(rds, f)
		}
		input = io.MultiReader(rds...)
	}

	prog, err := compiler.Compile(srcCode)
	if err != nil {
		log.Fatalf("%s:%v", name, err)
	}
	if err := prog.Run(os.Stdout, input); err != nil {
		log.Fatalf("executing: %v", err)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, `Usage: hawk 'program text' [file ...]
  or:  hawk -f program-file [file ...]

Hawk is an Awk clone. Program is a set of PATTERN { ACTION } pairs. Hawk reads
from all of the present files and for each line of each file executes all the
provided pairs. If no files are present, hawk reads from stdin.

Flags:`)
	flag.PrintDefaults()
}
