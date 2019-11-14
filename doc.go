//Package Math Eval implements a math evaluator.
//for using eval is necessary to create an Evaluator with
//	e := meval.New()
//An evaluator contains the variables and the functions
//	e.SetVar("x", 42)
//For evaluating expressions use the Eval method
//	e.Eval("x+2")
// Valid operators are: +, -, *, /, %(remainder), ^(power)
//
// Functions
//
//eval comes with some builtins functions, others may be added in future:
//ln, log10, log2, log, sin, cos, tan, sqrt
//	e.Eval("ln(2)")
//
// Math Constants
//
//eval comes with some predefine math constants, others may be added in future:
//pi, e
//
//	e.Eval("pi")
package meval
