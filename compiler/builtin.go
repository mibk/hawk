package compiler

import (
	"io"
	"log"
	"math"

	"github.com/mibk/hawk/value"
)

func evalArgs(w io.Writer, fname string, nargs int, args []Expr) []value.Value {
	if len(args) != nargs {
		// TODO: Get rid of log.Fatalf
		log.Fatalf("%s: %d != %d: argument count mismatch", fname, nargs, len(args))
	}
	vals := make([]value.Value, len(args))
	for i := range args {
		vals[i] = args[i].Eval(w)
	}
	return vals
}

func convertArgsToScalars(w io.Writer, fname string, nargs int, args []Expr) []*value.Scalar {
	if len(args) != nargs {
		// TODO: Get rid of log.Fatalf
		log.Fatalf("%s: %d != %d: argument count mismatch", fname, nargs, len(args))
	}
	vals := make([]*value.Scalar, len(args))
	for i := range args {
		if v, ok := args[i].Eval(w).Scalar(); ok {
			vals[i] = v
		} else {
			// TODO: Remove log.Fatalf and provide a better err message.
			log.Fatal("unimplemented operation")
		}
	}
	return vals
}

var aritFns = map[string]struct {
	narg int
	fn   func(io.Writer, []*value.Scalar) *value.Scalar
}{
	// Arithmetic functions:
	"atan2": {2, atan2},
	"cos":   {1, cos},
	"exp":   {1, exp},
	"log":   {1, _log},
	"sin":   {1, sin},
	"sqrt":  {1, sqrt},
}

func atan2(w io.Writer, vals []*value.Scalar) *value.Scalar {
	return value.NewNumber(math.Atan2(vals[0].Float64(), vals[1].Float64()))
}

func cos(w io.Writer, vals []*value.Scalar) *value.Scalar {
	return value.NewNumber(math.Cos(vals[0].Float64()))
}

func exp(w io.Writer, vals []*value.Scalar) *value.Scalar {
	return value.NewNumber(math.Exp(vals[0].Float64()))
}

func _log(w io.Writer, vals []*value.Scalar) *value.Scalar {
	return value.NewNumber(math.Log(vals[0].Float64()))
}

func sin(w io.Writer, vals []*value.Scalar) *value.Scalar {
	return value.NewNumber(math.Sin(vals[0].Float64()))
}

func sqrt(w io.Writer, vals []*value.Scalar) *value.Scalar {
	return value.NewNumber(math.Sqrt(vals[0].Float64()))
}
