package calc

import (
	"strings"
	"testing"
)

func TestWhileCount(t *testing.T) {
	// Counting loop: a mutable variable is incremented until the condition
	// goes false, then its final value is read back.
	got := runBatch(t, "x = 0\nwhile(x < 10, x = x + 1)\nx\nquit\n")
	if !strings.Contains(got, "\n10\n") {
		t.Fatalf("while count failed:\n%s", got)
	}
}

func TestWhileAccumulate(t *testing.T) {
	// Accumulation loop with a step > 1: acc 0->2->4->6->8->10.
	got := runBatch(t, "acc = 0\nwhile(acc < 10, acc = acc + 2)\nacc\nquit\n")
	if !strings.Contains(got, "\n10\n") {
		t.Fatalf("while accumulate failed:\n%s", got)
	}
}

func TestWhileConditionFalseImmediately(t *testing.T) {
	// Condition false on first check: body must not run, result is 0.
	got := runBatch(t, "x = 0\nwhile(x > 10, x = x + 1)\nwhile(x > 10, x = x + 1)\nx\nquit\n")
	if !strings.Contains(got, "\n0\n") {
		t.Fatalf("while immediate-false failed:\n%s", got)
	}
}

func TestWhileInfiniteGuard(t *testing.T) {
	// A non-progressing body must hit the iteration bound instead of hanging.
	got := runBatch(t, "x = 0\nwhile(x < 10, x = x + 0)\nquit\n")
	if !strings.Contains(got, "exceeded 10000 iterations") {
		t.Fatalf("expected infinite-loop guard:\n%s", got)
	}
}

func TestWhileMultiStatementBody(t *testing.T) {
	// Paren-aware statement splitting lets a while body sequence multiple
	// updates: acc accumulates x and x increments each iteration.
	got := runBatch(t, "x = 0\nacc = 0\nwhile(x < 5, acc = acc + x; x = x + 1)\nacc\nquit\n")
	if !strings.Contains(got, "\n10\n") {
		t.Fatalf("while multi-statement body failed:\n%s", got)
	}
}
