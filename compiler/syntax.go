package compiler

import (
	"fmt"
	"io"
	"log"

	"github.com/mibk/hawk/parse"
)

type Tree struct {
	begin    []Stmt
	pActions []Stmt
	end      []Stmt

	vars map[string]Value
}

func NewTree() *Tree {
	return &Tree{vars: make(map[string]Value)}
}

func (t Tree) Begin() {
	for _, a := range t.begin {
		a.Exec()
	}
}

func (t Tree) End() {
	for _, a := range t.end {
		a.Exec()
	}
}

func (t Tree) AnyPatternActions() bool {
	return len(t.pActions) > 0 || len(t.end) > 0
}

func (t Tree) Exec() {
	for _, a := range t.pActions {
		a.Exec()
	}
}

type BeginAction struct {
	Stmt
}

type EndAction struct {
	Stmt
}

type PatternAction struct {
	pattern Expr
	action  Stmt
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

type AssignStmt struct {
	tree *Tree
	name string
	expr Expr
}

func (a AssignStmt) Exec() {
	a.tree.vars[a.name] = a.expr.Val()
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

type Ident struct {
	tree *Tree
	name string
}

func (i Ident) Val() Value {
	return i.tree.vars[i.name]
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
