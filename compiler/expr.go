package compiler

import (
	"fmt"
	"io"
	"log"

	"github.com/mibk/hawk/scan"
	"github.com/mibk/hawk/value"
)

type Expr interface {
	Eval(io.Writer) value.Value
}

type TernaryExpr struct {
	cond Expr
	yes  Expr
	no   Expr
}

func (t *TernaryExpr) Eval(w io.Writer) value.Value {
	v, ok := t.cond.Eval(w).Scalar()
	if !ok {
		// TODO: Remove log.Fatal
		log.Fatal("invalid operation")
	}
	if v.Bool() {
		return t.yes.Eval(w)
	}
	return t.no.Eval(w)
}

type CallExpr struct {
	fun  string
	args []Expr
}

func (c *CallExpr) Eval(w io.Writer) value.Value {
	switch c.fun {
	case "len":
		vals := evalArgs(w, c.fun, 1, c.args)
		return value.NewNumber(float64(vals[0].Len()))
	case "sprintf":
		format, vals, err := formatPrintfArgs(w, "sprintf", c.args)
		if err != nil {
			// TODO: Get rid of log.Fatal.
			log.Fatal(err)
		}
		return value.NewString(fmt.Sprintf(format, vals...))
	}

	// Arithmetic functions:
	if dcl, ok := aritFns[c.fun]; ok {
		vals := evalArgs(w, c.fun, dcl.narg, c.args)
		return dcl.fn(w, vals)
	}

	// TODO: Get rid of log.Fatalf
	fn, ok := ast.funcs[c.fun]
	if !ok {
		log.Fatalf("unknown func %s", c.fun)
	}
	vals := evalArgs(w, c.fun, len(fn.args), c.args)
	fn.scope.Push()
	defer fn.scope.Pull()
	for i, n := range fn.args {
		fn.scope.SetVar(n, vals[i])
	}
	fn.body.Exec(w)
	if ast.retval != nil {
		v := ast.retval
		ast.retval = nil
		return v
	}
	return value.NewBool(false)
}

type Ident struct {
	scope Scope
	name  string
}

func (i *Ident) Eval(io.Writer) value.Value {
	return i.scope.Var(i.name)
}

type FieldExpr struct {
	sc  *scan.Scanner
	num Expr
}

func (f *FieldExpr) Eval(w io.Writer) value.Value {
	v, ok := f.num.Eval(w).Scalar()
	if !ok {
		// TODO: Remove log.Fatalf
		log.Fatal("invalid operation")
	}
	return value.NewString(f.sc.Field(v.Int()))
}

type IndexExpr struct {
	expr  Expr
	index Expr
}

func (ie *IndexExpr) Eval(w io.Writer) value.Value {
	a, ok := ie.expr.Eval(w).Array()
	if !ok {
		// TODO: Remove log.Fatalf
		log.Fatal("invalid operation; need array")
	}
	index, ok := ie.index.Eval(w).Scalar()
	if !ok {
		// TODO: Remove log.Fatalf
		log.Fatal("invalid operation")
	}
	v := a.Get(index)
	if v == nil {
		// TODO: Return a nil value?
		return value.NewBool(false)
	}
	return v
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

func (e *BinaryExpr) Eval(w io.Writer) value.Value {
	var z value.Scalar
	switch e.op {
	case Add, Sub, Mul, Div, Mod:
		l, ok := e.left.Eval(w).Scalar()
		r, ok2 := e.right.Eval(w).Scalar()
		if !ok && !ok2 {
			// TODO: Remove log.Fatal
			log.Fatal("invalid operation")
		}
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
		lval, ok := e.left.Eval(w).Scalar()
		if !ok {
			// TODO: Remove log.Fatal
			log.Fatal("invalid operation")
		}

		if e.op == OrOr {
			if lval.Bool() {
				return value.NewBool(true)
			}
			rval, ok := e.right.Eval(w).Scalar()
			if !ok {
				// TODO: Remove log.Fatal
				log.Fatal("invalid operation")
			}
			return value.NewBool(rval.Bool())
		}
		if !lval.Bool() {
			return value.NewBool(false)
		}
		rval, ok := e.right.Eval(w).Scalar()
		if !ok {
			// TODO: Remove log.Fatal
			log.Fatal("invalid operation")
		}
		return value.NewBool(rval.Bool())
	default:
		cmp := e.left.Eval(w).Cmp(e.right.Eval(w))
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

func (e *UnaryExpr) Eval(w io.Writer) value.Value {
	v, ok := e.expr.Eval(w).Scalar()
	if !ok {
		// TODO: Remove log.Fatal
		log.Fatal("invalid operation")
	}
	var z value.Scalar
	switch e.op {
	case Minus:
		z.Neg(v)
	case Not:
		return value.NewBool(!v.Bool())
	default:
		panic("unreachable")
	}
	return &z
}

type Lit int

func (l Lit) Eval(io.Writer) value.Value {
	return value.NewNumber(float64(l))
}

type StringLit string

func (s StringLit) Eval(io.Writer) value.Value {
	return value.NewString(string(s))
}

type ArrayLit struct {
	exprs []Expr
}

func (al *ArrayLit) Eval(w io.Writer) value.Value {
	arr := value.NewArray()
	for _, e := range al.exprs {
		arr.Put(nil, e.Eval(w))
	}
	return arr
}
