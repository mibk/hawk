package compiler

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
	expr Expr
}

func (e *ExprStmt) Exec(w io.Writer) Status {
	e.expr.Eval(w)
	return StatusNone
}

type BlockStmt struct {
	stmts []Stmt
}

func (b *BlockStmt) Exec(w io.Writer) Status {
	for _, stmt := range b.stmts {
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
	stmt Stmt
	cmd  string
}

func (p *PipeStmt) Exec(w io.Writer) Status {
	// TODO: better method for argument parsing. (Arguments could be in quotes.)
	args := strings.Fields(p.cmd)
	if len(args) == 0 {
		throw("pipe statement: no command specified")
	}
	name := args[0]
	cmd := exec.Command(name, args[1:]...)
	cmd.Stdout, cmd.Stderr = w, w
	wc, err := cmd.StdinPipe()
	if err != nil {
		throw("%s: %v", name, err)
	}
	if err := cmd.Start(); err != nil {
		throw("%s: %v", name, err)
	}
	st := p.stmt.Exec(wc)
	if err := wc.Close(); err != nil {
		throw("%s: %v", name, err)
	}
	if err := cmd.Wait(); err != nil {
		throw("%s: %v", name, err)
	}
	return st
}

type AssignStmt struct {
	scope Scope
	left  Expr
	right Expr
}

func (a *AssignStmt) Exec(w io.Writer) Status {
	v := a.right.Eval(w)
	switch e := a.left.(type) {
	case *Ident:
		a.scope.SetVar(e.name, v)
	case *IndexExpr:
		a, ok := e.expr.Eval(w).Array()
		if !ok {
			throw("assigning to a scalar value using index expression")
		}
		var index *value.Scalar
		if e.index != nil {
			index, ok = e.index.Eval(w).Scalar()
			if !ok {
				throw("indexing an array using a non-scalar value")
			}
		}
		a.Put(index, v)
	default:
		panic(fmt.Sprintf("unknown assignment type: %T", e))
	}
	return StatusNone
}

type IfStmt struct {
	expr     Expr
	stmt     *BlockStmt
	elseStmt Stmt
}

func (i *IfStmt) Exec(w io.Writer) Status {
	v, ok := i.expr.Eval(w).Scalar()
	if !ok {
		throw("non-scalar value used as a condition")
	}
	if v.Bool() {
		return i.stmt.Exec(w)
	} else if i.elseStmt != nil {
		return i.elseStmt.Exec(w)
	}
	return StatusNone
}

type ForStmt struct {
	init Stmt
	cond Expr
	step Stmt
	body *BlockStmt
}

func (f *ForStmt) Exec(w io.Writer) Status {
	if f.init != nil {
		f.init.Exec(w)
	}
	for {
		if f.cond != nil {
			v, ok := f.cond.Eval(w).Scalar()
			if !ok {
				throw("non-scalar value used as a condition")
			}
			if !v.Bool() {
				break
			}
		}
		switch f.body.Exec(w) {
		case StatusBreak:
			break
		case StatusReturn:
			return StatusReturn
		}
		if f.step != nil {
			f.step.Exec(w)
		}
	}
	return StatusNone
}

type ForeachStmt struct {
	key  *Ident
	val  *Ident
	expr Expr
	body *BlockStmt
}

func (fs ForeachStmt) Exec(w io.Writer) Status {
	a, ok := fs.expr.Eval(w).Array()
	if !ok {
		throw("attempting to range over a scalar value")
	}
	for _, k := range a.Keys() {
		if fs.key != nil {
			fs.key.scope.SetVar(fs.key.name, k)
		}
		if fs.val != nil {
			fs.val.scope.SetVar(fs.val.name, a.Get(k))
		}
		switch fs.body.Exec(w) {
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
	tree *Program
	expr Expr
}

func (r *ReturnStmt) Exec(w io.Writer) Status {
	if r.expr != nil {
		r.tree.retval = r.expr.Eval(w)
	}
	return StatusReturn
}

type PrintStmt struct {
	fun  string
	args []Expr
}

func (p *PrintStmt) Exec(w io.Writer) Status {
	switch p.fun {
	case "print":
		var vals []interface{}
		for _, e := range p.args {
			vals = append(vals, e.Eval(w))
		}
		fmt.Fprintln(w, vals...)
	case "printf":
		format, vals, err := formatPrintfArgs(w, "printf", p.args)
		if err != nil {
			throw("%v", err)
		}
		fmt.Fprintf(w, format, vals...)
	default:
		panic("unknown print function: " + p.fun)
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
