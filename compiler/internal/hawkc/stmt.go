package hawkc

import (
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/mibk/hawk/value"
)

type Status int

const (
	StatusNone Status = iota
	StatusBreak
	StatusContinue
	StatusReturn
)

type Stmt interface {
	Exec(io.Writer) Status
}

type ExprStmt struct {
	X Expr
}

func (e *ExprStmt) Exec(w io.Writer) Status {
	e.X.Eval(w)
	return StatusNone
}

type BlockStmt struct {
	List []Stmt
}

func (b *BlockStmt) Exec(w io.Writer) Status {
	for _, stmt := range b.List {
		switch s := stmt.Exec(w); s {
		case StatusBreak, StatusReturn:
			return s
		case StatusContinue:
			return StatusNone
		}
	}
	return StatusNone
}

type PipeStmt struct {
	debugInfo
	Stmt Stmt
	Cmd  string
}

func (p *PipeStmt) Exec(w io.Writer) Status {
	// TODO: better method for argument parsing. (Arguments could be in quotes.)
	args := strings.Fields(p.Cmd)
	if len(args) == 0 {
		p.throw("pipe statement: no command specified")
	}
	name := args[0]
	cmd := exec.Command(name, args[1:]...)
	cmd.Stdout, cmd.Stderr = w, w
	wc, err := cmd.StdinPipe()
	if err != nil {
		p.throw("%s: %v", name, err)
	}
	if err := cmd.Start(); err != nil {
		p.throw("%s: %v", name, err)
	}
	st := p.Stmt.Exec(wc)
	if err := wc.Close(); err != nil {
		p.throw("%s: %v", name, err)
	}
	if err := cmd.Wait(); err != nil {
		p.throw("%s: %v", name, err)
	}
	return st
}

type AssignStmt struct {
	debugInfo
	scope Scope
	Left  Expr
	Right Expr
}

func (as *AssignStmt) Exec(w io.Writer) Status {
	v := as.Right.Eval(w)
	switch e := as.Left.(type) {
	case *Ident:
		as.scope.Put(e.Name, v)
	case *IndexExpr:
		a, ok := e.X.Eval(w).Array()
		if !ok {
			as.throw("assigning to a scalar value using index expression")
		}
		var index *value.Scalar
		if e.Index != nil {
			index, ok = e.Index.Eval(w).Scalar()
			if !ok {
				as.throw("indexing an array using a non-scalar value")
			}
		}
		a.Put(index, v)
	default:
		panic(fmt.Sprintf("unknown assignment type: %T", e))
	}
	return StatusNone
}

type IfStmt struct {
	debugInfo
	X    Expr
	Body *BlockStmt
	Else Stmt
}

func (is *IfStmt) Exec(w io.Writer) Status {
	v, ok := is.X.Eval(w).Scalar()
	if !ok {
		is.throw("non-scalar value used as a condition")
	}
	if v.Bool() {
		return is.Body.Exec(w)
	} else if is.Else != nil {
		return is.Else.Exec(w)
	}
	return StatusNone
}

type ForStmt struct {
	debugInfo
	Init Stmt
	Cond Expr
	Post Stmt
	Body *BlockStmt
}

func (f *ForStmt) Exec(w io.Writer) Status {
	if f.Init != nil {
		f.Init.Exec(w)
	}
	for {
		if f.Cond != nil {
			v, ok := f.Cond.Eval(w).Scalar()
			if !ok {
				f.throw("non-scalar value used as a condition")
			}
			if !v.Bool() {
				break
			}
		}
		switch f.Body.Exec(w) {
		case StatusBreak:
			break
		case StatusReturn:
			return StatusReturn
		}
		if f.Post != nil {
			f.Post.Exec(w)
		}
	}
	return StatusNone
}

type ForeachStmt struct {
	debugInfo
	Key  *Ident
	Val  *Ident
	X    Expr
	Body *BlockStmt
}

func (fs ForeachStmt) Exec(w io.Writer) Status {
	a, ok := fs.X.Eval(w).Array()
	if !ok {
		fs.throw("attempting to range over a scalar value")
	}
	for _, k := range a.Keys() {
		if fs.Key != nil {
			fs.Key.scope.Put(fs.Key.Name, k)
		}
		if fs.Val != nil {
			fs.Val.scope.Put(fs.Val.Name, a.Get(k))
		}
		switch fs.Body.Exec(w) {
		case StatusBreak:
			break
		case StatusReturn:
			return StatusReturn
		}
	}
	return StatusNone
}

type StatusStmt struct {
	status Status
}

func (s *StatusStmt) Exec(io.Writer) Status {
	return s.status
}

type ReturnStmt struct {
	root *Program
	X    Expr
}

func (r *ReturnStmt) Exec(w io.Writer) Status {
	if r.X != nil {
		r.root.retval = r.X.Eval(w)
	}
	return StatusReturn
}

type PrintStmt struct {
	debugInfo
	Fun  string
	Args []Expr
}

func (p *PrintStmt) Exec(w io.Writer) Status {
	switch p.Fun {
	case "print":
		var vals []interface{}
		for _, e := range p.Args {
			vals = append(vals, e.Eval(w))
		}
		fmt.Fprintln(w, vals...)
	case "printf":
		format, vals, err := formatPrintfArgs(w, "printf", p.Args)
		if err != nil {
			p.throw("%v", err)
		}
		fmt.Fprintf(w, format, vals...)
	default:
		panic("unknown print function: " + p.Fun)
	}
	return StatusNone
}

func formatPrintfArgs(w io.Writer, fname string, exprs []Expr) (string, []interface{}, error) {
	if len(exprs) == 0 {
		return "", nil, fmt.Errorf("%s: not enough arguments: 0", fname)
	}
	format, args := exprs[0], exprs[1:]
	var vals []interface{}
	for _, e := range args {
		vals = append(vals, e.Eval(w))
	}
	return format.Eval(w).String(), vals, nil
}
