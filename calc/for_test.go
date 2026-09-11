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

func TestForLoopDirectBreak(t *testing.T) {
	got := runBatch(t, "s = 0\nfor(i, 1, 10, s = s + i; break)\ns\nquit\n")
	// break after the first iteration: only i=1 is added.
	if !strings.Contains(got, "\n1\n") {
		t.Fatalf("for direct break: expected s=1, got:\n%s", got)
	}
}

func TestForLoopBreakInsideBegin(t *testing.T) {
	got := runBatch(t, "s = 0\nfor(i, 1, 10, begin(s = s + i; break))\ns\nquit\n")
	if !strings.Contains(got, "\n1\n") {
		t.Fatalf("for begin break: expected s=1, got:\n%s", got)
	}
}

func TestRepeatLoopContinue(t *testing.T) {
	got := runBatch(t, "s = 0\nrepeat(5, n, begin(n; continue; s = s + 100))\ns\nquit\n")
	// continue skips the s assignment every iteration, so s stays 0.
	if !strings.Contains(got, "\n0\n") {
		t.Fatalf("repeat continue: expected s=0, got:\n%s", got)
	}
}

func TestRepeatLoopBreakInsideBegin(t *testing.T) {
	got := runBatch(t, "s = 0\nrepeat(5, n, begin(s = s + 1; break))\ns\nquit\n")
	// break after first iteration: s = 1.
	if !strings.Contains(got, "\n1\n") {
		t.Fatalf("repeat begin break: expected s=1, got:\n%s", got)
	}
}
