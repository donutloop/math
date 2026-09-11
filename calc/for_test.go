package calc

import (
	"strings"
	"testing"
)

func TestForLoopSum(t *testing.T) {
	// for(var, lo, hi, body) binds var to each integer in [lo,hi] inclusive;
	// the body may be an assignment that accumulates a mutable variable.
	got := runBatch(t, "s = 0\nfor(i, 1, 5, s = s + i)\ns\nquit\n")
	if !strings.Contains(got, "\n15\n") {
		t.Fatalf("for loop sum failed:\n%s", got)
	}
}

func TestForLoopValue(t *testing.T) {
	// Non-assignment bodies return the last body value (range 2..4 => 4*4).
	got := runBatch(t, "for(i, 2, 4, i * i)\nquit\n")
	if !strings.Contains(got, "16") {
		t.Fatalf("for loop value failed (got: %q):\n%s", got, got)
	}
}
