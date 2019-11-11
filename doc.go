//Package eval implements a math evaluator.
//for using eval is necessary to create an Evaluator with
//	e := eval.New()
//An evaluator contains the variables and the functions
//	e.SetVar("x", 42)
//For evaluating expressions use the Eval method
//	e.Eval("x+2")
package eval
