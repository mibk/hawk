package compiler

import (
	"fmt"
	"io"
	"log"

	"github.com/mibk/hawk/parse"
	"github.com/mibk/hawk/value"
)

type Tree struct {
	begin    []Stmt
	pActions []Stmt
	end      []Stmt

	vars map[string]value.Value
}

func NewTree() *Tree {
	return &Tree{vars: make(map[string]value.Value)}
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

	if p.pattern == nil || p.pattern.Eval().Bool() {
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
	e.expr.Eval()
}

type AssignStmt struct {
	tree *Tree
	name string
	expr Expr
}

func (a AssignStmt) Exec() {
	a.tree.vars[a.name] = a.expr.Eval()
}

type IfStmt struct {
	expr     Expr
	stmt     Stmt
	elseStmt Stmt
}

func (i IfStmt) Exec() {
	if i.expr.Eval().Bool() {
		i.stmt.Exec()
	} else if i.elseStmt != nil {
		i.elseStmt.Exec()
	}
}

type ForStmt struct {
	init Stmt
	cond Expr
	step Stmt
	body Stmt
}

func (f ForStmt) Exec() {
	if f.init != nil {
		f.init.Exec()
	}
	for f.cond == nil || f.cond.Eval().Bool() {
		f.body.Exec()
		if f.step != nil {
			f.step.Exec()
		}
	}
}

type CallStmt CallExpr

func (c CallStmt) Exec() {
	CallExpr(c).Eval()

}

type Expr interface {
	Eval() value.Value
}

type CallExpr struct {
	writer io.Writer
	fun    string
	args   []Expr
}

func (c CallExpr) Eval() value.Value {
	switch c.fun {
	case "print":
		var vals []interface{}
		for _, e := range c.args {
			vals = append(vals, e.Eval())
		}
		fmt.Fprintln(c.writer, vals...)
	default:
		log.Fatalf("unknown func %s", c.fun)
	}
	return value.NewBool(false)
}

type Ident struct {
	tree *Tree
	name string
}

func (i Ident) Eval() value.Value {
	return i.tree.vars[i.name]
}

type Col struct {
	p   *parse.Parser
	Num Expr
}

func (c Col) Eval() value.Value {
	n := c.Num.Eval().Int()
	return value.NewString(c.p.Field(n))
}

type BinaryExpr struct {
	typ   int
	left  Expr
	right Expr
}

const (
	ADD = iota
	SUB
	MUL
	DIV
)

func (e BinaryExpr) Eval() value.Value {
	switch e.typ {
	case ADD, SUB, MUL, DIV:
		l := e.left.Eval().Float64()
		r := e.right.Eval().Float64()
		var f float64
		switch e.typ {
		case ADD:
			f = l + r
		case SUB:
			f = l - r
		case MUL:
			f = l * r
		case DIV:
			f = l / r
		default:
			panic("unreachable")
		}
		return value.NewNumber(f)
	}
	switch e.typ {
	case LOR, LAND:
		lval := e.left.Eval()
		if e.typ == LOR {
			if lval.Bool() {
				return value.NewBool(true)
			}
			return value.NewBool(e.right.Eval().Bool())
		}
		if !lval.Bool() {
			return value.NewBool(false)
		}
		return value.NewBool(e.right.Eval().Bool())
	}
	cmp := e.left.Eval().Cmp(e.right.Eval())
	var b bool
	switch e.typ {
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
	case SUB:
	case MUL:
	case DIV:
	default:
		panic("unreachable")
	}
	return value.NewBool(b)
}

type UnaryExpr struct {
	typ  int
	expr Expr
}

func (e UnaryExpr) Eval() value.Value {
	switch e.typ {
	case SUB:
		return value.NewNumber(-e.expr.Eval().Float64())
	default:
		panic("unreachable")
	}
}

type Lit int

func (l Lit) Eval() value.Value {
	return value.NewNumber(float64(l))
}

type StringLit string

func (s StringLit) Eval() value.Value {
	return value.NewString(string(s))
}
