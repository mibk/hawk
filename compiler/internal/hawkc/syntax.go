package hawkc

import (
	"fmt"
	"io"

	"github.com/mibk/hawk/scan"
	"github.com/mibk/hawk/value"
)

type Decl interface{}

type Scope interface {
	Get(name string) value.Value
	Put(name string, v value.Value)
}

type Program struct {
	sc     *scan.Scanner
	vars   map[string]value.Value
	funcs  map[string]*FuncDecl
	retval value.Value

	Begins   []Stmt
	Pactions []Stmt
	Ends     []Stmt
}

func NewProgram(sc *scan.Scanner) *Program {
	return &Program{
		sc:    sc,
		vars:  make(map[string]value.Value),
		funcs: make(map[string]*FuncDecl),
	}
}

func (p *Program) Get(name string) value.Value {
	if v, ok := p.vars[name]; ok {
		return v
	}
	// Global "magic" variables.
	switch name {
	case "NR":
		return value.NewNumber(float64(p.sc.RecordNumber()))
	case "NF":
		return value.NewNumber(float64(p.sc.FieldCount()))
	case "FILENAME":
		return value.NewString(p.sc.Filename())
	case "FNR":
		return value.NewNumber(float64(p.sc.FileRecordNumber()))
	}
	v := &value.Undefined{}
	p.vars[name] = v
	return v
}

func (p *Program) Put(name string, v value.Value) {
	p.vars[name] = v
}

func (p *Program) Run(out io.Writer, in scan.Source) (err error) {
	defer func() {
		if err == nil {
			if v := recover(); v != nil {
				e, ok := v.(*runtimeError)
				if !ok {
					panic(v)
				}
				err = e
			}
		}
	}()
	p.Begin(out)
	if p.anyPatternActions() {
		p.sc.SetSource(in)
		for p.sc.Scan() {
			for _, a := range p.Pactions {
				a.Exec(out)
			}
		}
		if err := p.sc.Err(); err != nil {
			return err
		}
		p.End(out)
	}
	// TODO: return something like p.Err()
	return nil
}

func (p *Program) Begin(w io.Writer) {
	for _, a := range p.Begins {
		a.Exec(w)
	}
}

func (p *Program) End(w io.Writer) {
	for _, a := range p.Ends {
		a.Exec(w)
	}
}

func (p *Program) anyPatternActions() bool {
	return len(p.Pactions) > 0 || len(p.Ends) > 0
}

type BeginAction struct {
	Stmt
}

type EndAction struct {
	Stmt
}

type PatternAction struct {
	Pattern Expr
	Body    Stmt
}

func (p *PatternAction) Exec(w io.Writer) Status {
	if p.Pattern != nil {
		v, ok := p.Pattern.Eval(w).Scalar()
		if !ok {
			throw("Pattern in an Body must be a scalar value")
		}
		if !v.Bool() {
			return StatusNone
		}
	}
	p.Body.Exec(w)
	return StatusNone
}

type FuncDecl struct {
	scope *FuncScope
	Name  string
	Args  []string
	Body  Stmt
}

type FuncScope struct {
	stack []map[string]value.Value
}

func (f *FuncScope) Push() {
	f.stack = append(f.stack, make(map[string]value.Value))
}

func (f *FuncScope) Pull() {
	f.stack = f.stack[:len(f.stack)-1]
}

func (f *FuncScope) Get(name string) value.Value {
	s := f.currScope()
	if v, ok := s[name]; ok {
		return v
	}
	v := &value.Undefined{}
	s[name] = v
	return v
}

func (f *FuncScope) Put(name string, v value.Value) {
	f.currScope()[name] = v
}

func (f *FuncScope) currScope() map[string]value.Value {
	if f.stack == nil {
		panic("stack shouldn't be nil")
	}
	return f.stack[len(f.stack)-1]
}

func throw(format string, args ...interface{}) {
	panic(&runtimeError{fmt.Errorf(format, args...)})
}

type runtimeError struct {
	error
}

type debugInfo struct {
	line int
}

func genDebugInfo() debugInfo {
	return debugInfo{lexlineno}
}

func (di debugInfo) throw(format string, args ...interface{}) {
	throw(fmt.Sprintf("line %d:", di.line)+format, args...)
}
