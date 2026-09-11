package calc

import (
	"strings"
	"testing"
)

func TestFunctionLocalScope(t *testing.T) {
	// Assignments inside a function body must not leak to the caller: f(5)
	// returns 15, but acc is undefined afterward.
	got := runBatch(t, "f(n) = begin(acc = 0; for(i, 1, n, acc = acc + i); acc)\ny = f(5)\ny\nacc\nquit\n")
	if !strings.Contains(got, "\n15\n") {
		t.Fatalf("function result wrong:\n%s", got)
	}
	if strings.Contains(got, "\n55\n") {
		t.Fatalf("function-local acc leaked:\n%s", got)
	}
}
