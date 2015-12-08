package compiler

import (
	"io"
	"log"
	"os/exec"
	"strings"
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

func (e ExprStmt) Exec(w io.Writer) Status {
	e.expr.Eval(w)
	return StatusNone
}

type BlockStmt struct {
	stmts []Stmt
}

func (b BlockStmt) Exec(w io.Writer) Status {
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

func (p PipeStmt) Exec(w io.Writer) Status {
	// TODO: better method for argument parsing. (Arguments could be in quotes.)
	args := strings.Fields(p.cmd)
	if len(args) == 0 {
		log.Fatal("no command specified")
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout, cmd.Stderr = w, w
	wc, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	st := p.stmt.Exec(wc)
	if err := wc.Close(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	return st
}

type AssignStmt struct {
	scope Scope
	name  string
	expr  Expr
}

func (a AssignStmt) Exec(w io.Writer) Status {
	a.scope.SetVar(a.name, a.expr.Eval(w))
	return StatusNone
}

type IfStmt struct {
	expr     Expr
	stmt     Stmt
	elseStmt Stmt
}

func (i IfStmt) Exec(w io.Writer) Status {
	if i.expr.Eval(w).Bool() {
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
	body Stmt
}

func (f ForStmt) Exec(w io.Writer) Status {
	if f.init != nil {
		f.init.Exec(w)
	}
	for f.cond == nil || f.cond.Eval(w).Bool() {
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

type StatusStmt struct {
	status Status
}

func (s StatusStmt) Exec(io.Writer) Status {
	return s.status
}

type ReturnStmt struct {
	tree *Program
	expr Expr
}

func (r ReturnStmt) Exec(w io.Writer) Status {
	if r.expr != nil {
		r.tree.retval = r.expr.Eval(w)
	}
	return StatusReturn
}

type CallStmt CallExpr

func (c CallStmt) Exec(w io.Writer) Status {
	CallExpr(c).Eval(w)
	return StatusNone
}
