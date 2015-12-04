package compiler

import (
	"fmt"
	"io"
	"log"

	"github.com/mibk/hawk/parse"
	"github.com/mibk/hawk/value"
)

type Expr interface {
	Eval() value.Value
}

type TernaryExpr struct {
	cond Expr
	yes  Expr
	no   Expr
}

func (t TernaryExpr) Eval() value.Value {
	if t.cond.Eval().Bool() {
		return t.yes.Eval()
	}
	return t.no.Eval()
}

type CallExpr struct {
	writer io.Writer
	fun    string
	args   []Expr
}

func (c CallExpr) Eval() value.Value {
	switch c.fun {
	case "add":
		// Add some ugly function just for the purpose to have one
		// as print cannot be used as a func.
		// TODO: remove this function.
		f := .0
		for _, e := range c.args {
			f += e.Eval().Float64()
		}
		return value.NewNumber(f)
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
	tree *Program
	name string
}

func (i Ident) Eval() value.Value {
	return i.tree.vars[i.name]
}

type FieldExpr struct {
	p   *parse.Parser
	num Expr
}

func (f FieldExpr) Eval() value.Value {
	n := f.num.Eval().Int()
	return value.NewString(f.p.Field(n))
}

type ExprOp int

const (
	_      ExprOp = iota
	Add           // left + right
	Sub           // left - right
	Mul           // left * right
	Div           // left / right
	Mod           // left % right
	OrOr          // left || right
	AndAnd        // left && right
	Eq            // left == right
	NotEq         // left != right
	Lt            // left < right
	LtEq          // left <= right
	Gt            // left > right
	GtEq          // left >= right

	Plus  // +expr
	Minus // -expr
	Not   // !expr
)

type BinaryExpr struct {
	op    ExprOp
	left  Expr
	right Expr
}

func (e BinaryExpr) Eval() value.Value {
	switch e.op {
	case Add, Sub, Mul, Div, Mod:
		l := e.left.Eval().Float64()
		r := e.right.Eval().Float64()
		var f float64
		switch e.op {
		case Add:
			f = l + r
		case Sub:
			f = l - r
		case Mul:
			f = l * r
		case Div:
			f = l / r
		case Mod:
			f = float64(int(l) % int(r))
		default:
			panic("unreachable")
		}
		return value.NewNumber(f)
	}
	switch e.op {
	case OrOr, AndAnd:
		lval := e.left.Eval()
		if e.op == OrOr {
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
	switch e.op {
	case Eq:
		b = cmp == 0
	case NotEq:
		b = cmp != 0
	case Lt:
		b = cmp == -1
	case LtEq:
		b = cmp <= 0
	case Gt:
		b = cmp == 1
	case GtEq:
		b = cmp >= 0
	default:
		panic("unreachable")
	}
	return value.NewBool(b)
}

type UnaryExpr struct {
	op   ExprOp
	expr Expr
}

func (e UnaryExpr) Eval() value.Value {
	switch e.op {
	case Minus:
		return value.NewNumber(-e.expr.Eval().Float64())
	case Not:
		return value.NewBool(!e.expr.Eval().Bool())
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
