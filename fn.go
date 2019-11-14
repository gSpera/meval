package meval

import (
	"fmt"
	"math"
)

//Function if a math function
type Function func(args ...float64) (v float64, err error)

var fns = map[string]func(args ...float64) (float64, error){
	"ln":    simpleWrapper(math.Log),
	"log10": simpleWrapper(math.Log10),
	"log2":  simpleWrapper(math.Log2),
	"log":   logFunction,
	"sin":   simpleWrapper(math.Sin),
	"cos":   simpleWrapper(math.Cos),
	"tan":   simpleWrapper(math.Tan),
	"sqrt":  simpleWrapper(math.Sqrt),
}

func simpleWrapper(f func(float64) float64) Function {
	return func(args ...float64) (float64, error) {
		if len(args) != 1 {
			return 0, fmt.Errorf("calling function with wrong number of argument: expected: 1; got: %d", len(args))
		}
		return f(args[0]), nil
	}
}

func logFunction(args ...float64) (float64, error) {
	if len(args) != 2 {
		return 0, fmt.Errorf("calling function with wrong number of argument: expected: 2; got: %d", len(args))
	}

	base := args[0]
	arg := args[1]

	//Use builtin functions
	switch base {
	case math.E:
		return math.Log(arg), nil
	case 10:
		return math.Log10(arg), nil
	case 2:
		return math.Log2(arg), nil
	}

	return math.Log(arg) / math.Log(base), nil
}
