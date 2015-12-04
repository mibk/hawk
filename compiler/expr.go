package compiler

import (
	"fmt"
	"io"
	"log"

	"github.com/mibk/hawk/parse"
	"github.com/mibk/hawk/value"
)

type Expr interface {
	Eval() *value.Value
}

type TernaryExpr struct {
	cond Expr
	yes  Expr
	no   Expr
}

func (t TernaryExpr) Eval() *value.Value {
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

func (c CallExpr) Eval() *value.Value {
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

func (i Ident) Eval() *value.Value {
	v, ok := i.tree.vars[i.name]
	if !ok {
		return new(value.Value)
	}
	return v
}

type FieldExpr struct {
	p   *parse.Parser
	num Expr
}

func (f FieldExpr) Eval() *value.Value {
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

func (e BinaryExpr) Eval() *value.Value {
	var z value.Value
	switch e.op {
	case Add, Sub, Mul, Div, Mod:
		l := e.left.Eval()
		r := e.right.Eval()
		switch e.op {
		case Add:
			z.Add(l, r)
		case Sub:
			z.Sub(l, r)
		case Mul:
			z.Mul(l, r)
		case Div:
			z.Div(l, r)
		case Mod:
			z.Mod(l, r)
		default:
			panic("unreachable")
		}
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
	default:
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
	return &z
}

type UnaryExpr struct {
	op   ExprOp
	expr Expr
}

func (e UnaryExpr) Eval() *value.Value {
	var z value.Value
	switch e.op {
	case Minus:
		z.Neg(e.expr.Eval())
	case Not:
		return value.NewBool(!e.expr.Eval().Bool())
	default:
		panic("unreachable")
	}
	return &z
}

type Lit int

func (l Lit) Eval() *value.Value {
	return value.NewNumber(float64(l))
}

type StringLit string

func (s StringLit) Eval() *value.Value {
	return value.NewString(string(s))
}
