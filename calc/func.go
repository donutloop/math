package calc

import (
	"fmt"
	"prototype_kl/calc/control"
	"prototype_kl/eval"
	"prototype_kl/lexer"
	"strconv"
	"strings"

	"prototype_kl/parser"
)

// funcDef stores a user-defined function: its parameter names and body
// expression (as source text).
type funcDef struct {
	Params []string `json:"params"`
	Body   string   `json:"body"`
}

// parseFuncDef parses a function definition of the form
//
//	name(param1, param2) = body-expression
//
// It returns ok=false when the line is not a function definition (e.g. a
// bare variable assignment or an expression).
func parseFuncDef(line string) (name string, params []string, body string, ok bool) {
	i, n := 0, len(line)

	// First token must be an identifier (the function name).
	if i >= n || !lexer.IsIdentStart(line[i]) {
		return "", nil, "", false
	}
	j := i + 1
	for j < n && lexer.IsIdentChar(line[j]) {
		j++
	}
	name = line[i:j]
	i = j

	// Skip spaces, then require '('.
	for i < n && line[i] == ' ' {
		i++
	}
	if i >= n || line[i] != '(' {
		return "", nil, "", false
	}

	// Find matching ')' for the parameter list.
	end := findMatchingParen(line, i)
	if end < 0 {
		return "", nil, "", false
	}
	inner := line[i+1 : end]
	params, ok = parseParams(inner)
	if !ok {
		return "", nil, "", false
	}

	i = end + 1
	for i < n && line[i] == ' ' {
		i++
	}
	if i >= n || line[i] != '=' {
		return "", nil, "", false
	}

	body = strings.TrimSpace(line[i+1:])
	if body == "" {
		return "", nil, "", false
	}
	return name, params, body, true
}

// parseParams splits a parameter list on top-level commas and requires each
// entry to be a valid identifier.
func parseParams(inner string) ([]string, bool) {
	var params []string
	for _, p := range control.SplitArgs(inner) {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		if !isIdent(p) {
			return nil, false
		}
		params = append(params, p)
	}
	return params, true
}

// isIdent reports whether s is a complete identifier token.
func isIdent(s string) bool {
	if s == "" || !lexer.IsIdentStart(s[0]) {
		return false
	}
	for i := 1; i < len(s); i++ {
		if !lexer.IsIdentChar(s[i]) {
			return false
		}
	}
	return true
}

// findMatchingParen returns the index of the ')' matching the '(' at openIdx,
// or -1 if unmatched.
func findMatchingParen(s string, openIdx int) int {
	depth := 0
	for i := openIdx; i < len(s); i++ {
		switch s[i] {
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 {
				return i
			}
		}
	}
	return -1
}

// splitArgs splits a comma-separated argument list at top-level commas,
// respecting nested parentheses.

// expandLet implements local bindings: let(x = e1, y = e2, ..., body).
// The last argument is the body expression; every preceding argument must be
// a binding of the form name = expr. Bindings are substituted into the body
// (and into later bindings) so names stay local and never touch global state.
func (c *Calculator) expandLet(inner string) (string, error) {
	args := control.SplitArgs(inner)
	if len(args) < 2 {
		return "", fmt.Errorf("let expects at least one binding (name = expr) and a body expression")
	}
	// First pass: resolve every binding, substituting earlier bindings into
	// later ones so let(x = 1, y = x + 1, x * y) works.
	bindings := args[:len(args)-1]
	body := args[len(args)-1]
	subs := make([]subst, 0, len(bindings))
	for _, b := range bindings {
		eq := findAssignEq(b)
		if eq < 0 {
			return "", fmt.Errorf("let binding %q must be of the form name = expr", strings.TrimSpace(b))
		}
		name := strings.TrimSpace(b[:eq])
		if !isIdent(name) {
			return "", fmt.Errorf("let binding name %q is not a valid identifier", name)
		}
		expr := strings.TrimSpace(b[eq+1:])
		if expr == "" {
			return "", fmt.Errorf("let binding for %q has an empty expression", name)
		}
		// Substitute earlier bindings into this binding's expression.
		for _, sb := range subs {
			expr = control.ReplaceIdent(expr, sb.name, sb.expr)
		}
		ex, err := c.expand(expr)
		if err != nil {
			return "", err
		}
		subs = append(subs, subst{name: name, expr: "(" + ex + ")"})
	}
	// Apply all bindings to the body, then evaluate it.
	for _, sb := range subs {
		body = control.ReplaceIdent(body, sb.name, sb.expr)
	}
	eb, err := c.expand(body)
	if err != nil {
		return "", err
	}
	return "(" + eb + ")", nil
}

// subst records a resolved local binding for let.
type subst struct {
	name string
	expr string
}

// defineFunc validates and stores a user-defined function.
func (c *Calculator) defineFunc(name string, params []string, body string) error {
	if name == "ans" || name == "mem" || name == "pi" || name == "e" || name == "convert" || name == "if" || name == "and" || name == "or" || name == "not" || name == "clamp" || name == "lerp" || name == "step" || name == "diff" || name == "pct" || name == "smoothstep" || name == "remap" || name == "var" || name == "stddev" || name == "median" || name == "count" || name == "mode" || name == "countif" || name == "sumif" || name == "let" {
		return fmt.Errorf("cannot define function %q (reserved name)", name)
	}
	if _, ok := parser.SupportedFunctions[name]; ok {
		return fmt.Errorf("cannot redefine built-in function %q", name)
	}
	if _, ok := c.vars[name]; ok {
		return fmt.Errorf("cannot define function %q: name is already a variable", name)
	}
	seen := map[string]bool{}
	for _, p := range params {
		if seen[p] {
			return fmt.Errorf("duplicate parameter %q", p)
		}
		seen[p] = true
	}
	if c.funcs == nil {
		c.funcs = make(map[string]*funcDef)
	}
	c.funcs[name] = &funcDef{Params: params, Body: body}
	fmt.Fprintf(c.out, "%s(%s) defined\n", name, strings.Join(params, ", "))
	return nil
}

// expand recursively substitutes user function calls and variables in s.
// User function calls are replaced with their parameter-bound bodies, wrapped
// in parentheses to preserve precedence.
func (c *Calculator) expand(s string) (string, error) {
	c.expandDepth++
	defer func() { c.expandDepth-- }()
	if c.expandDepth > 500 {
		return "", fmt.Errorf("recursion too deep")
	}
	var b strings.Builder
	i := 0
	n := len(s)
	for i < n {
		if !lexer.IsIdentStart(s[i]) {
			b.WriteByte(s[i])
			i++
			continue
		}
		j := i + 1
		for j < n && lexer.IsIdentChar(s[j]) {
			j++
		}
		ident := s[i:j]
		// Skip spaces to detect a call.
		k := j
		for k < n && s[k] == ' ' {
			k++
		}
		if k < n && s[k] == '(' && ident == "convert" {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to convert")
			}
			conv, err := c.expandConvert(s[k+1 : end])
			if err != nil {
				return "", err
			}
			b.WriteString(conv)
			i = end + 1
			continue
		}
		if k < n && s[k] == '(' && ident == "let" {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to let")
			}
			le, err := c.expandLet(s[k+1 : end])
			if err != nil {
				return "", err
			}
			b.WriteString(le)
			i = end + 1
			continue
		}
		if k < n && s[k] == '(' && ident == "step" {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to step")
			}
			stepped, err := c.expandStep(s[k+1 : end])
			if err != nil {
				return "", err
			}
			b.WriteString(stepped)
			i = end + 1
			continue
		}
		if k < n && s[k] == '(' && ident == "lerp" {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to lerp")
			}
			lerped, err := c.expandLerp(s[k+1 : end])
			if err != nil {
				return "", err
			}
			b.WriteString(lerped)
			i = end + 1
			continue
		}
		if k < n && s[k] == '(' && ident == "clamp" {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to clamp")
			}
			clamped, err := c.expandClamp(s[k+1 : end])
			if err != nil {
				return "", err
			}
			b.WriteString(clamped)
			i = end + 1
			continue
		}
		if k < n && s[k] == '(' && (ident == "and" || ident == "or") {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to %s", ident)
			}
			boolExpr, err := c.expandAndOr(s[k+1:end], ident)
			if err != nil {
				return "", err
			}
			b.WriteString(boolExpr)
			i = end + 1
			continue
		}
		if k < n && s[k] == '(' && ident == "not" {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to not")
			}
			b.WriteString(s[k+1:end] + " ? 0 : 1")
			i = end + 1
			continue
		}
		if k < n && s[k] == '(' && ident == "diff" {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to diff")
			}
			diffed, err := c.expandDiff(s[k+1 : end])
			if err != nil {
				return "", err
			}
			b.WriteString(diffed)
			i = end + 1
			continue
		}
		if k < n && s[k] == '(' && ident == "pct" {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to pct")
			}
			pcted, err := c.expandPct(s[k+1 : end])
			if err != nil {
				return "", err
			}
			b.WriteString(pcted)
			i = end + 1
			continue
		}
		if k < n && s[k] == '(' && ident == "smoothstep" {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to smoothstep")
			}
			smoothed, err := c.expandSmoothstep(s[k+1 : end])
			if err != nil {
				return "", err
			}
			b.WriteString(smoothed)
			i = end + 1
			continue
		}
		if k < n && s[k] == '(' && ident == "round" {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to round")
			}
			args := control.SplitArgs(s[k+1 : end])
			if len(args) == 2 {
				x, digits := strings.TrimSpace(args[0]), strings.TrimSpace(args[1])
				xE, err := c.expand(x)
				if err != nil {
					return "", err
				}
				b.WriteString("fix(" + xE + ", " + digits + ")")
				i = end + 1
				continue
			}
		}
		if k < n && s[k] == '(' && ident == "remap" {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to remap")
			}
			remapped, err := c.expandRemap(s[k+1 : end])
			if err != nil {
				return "", err
			}
			b.WriteString(remapped)
			i = end + 1
			continue
		}
		if k < n && s[k] == '(' && (ident == "floor" || ident == "ceil") {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to %s", ident)
			}
			args := control.SplitArgs(s[k+1 : end])
			if len(args) == 2 {
				x, digits := strings.TrimSpace(args[0]), strings.TrimSpace(args[1])
				xE, err := c.expand(x)
				if err != nil {
					return "", err
				}
				b.WriteString(ident + "(" + xE + " * pow(10, " + digits + ")) / pow(10, " + digits + ")")
				i = end + 1
				continue
			}
		}
		if k < n && s[k] == '(' && ident == "if" {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to if")
			}
			condExpr, err := c.expandIf(s[k+1 : end])
			if err != nil {
				return "", err
			}
			b.WriteString(condExpr)
			i = end + 1
			continue
		}
		if k < n && s[k] == '(' && (ident == "sum" || ident == "prod" || ident == "count" || ident == "avg" || ident == "min" || ident == "max" || ident == "gcd" || ident == "lcm") {
			rangeOK := true
			parenEnd := findMatchingParen(s, k)
			if parenEnd < 0 {
				rangeOK = false
			} else {
				args := control.SplitArgs(s[k+1 : parenEnd])
				if len(args) != 4 && len(args) != 5 {
					rangeOK = false
				} else if !isIdent(strings.TrimSpace(args[0])) {
					rangeOK = false
				}
			}
			if rangeOK {
				end := findMatchingParen(s, k)
				if end < 0 {
					return "", fmt.Errorf("unmatched '(' in call to %s", ident)
				}
				r, err := c.expandRangeLoop(s[k+1:end], ident)
				if err != nil {
					return "", err
				}
				b.WriteString(r)
				i = end + 1
				continue
			}
		}
		if k < n && s[k] == '(' && ident == "repeat" {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to repeat")
			}
			args := control.SplitArgs(s[k+1 : end])
			if len(args) != 3 {
				return "", fmt.Errorf("repeat expects 3 arguments (n, var, body), got %d", len(args))
			}
			r, err := c.expandRepeat(strings.TrimSpace(args[0]), strings.TrimSpace(args[1]), strings.TrimSpace(args[2]))
			if err != nil {
				return "", err
			}
			b.WriteString(r)
			i = end + 1
			continue
		}
		if k < n && s[k] == '(' && ident == "for" {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to for")
			}
			args := control.SplitArgs(s[k+1 : end])
			if len(args) != 4 && len(args) != 5 {
				return "", fmt.Errorf("for requires 4 or 5 arguments (var, lo, hi[, step], body)")
			}
			step := "1"
			body := ""
			if len(args) == 5 {
				step = strings.TrimSpace(args[3])
				body = strings.TrimSpace(args[4])
			} else {
				body = strings.TrimSpace(args[3])
			}
			expanded, err := c.expandFor(strings.TrimSpace(args[0]), strings.TrimSpace(args[1]), strings.TrimSpace(args[2]), step, body)
			if err != nil {
				return "", err
			}
			b.WriteString(expanded)
			i = end + 1
			continue
		}
		if k < n && s[k] == '(' && ident == "while" {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to while")
			}
			args := control.SplitArgs(s[k+1 : end])
			if len(args) != 2 {
				return "", fmt.Errorf("while expects 2 arguments (cond, body), got %d", len(args))
			}
			r, err := c.expandWhile(strings.TrimSpace(args[0]), strings.TrimSpace(args[1]))
			if err != nil {
				return "", err
			}
			b.WriteString(r)
			i = end + 1
			continue
		}
		if k < n && s[k] == '(' && ident == "begin" {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to begin")
			}
			r, err := c.expandBegin(s[k+1 : end])
			if err != nil {
				return "", err
			}
			b.WriteString(r)
			i = end + 1
			continue
		}
		if k < n && s[k] == '(' && (ident == "countif" || ident == "sumif") {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to %s", ident)
			}
			args := control.SplitArgs(s[k+1 : end])
			if len(args) != 3 {
				return "", fmt.Errorf("%s expects 3 argument(s) (cond, a, b), got %d", ident, len(args))
			}
			cond, lo, hi := strings.TrimSpace(args[0]), strings.TrimSpace(args[1]), strings.TrimSpace(args[2])
			loE, err := c.expand(lo)
			if err != nil {
				return "", err
			}
			hiE, err := c.expand(hi)
			if err != nil {
				return "", err
			}
			loNum, err := eval.Evaluate(loE)
			if err != nil {
				return "", fmt.Errorf("countif/sumif bounds must be numeric: %v", err)
			}
			hiNum, err := eval.Evaluate(hiE)
			if err != nil {
				return "", fmt.Errorf("countif/sumif bounds must be numeric: %v", err)
			}
			if ident == "countif" {
				b.WriteString(expandCountIf(cond, loNum, hiNum))
			} else {
				b.WriteString(expandSumIf(cond, loNum, hiNum))
			}
			i = end + 1
			continue
		}
		if k < n && s[k] == '(' && c.isFunc(ident) {
			end := findMatchingParen(s, k)
			if end < 0 {
				return "", fmt.Errorf("unmatched '(' in call to %s", ident)
			}
			inner := s[k+1 : end]
			var args []string
			if inner != "" {
				args = control.SplitArgs(inner)
			}
			def := c.funcs[ident]
			if len(args) != len(def.Params) {
				return "", fmt.Errorf("%s expects %d argument(s), got %d", ident, len(def.Params), len(args))
			}
			// Bind each argument to its evaluated numeric value so recursive calls
			// (e.g. f(n-1)) substitute a number, not a raw expression string. This
			// keeps bodies like if(n <= 1, ...) well-formed at every recursion.
			vals := make([]string, len(args))
			for ai, a := range args {
				av, err := c.eval(strings.TrimSpace(a))
				if err != nil {
					return "", err
				}
				vals[ai] = fmt.Sprintf("%g", av)
			}
			body := def.Body
			for pi, p := range def.Params {
				body = control.ReplaceIdent(body, p, vals[pi])
			}
			// Function-local scope: snapshot variables so assignments inside the
			// body (e.g. acc = ...) do not leak to the caller or between calls.
			// The snapshot is restored after the body expands, keeping recursion
			// safe and functions pure.
			snap := make(map[string]float64, len(c.vars))
			for k, v := range c.vars {
				snap[k] = v
			}
			eb, err := c.expand(body)
			for k := range c.vars {
				delete(c.vars, k)
			}
			for k, v := range snap {
				c.vars[k] = v
			}
			if err != nil {
				return "", err
			}
			b.WriteByte('(')
			b.WriteString(eb)
			b.WriteByte(')')
			i = end + 1
			continue
		}
		// Variable / ans / mem substitution.
		switch {
		case c.hasAns && ident == "ans":
			b.WriteString(strconv.FormatFloat(c.ans, 'g', -1, 64))
		case c.hasMem && ident == "mem":
			b.WriteString(strconv.FormatFloat(c.memory, 'g', -1, 64))
		default:
			if v, ok := c.vars[ident]; ok {
				b.WriteString(strconv.FormatFloat(v, 'g', -1, 64))
			} else {
				b.WriteString(ident)
			}
		}
		i = j
	}
	return b.String(), nil
}

// isFunc reports whether name is a defined user function.
func (c *Calculator) isFunc(name string) bool {
	if c.funcs == nil {
		return false
	}
	_, ok := c.funcs[name]
	return ok
}

// funcsState returns a copy of the user function definitions as value structs.
func (c *Calculator) funcsState() map[string]funcDef {
	st := map[string]funcDef{}
	if c.funcs == nil {
		return st
	}
	for name, f := range c.funcs {
		st[name] = funcDef{Params: f.Params, Body: f.Body}
	}
	return st
}

// setFuncs installs user function definitions from a persisted map.
func (c *Calculator) setFuncs(st map[string]funcDef) {
	if len(st) == 0 {
		c.funcs = nil
		return
	}
	c.funcs = make(map[string]*funcDef, len(st))
	for name, f := range st {
		f := f
		c.funcs[name] = &f
	}
}
