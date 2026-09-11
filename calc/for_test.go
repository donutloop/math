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

func TestBreakOutsideLoopClearError(t *testing.T) {
	got := runBatch(t, "begin(x = 1; break)\nquit\n")
	// top-level break must report a clear error, not a parse failure.
	if !strings.Contains(got, "break outside a loop") {
		t.Fatalf("expected clear break-outside-loop error, got:\n%s", got)
	}
}

func TestContinueOutsideLoopClearError(t *testing.T) {
	got := runBatch(t, "begin(x = 1; continue)\nquit\n")
	if !strings.Contains(got, "continue outside a loop") {
		t.Fatalf("expected clear continue-outside-loop error, got:\n%s", got)
	}
}

func TestBreakBlockAsAssignmentValue(t *testing.T) {
	got := runBatch(t, "x = begin(1; break)\nquit\n")
	if !strings.Contains(got, "break outside a loop") {
		t.Fatalf("expected clear error when a break block is an assignment value, got:\n%s", got)
	}
}

func TestWhileBreakValue(t *testing.T) {
	got := runBatch(t, "x = 0\nwhile(x < 100, begin(x = x + 1; break x))\nquit\n")
	if !strings.Contains(got, "\n1\n") {
		t.Fatalf("while break value: expected 1, got:\n%s", got)
	}
}

func TestForBreakValue(t *testing.T) {
	got := runBatch(t, "for(i, 1, 10, begin(i; break i * 100))\nquit\n")
	if !strings.Contains(got, "100") {
		t.Fatalf("for break value: expected 100, got:\n%s", got)
	}
}

func TestRepeatBreakValue(t *testing.T) {
	got := runBatch(t, "repeat(5, n, begin(n; break n * 1000))\nquit\n")
	if !strings.Contains(got, "1000") {
		t.Fatalf("repeat break value: expected 1000, got:\n%s", got)
	}
}

func TestBreakInsideIfBranch(t *testing.T) {
	got := runBatch(t, "i = 0\nwhile(i < 5, begin(i = i + 1; if(i == 3, break, 0)))\ni\nquit\n")
	if !strings.Contains(got, "3") {
		t.Fatalf("break in if branch: expected i=3, got:\n%s", got)
	}
}

func TestContinueInsideIfBranch(t *testing.T) {
	got := runBatch(t, "i = 0\ns = 0\nwhile(i < 5, begin(i = i + 1; if(i == 2, continue, 0); s = s + i))\ns\nquit\n")
	// continue skips s += i when i == 2, so s = 1 + 3 + 4 + 5 = 13.
	if !strings.Contains(got, "13") {
		t.Fatalf("continue in if branch: expected s=13, got:\n%s", got)
	}
}

func TestBreakValueInsideIfBranch(t *testing.T) {
	got := runBatch(t, "i = 0\nwhile(i < 5, begin(i = i + 1; if(i == 3, break i * 100, 0)))\nquit\n")
	if !strings.Contains(got, "300") {
		t.Fatalf("break value in if branch: expected 300, got:\n%s", got)
	}
}
