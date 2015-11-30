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
	tree *Root
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
	MOD

	NOT
)

func (e BinaryExpr) Eval() value.Value {
	switch e.typ {
	case ADD, SUB, MUL, DIV, MOD:
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
		case MOD:
			f = float64(int(l) % int(r))
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
	case NOT:
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
