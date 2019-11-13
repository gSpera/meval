//Package eval implements a math evaluator.
//for using eval is necessary to create an Evaluator with
//	e := eval.New()
//An evaluator contains the variables and the functions
//	e.SetVar("x", 42)
//For evaluating expressions use the Eval method
//	e.Eval("x+2")
//
// Functions
//
//eval comes with some builtins functions, others may be added in future:
//Ln, Log10, Log2, Sin, Cos, Tan, Sqrt
//	e.Eval("ln(2)")
//
// Math Constants
//
//eval comes with some predefine math constants, others may be added in future:
//Pi, E
//
//	e.Eval("pi")
package eval
