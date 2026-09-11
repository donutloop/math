package calc

import (
	"strings"
	"testing"
)

func TestRecursiveFactorial(t *testing.T) {
	// Recursion: f(n) = if(n <= 1, 1, n * f(n-1)). Arguments must bind to
	// evaluated numbers so f(n-1) substitutes 3 (not "n-1"/"4-1" strings).
	got := runBatch(t, "f(n) = if(n <= 1, 1, n * f(n-1))\nf(5)\nf(10)\nquit\n")
	if !strings.Contains(got, "\n120\n") || !strings.Contains(got, "\n3628800\n") {
		t.Fatalf("recursive factorial failed:\n%s", got)
	}
}

func TestRecursiveFibonacci(t *testing.T) {
	got := runBatch(t, "fib(n) = if(n <= 1, n, fib(n-1) + fib(n-2))\nfib(6)\nquit\n")
	if !strings.Contains(got, "\n8\n") {
		t.Fatalf("recursive fibonacci failed:\n%s", got)
	}
}
