package calc

import (
	"strings"
	"testing"
)

// runBatch feeds input lines to the calculator and returns all output.
func runBatch(t *testing.T, input string) string {
	t.Helper()
	var out strings.Builder
	NewBatch(strings.NewReader(input), &out).Run()
	return out.String()
}

func TestLetSingleBinding(t *testing.T) {
	got := runBatch(t, "let(x = 2, x^2 + x)\nquit\n")
	if !strings.Contains(got, "6") {
		t.Errorf("let(x=2, x^2+x) = 6, got:\n%s", got)
	}
}

func TestLetMultiBinding(t *testing.T) {
	got := runBatch(t, "let(x = 1, y = x + 1, x * y)\nquit\n")
	if !strings.Contains(got, "2") {
		t.Errorf("let(x=1, y=x+1, x*y) = 2, got:\n%s", got)
	}
}

func TestLetThreeBindings(t *testing.T) {
	got := runBatch(t, "let(a = 5, b = 10, c = a + b, c * 2)\nquit\n")
	if !strings.Contains(got, "30") {
		t.Errorf("let(a=5,b=10,c=a+b,c*2) = 30, got:\n%s", got)
	}
}

func TestLetScopingDoesNotPolluteGlobals(t *testing.T) {
	got := runBatch(t, "let(x = 99, x + 1)\nx\nquit\n")
	// The global x must remain undefined; 'x' alone should error, not print 99.
	if strings.Contains(got, "99") && !strings.Contains(got, "undefined") {
		t.Errorf("let binding leaked into global scope:\n%s", got)
	}
}

func TestLetErrors(t *testing.T) {
	got := runBatch(t, "let(x)\nlet(x = , 1)\nquit\n")
	if !strings.Contains(got, "expects at least one binding") {
		t.Errorf("let with no body should error, got:\n%s", got)
	}
	if !strings.Contains(got, "empty expression") {
		t.Errorf("let with empty binding expr should error, got:\n%s", got)
	}
}
