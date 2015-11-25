package compiler

import (
	"fmt"
	"io"
	"log"

	"github.com/mibk/hawk/parse"
)

type Tree struct {
	pAction PatternAction
}

func (t Tree) Exec() { t.pAction.Exec() }

type PatternAction struct {
	pattern Expr
	action  BlockStmt
}

func (p PatternAction) Exec() {
	if p.pattern.Val().Cmp(NewBoolValue(true)) == 0 {
		p.action.Exec()
	}
}

type BlockStmt struct {
	stmts []Stmt
}

func (b BlockStmt) Exec() {
	for _, stmt := range b.stmts {
		stmt.Exec()
	}
}

type Stmt interface {
	Exec()
}

type ExprStmt struct {
	expr Expr
}

func (e ExprStmt) Exec() {
	e.expr.Val()
}

type Expr interface {
	Val() Value
}

type CallExpr struct {
	writer io.Writer
	fun    string
	args   []Expr
}

func (c CallExpr) Val() Value {
	switch c.fun {
	case "print":
		var vals []interface{}
		for _, e := range c.args {
			vals = append(vals, e.Val())
		}
		fmt.Fprintln(c.writer, vals...)
	default:
		log.Fatalf("unknown func %s", c.fun)
	}
	return NewBoolValue(false)
}

type Col struct {
	p   *parse.Parser
	Num Expr
}

func (c Col) Val() Value {
	n := c.Num.Val().Int()
	return NewStringValue(c.p.Field(n))
}

type BinaryOp struct {
	typ   int
	left  Expr
	right Expr
}

func (c BinaryOp) Val() Value {
	cmp := c.left.Val().Cmp(c.right.Val())
	var b bool
	switch c.typ {
	case EQ:
		b = cmp == 0
	case NE:
		b = cmp != 0
	case LT:
		b = cmp == -1
	case LE:
		b = cmp <= 0
	case GT:
		b = cmp == 1
	case GE:
		b = cmp >= 0
	default:
		panic("unreachable")
	}
	return NewBoolValue(b)
}

type Lit int

func (l Lit) Val() Value {
	return NewNumberValue(float64(l))
}
