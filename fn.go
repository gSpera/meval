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
