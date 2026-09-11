package calc

import (
	"fmt"
	"prototype_kl/eval"
	"prototype_kl/parser"
)

// eval.go — expression evaluation pipeline: substitute (expand constructs
// and variables), degree/grade conversion, and parser evaluation.
func (c *Calculator) assign(name, expr string) error {
	if _, ok := parser.SupportedFunctions[name]; ok {
		return fmt.Errorf("cannot assign to function name %q", name)
	}
	if _, ok := parser.SupportedConstants[name]; ok {
		return fmt.Errorf("cannot assign to constant %q", name)
	}
	if name == "ans" {
		return fmt.Errorf("'ans' is reserved")
	}
	expanded, err := c.substitute(expr)
	if err != nil {
		return err
	}
	expr = expanded
	v, err := c.evalExpanded(expr)
	if err != nil {
		return err
	}
	c.vars[name] = v
	c.results = append(c.results, name+" = "+c.format(v))
	if !c.quietAssign {
		fmt.Fprintf(c.out, "%s = %s\n", name, c.format(v))
	}
	return nil
}

func (c *Calculator) eval(line string) (float64, error) {
	if !c.hasAns && hasIdent(line, "ans") {
		return 0, fmt.Errorf("no previous result yet")
	}
	expanded, err := c.substitute(line)
	if err != nil {
		return 0, err
	}
	return c.evalExpanded(expanded)
}

func (c *Calculator) evalExpanded(expanded string) (float64, error) {
	if expanded == breakSentinel {
		return 0, fmt.Errorf("break outside a loop")
	}
	if expanded == continueSentinel {
		return 0, fmt.Errorf("continue outside a loop")
	}
	if c.degMode {
		expanded = applyDeg(expanded)
	} else if c.gradMode {
		expanded = applyGrad(expanded)
	}
	return eval.Evaluate(expanded)
}

func (c *Calculator) substitute(expr string) (string, error) {
	// User-defined functions are expanded inline; variables, "ans", and "mem"
	// are substituted to their numeric literals.
	return c.expand(expr)
}
