package hawkc

import (
	"fmt"
	"io"
	"regexp"

	"github.com/mibk/hawk/scan"
	"github.com/mibk/hawk/value"
)

type Expr interface {
	Eval(io.Writer) value.Value
}

type TernaryExpr struct {
	debugInfo
	Cond Expr
	Yes  Expr
	No   Expr
}

func (t *TernaryExpr) Eval(w io.Writer) value.Value {
	v, ok := t.Cond.Eval(w).Scalar()
	if !ok {
		t.throw("non-scalar value used as a condition")
	}
	if v.Bool() {
		return t.Yes.Eval(w)
	}
	return t.No.Eval(w)
}

type CallExpr struct {
	debugInfo
	Fun  string
	Args []Expr
}

func (c *CallExpr) Eval(w io.Writer) value.Value {
	switch c.Fun {
	case "len":
		vals := evalArgs(c.debugInfo, w, c.Fun, 1, c.Args)
		return value.NewNumber(float64(vals[0].Len()))
	case "sprintf":
		format, vals, err := formatPrintfArgs(w, "sprintf", c.Args)
		if err != nil {
			c.throw("%v", err)
		}
		return value.NewString(fmt.Sprintf(format, vals...))
	}

	// Arithmetic functions:
	if dcl, ok := aritFns[c.Fun]; ok {
		vals := convertArgsToScalars(c.debugInfo, w, c.Fun, dcl.narg, c.Args)
		return dcl.fn(vals)
	}

	fn, ok := ast.funcs[c.Fun]
	if !ok {
		c.throw("unknown function: %s", c.Fun)
	}
	vals := convertArgsToScalars(c.debugInfo, w, c.Fun, len(fn.Args), c.Args)
	fn.scope.Push()
	defer fn.scope.Pull()
	for i, n := range fn.Args {
		fn.scope.Put(n, vals[i])
	}
	fn.Body.Exec(w)
	if ast.retval != nil {
		v := ast.retval
		ast.retval = nil
		return v
	}
	return value.NewBool(false)
}

type Ident struct {
	scope Scope
	Name  string
}

func (i *Ident) Eval(io.Writer) value.Value {
	return i.scope.Get(i.Name)
}

type FieldExpr struct {
	debugInfo
	sc *scan.Scanner
	X  Expr
}

func (f *FieldExpr) Eval(w io.Writer) value.Value {
	v, ok := f.X.Eval(w).Scalar()
	if !ok {
		f.throw("attempting to access a field using a non-scalar value")
	}
	i := v.Int()
	if i < 0 {
		f.throw("attempting to access a field using a negative index")
	}
	return value.NewString(f.sc.Field(i))
}

type IndexExpr struct {
	debugInfo
	X     Expr
	Index Expr
}

func (ie *IndexExpr) Eval(w io.Writer) value.Value {
	a, ok := ie.X.Eval(w).Array()
	if !ok {
		// TODO: This might be permitted e.g. for string.
		ie.throw("attempting to get an index of a scalar value")
	}
	index, ok := ie.Index.Eval(w).Scalar()
	if !ok {
		ie.throw("indexing an array using a non-scalar value")
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
	Add           // x + y
	Sub           // x - y
	Mul           // x * y
	Div           // x / y
	Mod           // x % y
	OrOr          // x || y
	AndAnd        // x && y
	Eq            // x == y
	NotEq         // x != y
	Lt            // x < y
	LtEq          // x <= y
	Gt            // x > y
	GtEq          // x >= y

	Plus  // +expr
	Minus // -expr
	Not   // !expr

	Concat // x . y
)

type BinaryExpr struct {
	debugInfo
	Op ExprOp
	X  Expr
	Y  Expr
}

func (e *BinaryExpr) Eval(w io.Writer) value.Value {
	var z value.Scalar
	switch e.Op {
	case Add, Sub, Mul, Div, Mod, Concat:
		v, v2 := e.X.Eval(w), e.Y.Eval(w)
		l, ok := v.Scalar()
		r, ok2 := v2.Scalar()
		if !ok || !ok2 {
			if e.Op == Add {
				a, ok := v.Array()
				a2, ok2 := v2.Array()
				if ok && ok2 {
					return value.MergeArrays(a, a2)
				}
			}
			e.throw("unsupported types for binary expression: %V and %V", v, v2)
		}
		switch e.Op {
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
		case Concat:
			z.Concat(l, r)
		default:
			panic("unreachable")
		}
	case OrOr, AndAnd:
		lval, ok := e.X.Eval(w).Scalar()
		if !ok {
			e.throw("unsupported type for binary expression")
		}

		if e.Op == OrOr {
			if lval.Bool() {
				return value.NewBool(true)
			}
			rval, ok := e.Y.Eval(w).Scalar()
			if !ok {
				e.throw("unsupported type for binary expression")
			}
			return value.NewBool(rval.Bool())
		}
		if !lval.Bool() {
			return value.NewBool(false)
		}
		rval, ok := e.Y.Eval(w).Scalar()
		if !ok {
			e.throw("unsupported type for binary expression")
		}
		return value.NewBool(rval.Bool())
	default:
		l, r := e.X.Eval(w), e.Y.Eval(w)
		cmp, ok := l.Cmp(r)
		if !ok && e.Op != Eq && e.Op != NotEq {
			e.throw("cannot compare %V and %V using <, >, <=, or >=", l, r)
		}
		var b bool
		switch e.Op {
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
			panic("unknown binary operation")
		}
		return value.NewBool(b)
	}
	return &z
}

type UnaryExpr struct {
	debugInfo
	Op ExprOp
	X  Expr
}

func (e *UnaryExpr) Eval(w io.Writer) value.Value {
	v, ok := e.X.Eval(w).Scalar()
	if !ok {
		e.throw("unsupported type for unary expression")
	}
	var z value.Scalar
	switch e.Op {
	case Minus:
		z.Neg(v)
	case Not:
		return value.NewBool(!v.Bool())
	default:
		panic("unknown unary operation")
	}
	return &z
}

type MatchExpr struct {
	debugInfo
	X     Expr
	Y     Expr
	Match bool
}

func (m *MatchExpr) Eval(w io.Writer) value.Value {
	l, r := m.X.Eval(w), m.Y.Eval(w)
	x, ok := l.Scalar()
	y, ok2 := r.Scalar()
	if !ok || !ok2 || y.Type() != value.String {
		m.throw("invalid types for regexp matching: %V ~ %V", l, r)
	}
	rx, err := regexp.Compile(y.String())
	if err != nil {
		m.throw("invalid regexp")
	}
	return value.NewBool(rx.MatchString(x.String()) == m.Match)
}

type BasicLit struct {
	Val value.Value
}

func (b BasicLit) Eval(io.Writer) value.Value { return b.Val }

type ArrayLit struct {
	Elems []Expr
}

func (al *ArrayLit) Eval(w io.Writer) value.Value {
	arr := value.NewArray()
	for _, e := range al.Elems {
		arr.Put(nil, e.Eval(w))
	}
	return arr
}
