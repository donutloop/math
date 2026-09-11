package control

import (
	"prototype_kl/lexer"
	"strings"
)

// Control-flow sentinels used by loop expansion.
const (
	BreakValuePrefix = "__BREAK_VALUE__:"
	BreakSentinel    = "__BREAK__"
	ContinueSentinel = "__CONTINUE__"
)

func SplitArgs(s string) []string {
	var args []string
	depth := 0
	start := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			depth++
		case ')':
			depth--
		case ',':
			if depth == 0 {
				args = append(args, s[start:i])
				start = i + 1
			}
		}
	}
	args = append(args, s[start:])
	return args
}
func ReplaceIdent(s, target, replacement string) string {
	var b strings.Builder
	i := 0
	for i < len(s) {
		if lexer.IsIdentStart(s[i]) {
			j := i + 1
			for j < len(s) && lexer.IsIdentChar(s[j]) {
				j++
			}
			if s[i:j] == target {
				b.WriteString(replacement)
			} else {
				b.WriteString(s[i:j])
			}
			i = j
			continue
		}
		b.WriteByte(s[i])
		i++
	}
	return b.String()
}
func BreakValue(t string) (string, bool) {
	if t == "break" {
		return "", false
	}
	if !strings.HasPrefix(t, "break") {
		return "", false
	}
	rest := strings.TrimSpace(t[len("break"):])
	if rest == "" {
		return "", false
	}
	return rest, true
}
