package meval

import (
	"fmt"
	"go/ast"
	"go/token"
	"math"
	"testing"
)

func TestEvaluatorevaluate(t *testing.T) {
	e := New()
	tm := []struct {
		name  string
		input ast.Expr
	}{
		{
			name: "call expression",
			input: &ast.CallExpr{
				Fun: &ast.BasicLit{},
			},
		},
		{
			name: "int parsing",
			input: &ast.BasicLit{
				Kind:  token.INT,
				Value: "not an int",
			},
		},
		{
			name: "float parsing",
			input: &ast.BasicLit{
				Kind:  token.FLOAT,
				Value: "not an int",
			},
		},
		{
			name: "complex parsing",
			input: &ast.BasicLit{
				Kind:  token.IMAG,
				Value: "not an int",
			},
		},
		{
			name: "string parsing",
			input: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "not an int",
			},
		},
		{
			name:  "unkown token",
			input: &ast.BadExpr{},
		},
		{
			name: "basic lit",
			input: &ast.BasicLit{
				Kind: token.RETURN,
			},
		},
	}

	for _, tt := range tm {
		t.Run(tt.name, func(t *testing.T) {
			_, err := e.evaluate(tt.input)
			if err == nil {
				t.Errorf("expected error; got nil")
			}
		})
	}
}

func TestEvaluatorBinaryExpr(t *testing.T) {
	e := New()
	_, err := e.binaryExpr(&ast.BinaryExpr{
		X:  &ast.BasicLit{Kind: token.INT, Value: "42"},
		Y:  &ast.BasicLit{Kind: token.INT, Value: "42"},
		Op: token.RETURN,
	})

	if err == nil {
		t.Errorf("expected error; got nil")
	}
}

func TestUnaryExpr(t *testing.T) {
	e := New()
	t.Run("invalid", func(t *testing.T) {
		_, err := e.unaryExpr(&ast.UnaryExpr{
			Op: token.INT,
		})
		if err == nil {
			t.Errorf("expected error; got nil")
		}
	})
}

func TestFunctionsimpleWrapper(t *testing.T) {
	_, err := fns["ln"](1, 2, 3)
	if err == nil {
		t.Errorf("expected error; got nil")
	}
}

func TestLogFunction(t *testing.T) {
	t.Run("arguments", func(t *testing.T) {
		_, err := logFunction(1, 2, 3)
		if err == nil {
			t.Errorf("expected error; got nil")
		}
	})

	tm := []struct {
		base float64
		arg  float64
		out  float64
	}{
		{10, 42, math.Log10(42)},
		{math.E, 42, math.Log(42)},
		{2, 42, math.Log2(42)},

		{3, 7, 1.7712437491614224},
	}

	for _, tt := range tm {
		t.Run(fmt.Sprintf("log_%g(%g) = %g", tt.base, tt.arg, tt.out), func(t *testing.T) {
			got, err := logFunction(tt.base, tt.arg)
			if err != nil {
				t.Errorf("got error: %v", err)
			}
			if got != tt.out {
				t.Errorf("expected: %g; got: %g; delta: %g", tt.out, got, tt.out-got)
			}
		})
	}
}

//BenchmarkLog calculates the best base for logarith to use in Log Function
func BenchmarkLog(b *testing.B) {
	logFunctions := []struct {
		name string
		fn   func(float64) float64
	}{
		{"e", math.Log},
		{"2", math.Log2},
		{"10", math.Log10},
	}

	sizes := []int{1, 2, 10, 100, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			arg := 3 << size
			base := 5 << size
			for _, f := range logFunctions {
				b.Run(f.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = f.fn(float64(arg)) / f.fn(float64(base))
					}
				})
			}
		})
	}
}
