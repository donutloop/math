package calc

import (
	"fmt"
			"strings"

	"prototype_kl/parser"
)

// expandRangeLoop rewrites a generalized range loop into a flat arithmetic
// expression by substituting the loop variable with each integer in the range.
//
//		sum(i, lo, hi, expr)         -> 0 + (expr@lo) + (expr@lo+1) + ...
//		sum(i, lo, hi, step, expr)   -> same, stepping by step
//		prod(i, lo, hi, expr)        -> 1 * (expr@lo) * (expr@lo+1) * ...
//		count(i, lo, hi, cond)       -> 0 + (cond@lo) + (cond@lo+1) + ...
//
// count accumulates the condition values directly, since comparisons already
// yield 1 (true) or 0 (false). This is a macro-level "for loop" over integers
// with a running variable, generalizing the plain numeric sum/prod/count forms.
func (c *Calculator) expandRangeLoop(inner, kind string) (string, error) {
	args := splitArgs(inner)
	if len(args) != 4 && len(args) != 5 {
		return "", fmt.Errorf("%s expects 4 or 5 arguments (var, lo, hi[, step], body)", kind)
	}
	varName := strings.TrimSpace(args[0])
	if !isIdent(varName) {
		return "", fmt.Errorf("%s loop variable %q is not a valid identifier", kind, varName)
	}
	loArg := strings.TrimSpace(args[1])
	hiArg := strings.TrimSpace(args[2])
	stepArg := "1"
	bodyArg := args[3]
	if len(args) == 5 {
		stepArg = strings.TrimSpace(args[3])
		bodyArg = args[4]
	}
	loV, err := c.expand(loArg)
	if err != nil {
		return "", err
	}
	hiV, err := c.expand(hiArg)
	if err != nil {
		return "", err
	}
	stepV, err := c.expand(stepArg)
	if err != nil {
		return "", err
	}
	loNum, err := parser.Evaluate(loV)
	if err != nil {
		return "", fmt.Errorf("%s lower bound must be numeric", kind)
	}
	hiNum, err := parser.Evaluate(hiV)
	if err != nil {
		return "", fmt.Errorf("%s upper bound must be numeric", kind)
	}
	stepNum, err := parser.Evaluate(stepV)
	if err != nil {
		return "", fmt.Errorf("%s step must be numeric", kind)
	}
	loN := int(loNum)
	hiN := int(hiNum)
	stepN := int(stepNum)
	if stepN < 1 {
		return "", fmt.Errorf("%s step must be >= 1", kind)
	}
	terms := []string{}
	for i := loN; i <= hiN; i += stepN {
		term := replaceIdent(bodyArg, varName, fmt.Sprintf("%d", i))
		te, err := c.expand(term)
		if err != nil {
			return "", err
		}
		terms = append(terms, "("+te+")")
	}
	switch kind {
	case "count":
		// count sums the body values (comparisons yield 1/0), so it is the sum.
		if len(terms) == 0 {
			return "0", nil
		}
		return strings.Join(terms, "+"), nil
	case "sum":
		if len(terms) == 0 {
			return "0", nil
		}
		return strings.Join(terms, "+"), nil
	case "prod":
		if len(terms) == 0 {
			return "1", nil
		}
		return strings.Join(terms, "*"), nil
	case "avg":
		if len(terms) == 0 {
			return "", fmt.Errorf("avg over an empty range")
		}
		return "(" + strings.Join(terms, "+") + ")/" + fmt.Sprintf("%d", len(terms)), nil
	case "min":
		if len(terms) == 0 {
			return "", fmt.Errorf("min over an empty range")
		}
		return "min(" + strings.Join(terms, ",") + ")", nil
	case "max":
		if len(terms) == 0 {
			return "", fmt.Errorf("max over an empty range")
		}
		return "max(" + strings.Join(terms, ",") + ")", nil
	}
	return "", fmt.Errorf("unknown range kind %q", kind)
}

