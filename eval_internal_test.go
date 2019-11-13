package meval

import (
	"go/ast"
	"go/token"
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

func TestFunctionsimpleWrapper(t *testing.T) {
	_, err := fns["ln"](1, 2, 3)
	if err == nil {
		t.Errorf("expected error; got nil")
	}
}
