package value

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

const (
	String = iota
	Bool
	Number
)

type Value interface {
	Scalar() (v *Scalar, ok bool)
	Array() (a *Array, ok bool)
	Cmp(Value) int
	String() string
	Len() int

	// Encode encodes value to string in such a way that the resulting
	// string is a lexicographically correct representation of the
	// value.
	Encode() string
}

type Scalar struct {
	typ    int
	string string
	number float64
}

func NewNumber(f float64) *Scalar {
	return &Scalar{Number, "", f}
}

func NewString(s string) *Scalar {
	return &Scalar{String, s, 0}
}

func NewBool(b bool) *Scalar {
	n := .0
	if b {
		n = 1
	}
	return &Scalar{Bool, "", n}
}

func (v *Scalar) Scalar() (w *Scalar, ok bool) { return v, true }
func (v *Scalar) Array() (w *Array, ok bool)   { return nil, false }

func (v *Scalar) Cmp(w Value) int {
	v2, ok := w.Scalar()
	if !ok {
		// TODO: Fix case when the values are uncomparable.
		return -1
	}
	if v.typ == v2.typ {
		return v.cmp(v2)
	}
	return v.Number().cmp(v2.Number())
}

func (v *Scalar) cmp(b *Scalar) int {
	switch v.typ {
	case String:
		return strings.Compare(v.string, b.string)
	case Number, Bool:
		if v.number < b.number {
			return -1
		} else if v.number > b.number {
			return 1
		}
		return 0
	}
	panic("unknown scalar type")
}

func (v *Scalar) Number() *Scalar {
	switch v.typ {
	case String:
		v.number, _ = strconv.ParseFloat(v.string, 64)
	}
	v.typ = Number
	return v
}

func (v *Scalar) Float64() float64 { return v.Number().number }
func (v *Scalar) Int() int         { return int(v.Number().number) }

func (v *Scalar) Bool() bool {
	return v.Cmp(NewBool(true)) == 0
}

func (v *Scalar) String() string {
	switch v.typ {
	case String:
		return v.string
	case Number:
		return strconv.FormatFloat(v.number, 'f', -1, 64)
	case Bool:
		if v.number == 1 {
			return "true"
		}
		return "false"
	}
	panic("unknown scalar type")
}

func (v *Scalar) Encode() string {
	switch v.typ {
	case String:
		return strconv.Quote(v.string)
	default:
		return v.String()
	}
}

func (v *Scalar) Format(s fmt.State, verb rune) {
	var val interface{}
	switch verb {
	case 'v':
		fmt.Fprint(s, v.String())
		return
	// Boolean:
	case 't':
		val = v.Bool()
	// Integer:
	case 'b', 'c', 'd', 'o', 'U':
		// TODO: %b is different for integer and float.
		val = v.Int()
	// Floating-point:
	case 'e', 'E', 'f', 'F', 'g', 'G':
		val = v.Float64()
	// String:
	case 's':
		val = v.String()
	// Common for String and Integer
	case 'q', 'x', 'X':
		if v.typ == String {
			val = v.string
		} else {
			val = v.Int()
		}
	}
	fmt.Fprintf(s, formatVerb(s, verb), val)
}

func formatVerb(s fmt.State, verb rune) string {
	var buf bytes.Buffer
	buf.WriteRune('%')
	for _, c := range []int{' ', '0'} {
		if s.Flag(c) {
			buf.WriteRune(rune(c))
		}
	}
	if wid, ok := s.Width(); ok {
		fmt.Fprint(&buf, wid)
	}
	if prec, ok := s.Precision(); ok {
		fmt.Fprintf(&buf, ".%d", prec)
	}
	buf.WriteRune(verb)
	return buf.String()
}

func (v *Scalar) Len() int {
	if v.typ == String {
		return len(v.string)
	}
	// Handle other cases properly.
	return 0
}

func (z *Scalar) Add(x, y *Scalar) *Scalar {
	a, b := toFloat64(x, y)
	z.typ = Number
	z.number = a + b
	return z
}

func (z *Scalar) Sub(x, y *Scalar) *Scalar {
	a, b := toFloat64(x, y)
	z.typ = Number
	z.number = a - b
	return z
}

func (z *Scalar) Mul(x, y *Scalar) *Scalar {
	a, b := toFloat64(x, y)
	z.typ = Number
	z.number = a * b
	return z
}

func (z *Scalar) Div(x, y *Scalar) *Scalar {
	a, b := toFloat64(x, y)
	z.typ = Number
	z.number = a / b // TODO: division by 0.
	return z
}

func (z *Scalar) Mod(x, y *Scalar) *Scalar {
	a, b := toFloat64(x, y)
	z.typ = Number
	z.number = float64(int(a) % int(b))
	return z
}

func (z *Scalar) Neg(x *Scalar) *Scalar {
	z.typ = Number
	z.number = -x.Float64()
	return z
}

func toFloat64(x, y *Scalar) (float64, float64) {
	return x.Float64(), y.Float64()
}

func (z *Scalar) Concat(x, y *Scalar) *Scalar {
	z.typ = String
	z.string = x.String() + y.String()
	return z
}
