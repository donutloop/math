package eval

import (
	"prototype_kl/parser"
)

// Evaluate parses (via the parser package) and then evaluates an expression
// string. Returns the float64 result or a typed error if the expression is
// invalid.
func Evaluate(expression string) (float64, error) {
	ast, err := parser.Parse(expression)
	if err != nil {
		return 0, err
	}
	return NewEvaluator().Evaluate(ast)
}
