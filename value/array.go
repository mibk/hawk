package value

import (
	"bytes"
	"fmt"
)

// MergeArrays merges two arrays and returns a new array that will
// be a union of the previous two. If at least one of the arrays
// is associative, the union will contain all the keys from the
// first array and all the non-conflicting keys from the second
// one. Otherwise, all the values from both arrays are one after
// one put under keys from 0 to len(a)+len(a2)-1 and the resulting
// array is non-associative.
func MergeArrays(a, a2 *Array) *Array {
	z := NewArray()
	for _, k := range a.keys {
		z.Put(k, a.m[*k])
	}
	if !a.associative && !a2.associative {
		for _, k := range a2.keys {
			z.Put(nil, a2.m[*k])
		}
	} else {
		for _, k := range a2.keys {
			if _, ok := z.m[*k]; ok {
				continue
			}
			z.Put(k, a2.m[*k])
		}
	}
	return z
}

type Array struct {
	ai          int // autoincrement
	associative bool

	keys []*Scalar
	m    map[Scalar]Value
}

func NewArray() *Array {
	return &Array{m: make(map[Scalar]Value)}
}

// Put puts value v under key k into a. If k is nil,
// autoincrement value is used as a key.
func (a *Array) Put(k *Scalar, v Value) {
	if k == nil {
		k = NewNumber(float64(a.ai))
		a.keys = append(a.keys, k)
		a.ai++
	} else {
		if _, ok := a.m[*k]; !ok {
			a.keys = append(a.keys, k)
			if !a.associative {
				if k.typ != Number || k.number != float64(a.ai) {
					a.associative = true
				}
				if k.typ == Number {
					a.ai = int(k.number) + 1
				}
			}
		}
	}
	a.m[*k] = v
}

func (a *Array) Get(k *Scalar) Value {
	return a.m[*k]
}

func (a *Array) Keys() []*Scalar {
	return a.keys
}

func (a *Array) Scalar() (v *Scalar, ok bool) { return nil, false }
func (a *Array) Array() (v *Array, ok bool)   { return a, true }

func (a *Array) Cmp(v Value) (cmp int, ok bool) {
	// TODO: If v is Undefined, this way it becomes an array.
	// This might not be a desired behaviour.
	a2, ok := v.Array()
	if !ok {
		return -1, false
	}

	// Always return false as the second return value
	// as it's not possible to compare arrays using
	// <, >, <= or >=.

	if len(a.keys) != len(a2.keys) {
		return -1, false
	}
	if a.associative != a2.associative {
		return -1, false
	}
	for i, k := range a.keys {
		k2 := a2.keys[i]
		if cmp, _ := k.Cmp(k2); cmp != 0 {
			return -1, false
		}
		if cmp, _ := a.m[*k].Cmp(a2.m[*k2]); cmp != 0 {
			return -1, false
		}
	}
	return 0, false
}

func (a *Array) String() string {
	var buf bytes.Buffer
	buf.WriteRune('[')
	for i, k := range a.keys {
		if i != 0 {
			buf.WriteString(", ")
		}
		if a.associative {
			buf.WriteString(k.Encode() + ": ")
		}
		buf.WriteString(a.m[*k].Encode())
	}
	buf.WriteRune(']')
	return buf.String()
}

func (a *Array) Encode() string { return a.String() }
func (a *Array) Len() int       { return len(a.keys) }

func (a *Array) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		fmt.Fprint(s, a.String())
		return
	case 'V':
		fmt.Fprint(s, "array")
		return
	}
	fmt.Fprintf(s, formatVerb(s, verb), nil)
}
