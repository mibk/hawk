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

type Value struct {
	typ    int
	string string
	number float64
}

func NewNumber(f float64) *Value {
	return &Value{Number, "", f}
}

func NewString(s string) *Value {
	return &Value{String, s, 0}
}

func NewBool(b bool) *Value {
	n := .0
	if b {
		n = 1
	}
	return &Value{Bool, "", n}
}

func (v *Value) Cmp(b *Value) int {
	if v.typ == b.typ {
		return v.cmp(b)
	}
	return v.Number().cmp(b.Number())
}

func (v *Value) cmp(b *Value) int {
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
	panic("unreachable")
}

func (v *Value) Number() *Value {
	switch v.typ {
	case String:
		v.number, _ = strconv.ParseFloat(v.string, 64)
	}
	v.typ = Number
	return v
}

func (v *Value) Float64() float64 { return v.Number().number }
func (v *Value) Int() int         { return int(v.Number().number) }

func (v *Value) Bool() bool {
	return v.Cmp(NewBool(true)) == 0
}

func (v *Value) String() string {
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
	return "<unknown>"
}

func (v *Value) Format(s fmt.State, verb rune) {
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

func (v *Value) Len() int {
	if v.typ == String {
		return len(v.string)
	}
	// Handle other cases properly.
	return 0
}

func (z *Value) Add(x, y *Value) *Value {
	a, b := toFloat64(x, y)
	z.typ = Number
	z.number = a + b
	return z
}

func (z *Value) Sub(x, y *Value) *Value {
	a, b := toFloat64(x, y)
	z.typ = Number
	z.number = a - b
	return z
}

func (z *Value) Mul(x, y *Value) *Value {
	a, b := toFloat64(x, y)
	z.typ = Number
	z.number = a * b
	return z
}

func (z *Value) Div(x, y *Value) *Value {
	a, b := toFloat64(x, y)
	z.typ = Number
	z.number = a / b // TODO: division by 0.
	return z
}

func (z *Value) Mod(x, y *Value) *Value {
	a, b := toFloat64(x, y)
	z.typ = Number
	z.number = float64(int(a) % int(b))
	return z
}

func (z *Value) Neg(x *Value) *Value {
	z.typ = Number
	z.number = -x.Float64()
	return z
}

func toFloat64(x, y *Value) (float64, float64) {
	return x.Float64(), y.Float64()
}
