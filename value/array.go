package value

import "bytes"

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
		if a.associative {
			// TODO: What should be the next autoincrement
			// value if the array is associative.
			panic("not implemented yet")
		}
		k = NewNumber(float64(a.ai))
		a.keys = append(a.keys, k)
		a.ai++
	} else {
		if _, ok := a.m[*k]; !ok {
			a.keys = append(a.keys, k)
			if !a.associative {
				a.ai++
				if k.typ != Number || k.number != float64(a.ai) {
					a.associative = true
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

func (a *Array) Cmp(Value) int { return -1 }
func (a *Array) String() string {
	var buf bytes.Buffer
	buf.WriteRune('[')
	// TODO: Print using a different method if the array is associative.
	for i, k := range a.keys {
		if i != 0 {
			buf.WriteString(", ")
		}
		// TODO: Use a different function for stringification of
		// the values. This method for example doesn't quote strings.
		buf.WriteString(a.m[*k].String())
	}
	buf.WriteRune(']')
	return buf.String()
}

func (a *Array) Len() int { return len(a.keys) }

// TODO
// func (v *Array) Format(s fmt.State, verb rune) {
// }
