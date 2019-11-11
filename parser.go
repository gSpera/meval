package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"math"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

var vars = map[string]float64{
	"x": 10,
}

func main() {
	expr, err := parser.ParseExpr("7+2*5-2+x+ln(x)")
	if err != nil {
		panic(err)
	}
	spew.Dump(expr)
	v, err := evaluate(expr)
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}

func evaluate(node ast.Node) (float64, error) {
	switch n := node.(type) {
	case *ast.BinaryExpr:
		return binaryExpr(n)
	case *ast.BasicLit:
		return basicLitToFloat(n)
	case *ast.Ident:
		v, ok := vars[n.Name]
		if !ok {
			return 0, fmt.Errorf("unkown variable: %q", n.Name)
		}
		return v, nil
	case *ast.CallExpr:
		ident, ok := n.Fun.(*ast.Ident)
		if !ok {
			return 0, fmt.Errorf("call expression function name not an ident: %T(%v)", n, n)
		}
		fn := ident.Name
		args := make([]float64, len(n.Args))
		for i, arg := range n.Args {
			v, err := evaluate(arg)
			if err != nil {
				return 0, fmt.Errorf("cannot evaluate argument %d while calling %q: %w", i, fn, err)
			}
			args[i] = v
		}
		f, ok := fns[fn]
		if !ok {
			return 0, fmt.Errorf("unkown function %q", fn)
		}
		return f(args...)
	}

	fmt.Fprintf(os.Stderr, "Token: %T(%+v)\n", node, node)
	panic("Unkown Token")
}

func binaryExpr(expr *ast.BinaryExpr) (float64, error) {
	x, err := evaluate(expr.X)
	if err != nil {
		return 0, fmt.Errorf("cannot  evaluate left operand: %w", err)
	}
	y, err := evaluate(expr.Y)
	if err != nil {
		return 0, fmt.Errorf("cannot evaluate right operand: %w", err)
	}

	var v float64
	switch expr.Op {
	case token.ADD:
		v = x + y
	case token.SUB:
		v = x - y
	case token.MUL:
		v = x * y
	case token.QUO:
		v = x / y
	case token.REM:
		v = math.Remainder(x, y)

	default:
		panic("Unkown Token")
	}

	return v, nil
}

func basicLitToFloat(b *ast.BasicLit) (float64, error) {
	switch b.Kind {
	case token.INT:
		v, err := strconv.Atoi(b.Value)
		if err != nil {
			return 0, fmt.Errorf("cannot parse int literal: %w", err)
		}
		return float64(v), nil
	case token.FLOAT:
		v, err := strconv.ParseFloat(b.Value, 64)
		if err != nil {
			return 0, fmt.Errorf("cannot parse float literal: %w", err)
		}
		return v, nil
	case token.IMAG:
		return 0, fmt.Errorf("imaginary/complex numbers are not supported")
	case token.CHAR, token.STRING:
		return 0, fmt.Errorf("found string literal")
	}

	return 0, fmt.Errorf("BasicLit: %T(%+v)", b, b)
}
