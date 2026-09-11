package calc

import (
	"strings"
	"testing"
)

func TestBeginBlock(t *testing.T) {
	// begin(...) sequences statements and returns the last value.
	got := runBatch(t, "a = 0\nbegin(a = a + 5; b = a * 2)\nb\nquit\n")
	if !strings.Contains(got, "\n10\n") {
		t.Fatalf("begin block failed:\n%s", got)
	}
}

func TestBeginAsLoopBody(t *testing.T) {
	// begin(...) as a while body: accumulate x and increment x each step.
	got := runBatch(t, "x = 0\nacc = 0\nwhile(x < 4, begin(acc = acc + x; x = x + 1))\nacc\nquit\n")
	if !strings.Contains(got, "\n6\n") {
		t.Fatalf("begin as loop body failed:\n%s", got)
	}
}
