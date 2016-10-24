package compiler

import (
	"io"
	"math"

	"github.com/mibk/hawk/value"
)

func evalArgs(w io.Writer, fname string, nargs int, args []Expr) []value.Value {
	if len(args) != nargs {
		throw("%s: %d != %d: argument count mismatch", fname, nargs, len(args))
	}
	vals := make([]value.Value, len(args))
	for i := range args {
		vals[i] = args[i].Eval(w)
	}
	return vals
}

func convertArgsToScalars(w io.Writer, fname string, nargs int, args []Expr) []*value.Scalar {
	if len(args) != nargs {
		throw("%s: %d != %d: argument count mismatch", fname, nargs, len(args))
	}
	vals := make([]*value.Scalar, len(args))
	for i := range args {
		if v, ok := args[i].Eval(w).Scalar(); ok {
			vals[i] = v
		} else {
			throw("%s: all arguments must be scalar values", fname)
		}
	}
	return vals
}

var aritFns = map[string]struct {
	narg int
	fn   func([]*value.Scalar) *value.Scalar
}{
	// Arithmetic functions:
	"atan2": {2, atan2},
	"cos":   {1, cos},
	"exp":   {1, exp},
	"log":   {1, _log},
	"sin":   {1, sin},
	"sqrt":  {1, sqrt},
}

func atan2(vals []*value.Scalar) *value.Scalar {
	return value.NewNumber(math.Atan2(vals[0].Float64(), vals[1].Float64()))
}

func cos(vals []*value.Scalar) *value.Scalar {
	return value.NewNumber(math.Cos(vals[0].Float64()))
}

func exp(vals []*value.Scalar) *value.Scalar {
	return value.NewNumber(math.Exp(vals[0].Float64()))
}

func _log(vals []*value.Scalar) *value.Scalar {
	return value.NewNumber(math.Log(vals[0].Float64()))
}

func sin(vals []*value.Scalar) *value.Scalar {
	return value.NewNumber(math.Sin(vals[0].Float64()))
}

func sqrt(vals []*value.Scalar) *value.Scalar {
	return value.NewNumber(math.Sqrt(vals[0].Float64()))
}
