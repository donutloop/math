package calc

import (
	"fmt"
	"math"
	"prototype_kl/calc/control"
	"prototype_kl/eval"
	"strconv"
	"strings"
)

// control.BreakSentinel and control.ContinueSentinel are internal signals that expandBegin
// raises when a begin(...) block encounters a break/continue statement. The
// enclosing loop (expandWhile) catches them at the string-expansion level
// before the value is parsed, so break/continue work inside begin blocks.

// expandRangeLoop rewrites a generalized range loop into a flat arithmetic
// expression by substituting the loop variable with each integer in the range.
//
//	sum(i, lo, hi, expr)         -> 0 + (expr@lo) + (expr@lo+1) + ...
//	sum(i, lo, hi, step, expr)   -> same, stepping by step
//	prod(i, lo, hi, expr)        -> 1 * (expr@lo) * (expr@lo+1) * ...
//	count(i, lo, hi, cond)       -> 0 + (cond@lo) + (cond@lo+1) + ...
//
// count accumulates the condition values directly, since comparisons already
// yield 1 (true) or 0 (false). This is a macro-level "for loop" over integers
// with a running variable, generalizing the plain numeric sum/prod/count forms.
func (c *Calculator) expandRangeLoop(inner, kind string) (string, error) {
	args := control.SplitArgs(inner)
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
	loNum, err := eval.Evaluate(loV)
	if err != nil {
		return "", fmt.Errorf("%s lower bound must be numeric", kind)
	}
	hiNum, err := eval.Evaluate(hiV)
	if err != nil {
		return "", fmt.Errorf("%s upper bound must be numeric", kind)
	}
	stepNum, err := eval.Evaluate(stepV)
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
		term := control.ReplaceIdent(bodyArg, varName, fmt.Sprintf("%d", i))
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
	case "gcd":
		if len(terms) == 0 {
			return "", fmt.Errorf("gcd over an empty range")
		}
		return "gcd(" + strings.Join(terms, ",") + ")", nil
	case "lcm":
		if len(terms) == 0 {
			return "", fmt.Errorf("lcm over an empty range")
		}
		return "lcm(" + strings.Join(terms, ",") + ")", nil
	}
	return "", fmt.Errorf("unknown range kind %q", kind)
}

// expandRepeat expands a repeat(n, var, body) macro: it evaluates body with
// var bound to each integer 1..n and returns the LAST evaluated value.
func (c *Calculator) expandRepeat(countS, vname, body string) (string, error) {
	count, err := c.eval(countS)
	if err != nil {
		return "", err
	}
	last := "0"
	for i := 1; i <= int(count); i++ {
		for _, st := range splitStatements(body) {
			b := control.ReplaceIdent(st, vname, strconv.Itoa(i))
			t := strings.TrimSpace(b)
			if t == "break" {
				return last, nil
			}
			if expr, ok := control.BreakValue(t); ok {
				v, err := c.eval(expr)
				if err != nil {
					return "", err
				}
				return fmt.Sprintf("%g", v), nil
			}
			if t == "continue" {
				break
			}
			if name, expr, ok := parseAssignment(b); ok {
				if err := c.assign(name, expr); err != nil {
					return "", err
				}
				last = fmt.Sprintf("%g", c.vars[name])
			} else {
				// Expand at string level so begin(...) blocks can raise
				// break/continue sentinels that are caught before parsing.
				expanded, err := c.substitute(b)
				if err != nil {
					return "", err
				}
				if expanded == control.BreakSentinel {
					return last, nil
				}
				if strings.HasPrefix(expanded, control.BreakValuePrefix) {
					v, err := c.eval(expanded[len(control.BreakValuePrefix):])
					if err != nil {
						return "", err
					}
					return fmt.Sprintf("%g", v), nil
				}
				if expanded == control.ContinueSentinel {
					break
				}
				v, err := c.evalExpanded(expanded)
				if err != nil {
					return "", err
				}
				last = fmt.Sprintf("%g", v)
			}
		}
	}
	return last, nil
}

// expandWhile implements a while(cond, body) loop — a general control-flow
// construct for the math programming language. cond is evaluated with the
// current variable bindings; while it is nonzero, body is evaluated as a
// statement (so an assignment like x = x + 1 mutates variables between
// iterations). A hard iteration bound prevents runaway/infinite programs.
// It returns the value of the last body evaluation (or the loop variable's
// final value when body is an assignment).
func (c *Calculator) expandWhile(cond, body string) (string, error) {
	const maxIters = 10000
	last := "0"
	iters := 0
	for {
		if iters >= maxIters {
			return "", fmt.Errorf("while loop exceeded %d iterations", maxIters)
		}
		iters++
		cv, err := c.eval(cond)
		if err != nil {
			return "", err
		}
		if cv == 0 {
			return last, nil
		}
		// Evaluate body as a sequence of statements (paren-aware split on ';'),
		// so multi-step bodies like acc = acc + x; x = x + 1 mutate variables.
		// Direct 'break' exits the loop; 'continue' skips the rest of the body.
		last = "0"
		for _, st := range splitStatements(body) {
			t := strings.TrimSpace(st)
			if t == "break" {
				return last, nil
			}
			if expr, ok := control.BreakValue(t); ok {
				v, err := c.eval(expr)
				if err != nil {
					return "", err
				}
				return fmt.Sprintf("%g", v), nil
			}
			if t == "continue" {
				break
			}
			if name, expr, ok := parseAssignment(st); ok {
				if err := c.assign(name, expr); err != nil {
					return "", err
				}
				last = fmt.Sprintf("%g", c.vars[name])
			} else {
				// Expand at string level so begin(...) blocks can raise
				// break/continue sentinels that are caught before parsing.
				expanded, err := c.substitute(st)
				if err != nil {
					return "", err
				}
				if expanded == control.BreakSentinel {
					return last, nil
				}
				if strings.HasPrefix(expanded, control.BreakValuePrefix) {
					v, err := c.eval(expanded[len(control.BreakValuePrefix):])
					if err != nil {
						return "", err
					}
					return fmt.Sprintf("%g", v), nil
				}
				if expanded == control.ContinueSentinel {
					break
				}
				bv, err := c.evalExpanded(expanded)
				if err != nil {
					return "", err
				}
				last = fmt.Sprintf("%g", bv)
			}
		}
	}
}

// expandBegin implements a begin(...) statement block: the statements inside
// are evaluated in sequence (paren-aware split on ';'), each as an assignment
// or expression, so they can mutate variables. It returns the last value.
func (c *Calculator) expandBegin(body string) (string, error) {
	last := "0"
	for _, st := range splitStatements(body) {
		t := strings.TrimSpace(st)
		if t == "break" {
			return control.BreakSentinel, nil
		}
		if expr, ok := control.BreakValue(t); ok {
			return control.BreakValuePrefix + expr, nil
		}
		if t == "continue" {
			return control.ContinueSentinel, nil
		}
		if name, expr, ok := parseAssignment(st); ok {
			if err := c.assign(name, expr); err != nil {
				return "", err
			}
			last = fmt.Sprintf("%g", c.vars[name])
		} else {
			// Expand at string level so nested break/continue (e.g. from an
			// if branch) can propagate up to the enclosing loop.
			expanded, err := c.substitute(st)
			if err != nil {
				return "", err
			}
			if expanded == control.BreakSentinel || expanded == control.ContinueSentinel ||
				strings.HasPrefix(expanded, control.BreakValuePrefix) {
				return expanded, nil
			}
			bv, err := c.evalExpanded(expanded)
			if err != nil {
				return "", err
			}
			last = fmt.Sprintf("%g", bv)
		}
	}
	return last, nil
}

// expandFor implements for(var, lo, hi, body): bind var to each integer in
// [lo, hi] inclusive and evaluate body (substituting var), returning the last
// body value. lo/hi are evaluated first so the range is fixed up front.
func (c *Calculator) expandFor(vname, loS, hiS, stepS, body string) (string, error) {
	lo, err := c.eval(loS)
	if err != nil {
		return "", err
	}
	hi, err := c.eval(hiS)
	if err != nil {
		return "", err
	}
	last := "0"
	step, err := c.eval(stepS)
	if err != nil {
		return "", err
	}
	if step == 0 {
		return "", fmt.Errorf("for step must be nonzero")
	}
	for i := int(lo); i <= int(hi); i += int(math.Abs(step)) {
		for _, st := range splitStatements(body) {
			b := control.ReplaceIdent(st, vname, strconv.Itoa(i))
			t := strings.TrimSpace(b)
			if t == "break" {
				return last, nil
			}
			if expr, ok := control.BreakValue(t); ok {
				v, err := c.eval(expr)
				if err != nil {
					return "", err
				}
				return fmt.Sprintf("%g", v), nil
			}
			if t == "continue" {
				break
			}
			if name, expr, ok := parseAssignment(b); ok {
				if err := c.assign(name, expr); err != nil {
					return "", err
				}
				last = fmt.Sprintf("%g", c.vars[name])
			} else {
				// Expand at string level so begin(...) blocks can raise
				// break/continue sentinels that are caught before parsing.
				expanded, err := c.substitute(b)
				if err != nil {
					return "", err
				}
				if expanded == control.BreakSentinel {
					return last, nil
				}
				if strings.HasPrefix(expanded, control.BreakValuePrefix) {
					v, err := c.eval(expanded[len(control.BreakValuePrefix):])
					if err != nil {
						return "", err
					}
					return fmt.Sprintf("%g", v), nil
				}
				if expanded == control.ContinueSentinel {
					break
				}
				v, err := c.evalExpanded(expanded)
				if err != nil {
					return "", err
				}
				last = fmt.Sprintf("%g", v)
			}
		}
	}
	return last, nil
}
