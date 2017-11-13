package compiler

import (
	"io"
	"sync"

	"github.com/mibk/hawk/compiler/internal/hawkc"
	"github.com/mibk/hawk/scan"
)

var mu sync.Mutex

// A Program represents a compiled Hawk program.
type Program struct {
	// FieldSep specifies the default field separator, FS.
	// If FieldSep is the empty string, Run runs the program
	// with the default behaviour, which is to split each
	// record into fields using one or more white spaces
	// characters as a separator.
	FieldSep string

	prog hawkc.Program
}

// Compile compiles a Hawk program (name) from src. name is there
// only for better error printing.
func Compile(name string, src io.Reader) (*Program, error) {
	mu.Lock()
	defer mu.Unlock()
	p, err := hawkc.Compile(name, src)
	if err != nil {
		return nil, err
	}
	return &Program{prog: *p}, nil
}

// Run runs the program. It scans src and writes output to w.
func (p *Program) Run(w io.Writer, src scan.Source) error {
	if p.FieldSep != "" {
		p.prog.SetFieldSep(p.FieldSep)
	}
	return p.prog.Run(w, src)
}
