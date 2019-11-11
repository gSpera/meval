package parser

import (
	"fmt"
	"math"
)

var fns = map[string]func(args ...float64) (float64, error){
	"ln": func(args ...float64) (float64, error) {
		if len(args) != 1 {
			return 0, fmt.Errorf("calling ln with wrong number of argument: expected: 1; got: %d", len(args))
		}
		return math.Log(args[0]), nil
	},
}
