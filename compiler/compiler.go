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
	return &Program{*p}, nil
}

// SetFieldSep sets the field separator, FS, to the sep value.
func (p *Program) SetFieldSep(sep string) { p.prog.SetFieldSep(sep) }

// Run runs the program. It scans src and writes output to w.
func (p *Program) Run(w io.Writer, src scan.Source) error {
	return p.prog.Run(w, src)
}
