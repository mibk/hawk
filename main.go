package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/mibk/hawk/compiler"
	"github.com/mibk/hawk/scan"
)

var (
	helpFlag = flag.Bool("help", false, "display an extended help")

	file     = flag.String("f", "", "read program from `file`")
	fieldSep = flag.String("F", "", "set the field separator, FS")
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("hawk: ")
	flag.Usage = usage
	flag.Parse()

	if *helpFlag {
		fmt.Fprint(os.Stderr, extendedHelp)
		os.Exit(2)
	}

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

	var input scan.Source = os.Stdin
	if len(args) > 0 {
		srcs := make([]scan.Source, 0, len(args))
		for _, file := range args {
			f, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			srcs = append(srcs, f)
		}
		input = scan.MultiSource(srcs...)
	}

	prog, err := compiler.Compile(name, srcCode)
	if err != nil {
		log.Fatal(err)
	}
	if *fieldSep != "" {
		prog.SetFieldSep(*fieldSep)
	}
	if err := prog.Run(os.Stdout, input); err != nil {
		log.Fatal(err)
	}
}

const shortHelp = `Usage: hawk 'program' [file ...]
  or:  hawk -f progfile [file ...]

Hawk is an Awk clone. Program is a set of pattern {action} pairs. Hawk reads
from all of the present files and for each line of each file executes all the
provided pairs. If no files are present, hawk reads from stdin.
`

func usage() {
	fmt.Fprintln(os.Stderr, shortHelp+`
Run hawk -help for a detailed help message.

Flags:`)
	flag.PrintDefaults()
}
