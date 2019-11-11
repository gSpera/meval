package eval_test

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/gSpera/txt-graph/eval"
)

func TestEvaluator(t *testing.T) {
	e := eval.New()

	tm := []struct {
		name   string
		input  string
		err    error
		output float64
	}{
		{
			"simple",
			"2+2",
			nil,
			4,
		},
		{
			"order",
			"2+2*2",
			nil,
			6,
		},
		{
			"function call",
			"ln(2)",
			nil,
			math.Log(2),
		},
		{
			"operators",
			"(2+2)+(2-2)+(2*2)+(2/2)+(2%2)",
			nil,
			9,
		},

		{
			"cannot parse",
			"cannot parse, do not try",
			errors.New(`cannot parse "cannot parse, do not try": 1:8: expected 'EOF', found parse`),
			0,
		},
		{
			"unkown value",
			"2+a",
			errors.New(`cannot evaluate "2+a": cannot evaluate right operand: unkown variable: "a"`),
			0,
		},
		{
			"unkown funcation",
			"function(42)",
			errors.New(`cannot evaluate "function(42)": unkown function "function"`),
			0,
		},
		{
			"parsing left argument",
			"function(x+2)",
			errors.New(`cannot evaluate "function(x+2)": cannot evaluate argument 0 while calling "function": cannot evaluate left operand: unkown variable: "x"`),
			0,
		},
		{
			"parsing right argument",
			"function(2+x)",
			errors.New(`cannot evaluate "function(2+x)": cannot evaluate argument 0 while calling "function": cannot evaluate right operand: unkown variable: "x"`),
			0,
		},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			v, err := e.Eval(tt.input)
			if !equalError(err, tt.err) {
				t.Errorf("wrong error: expected: %v; got: %v", tt.err, err)
			}

			if v != tt.output {
				t.Errorf("wrong value: expected: %v; got: %v", tt.output, v)
			}
		})
	}
}

func equalError(a, b error) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil || b == nil && a != nil {
		return false
	}

	return a.Error() == b.Error()
}

func ExampleEvaluator() {
	e := eval.New()
	value, _ := e.Eval("2+2")
	fmt.Println(value)
	//Output:
	//4
}
func ExampleEvaluator_x() {
	e := eval.New()
	e.SetVar("x", 7)
	value, _ := e.Eval("x * x")
	fmt.Println(value)
	//Output:
	//49
}

func ExampleEvaluator_functionCall() {
	e := eval.New()
	value, _ := e.Eval("sqrt(49.0)")
	fmt.Println(value)
	//Output:
	//7
}

func ExampleEvaluator_customFunction() {
	e := eval.New()
	sum := func(args ...float64) (float64, error) {
		var sum float64
		for _, arg := range args {
			sum += arg
		}
		return sum, nil
	}
	e.SetFn("sum", sum)

	value, _ := e.Eval("sum(1,2,42)")
	fmt.Println(value)
	//Output:
	//45
}
