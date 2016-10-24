package value

import "fmt"

type Undefined struct {
	// TODO: Definition of this type might be a bit
	// unfortunate, but until a better solution is found,
	// leave it like that.
	//
	// The reason for this is that in the current design
	// it is not possible to change the underlying value
	// of the value.Value interface that is delivered by
	// evaluating an expression. At that point we have
	// no idea what scope that variable is from, nor do
	// we know the name of the variable. It's just a value.
	//
	// So effectively Undefined is just a sentinel value
	// for pontential future arrays. Scalar values are
	// immutable, at least for now, so we only care about
	// arrays at this point.
	arr *Array
}

func (u *Undefined) Scalar() (v *Scalar, ok bool) {
	if u.arr != nil {
		return nil, false
	}
	return &Scalar{}, true
}

func (u *Undefined) Array() (a *Array, ok bool) {
	if u.arr == nil {
		u.arr = NewArray()
	}
	return u.arr, true
}

func (u *Undefined) Cmp(v Value) int {
	if u.arr != nil {
		return u.arr.Cmp(v)
	}
	var eq bool
	switch v := v.(type) {
	case *Scalar:
		switch v.typ {
		case String:
			eq = v.string == ""
		case Bool, Number:
			eq = v.number == 0
		}
	case *Array:
		eq = v.Len() == 0
	}
	if eq {
		return 0
	}
	return -1
}

func (u *Undefined) String() string {
	if u.arr != nil {
		return u.arr.String()
	}
	return "undefined"
}

func (u *Undefined) Len() int {
	if u.arr != nil {
		return u.arr.Len()
	}
	return 0
}
func (u *Undefined) Encode() string {
	if u.arr != nil {
		return u.Encode()
	}
	return "undefined"
}

func (u *Undefined) Format(s fmt.State, verb rune) {
	if u.arr != nil {
		u.arr.Format(s, verb)
		return
	}
	switch verb {
	case 'v', 'V':
		fmt.Fprint(s, "undefined")
		return
	}
	fmt.Fprintf(s, formatVerb(s, verb), nil)
}
