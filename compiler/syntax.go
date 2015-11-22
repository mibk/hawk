package compiler

import "github.com/mibk/hawk/parse"

type Expr interface {
	Val() Value
}

type Col struct {
	p   *parse.Parser
	Num Expr
}

func (c Col) Val() Value {
	n := c.Num.Val().Int()
	return NewStringValue(c.p.Field(n - 1))
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

type Tree struct {
	e Expr
}

func (t *Tree) Match() bool {
	return t.e.Val().Cmp(NewBoolValue(true)) == 0
}
