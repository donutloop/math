package integration_test

// Whole-program integration tests. Each case is a complete math program:
// a sequence of statements (variable setup, function definitions, loops,
// begin blocks) ending with a query that prints the program's result.
// The batch calculator is the runtime that interprets the program.

import (
	"strings"
	"testing"

	"prototype_kl/calc"
)

// runProgram feeds a whole program to the calculator and returns all output.
func runProgram(t *testing.T, program string) string {
	t.Helper()
	var out strings.Builder
	calc.NewBatch(strings.NewReader(program), &out).Run()
	return out.String()
}

// checkLine asserts the output contains an exact line equal to want.
func checkLine(t *testing.T, got, want string) {
	t.Helper()
	if !strings.Contains(got, "\n"+want+"\n") {
		t.Fatalf("expected line %q in program output:\n%s", want, got)
	}
}

func TestProgramFactorialLoop(t *testing.T) {
	// Complete program: compute 5! with a while loop and a begin block that
	// sequences two updates per iteration (multiply then decrement).
	prog := `n = 5
f = 1
while(n > 1, begin(f = f * n; n = n - 1))
f
quit
`
	got := runProgram(t, prog)
	checkLine(t, got, "120")
}

func TestProgramCumulativeSumForLoop(t *testing.T) {
	// Complete program: sum 1..10 using the range for-loop.
	prog := `s = 0
for(i, 1, 10, s = s + i)
s
quit
`
	got := runProgram(t, prog)
	checkLine(t, got, "55")
}

func TestProgramFibonacciRecursive(t *testing.T) {
	// Complete program: define a recursive fibonacci function, then call it.
	prog := `fib(n) = if(n <= 1, n, fib(n-1) + fib(n-2))
fib(7)
fib(10)
quit
`
	got := runProgram(t, prog)
	checkLine(t, got, "13")
	checkLine(t, got, "55")
}

func TestProgramScopedFunctionLocalVars(t *testing.T) {
	// Complete program: a function uses a local accumulator (snapshot/restore
	// scope) so its internal variables do not leak into the global session.
	prog := `sum_to(x) = begin(acc = 0; for(i, 1, x, acc = acc + i); acc)
sum_to(5)
acc
quit
`
	got := runProgram(t, prog)
	checkLine(t, got, "15")
	// acc must be undefined after the function call: querying it should not
	// print a number line, so the "15" line is the last numeric line.
	if strings.Count(got, "\n15\n") != 1 {
		t.Fatalf("acc leaked out of function scope:\n%s", got)
	}
}

func TestProgramSumOfSquares(t *testing.T) {
	// Complete program: sum of squares 1..10 == 385, via a while loop with a
	// paren-aware multi-statement body.
	prog := `i = 1
s = 0
while(i <= 10, s = s + i * i; i = i + 1)
s
quit
`
	got := runProgram(t, prog)
	checkLine(t, got, "385")
}

func TestProgramComposition(t *testing.T) {
	// Complete program combining all constructs: a helper function, a range
	// loop, a while loop, and a begin block, computing (1+2+...+10) + 3!.
	prog := `fact(n) = if(n <= 1, 1, n * fact(n-1))
s = 0
for(i, 1, 10, s = s + i)
s = s + fact(3)
s
quit
`
	got := runProgram(t, prog)
	checkLine(t, got, "61")
}

func TestProgramBreakInsideBegin(t *testing.T) {
	// break raised inside a begin(...) block propagates to the enclosing
	// while loop: the loop exits after the first iteration.
	prog := `x = 0
while(x < 100, begin(x = x + 1; break))
x
quit
`
	got := runProgram(t, prog)
	checkLine(t, got, "1")
}

func TestProgramContinueInsideBegin(t *testing.T) {
	// continue raised inside a begin(...) block skips the rest of the current
	// while iteration: the skipped statement never executes.
	prog := `i = 0
s = 0
while(i < 5, begin(i = i + 1; continue; s = s + 100))
s
i
quit
`
	got := runProgram(t, prog)
	checkLine(t, got, "0")
	checkLine(t, got, "5")
}

func TestProgramForLoopBreak(t *testing.T) {
	// break inside a begin(...) block propagates to the enclosing for loop:
	// only the first iteration executes, so the running sum is 1.
	prog := `s = 0
for(i, 1, 10, begin(s = s + i; break))
s
quit
`
	got := runProgram(t, prog)
	checkLine(t, got, "1")
}

func TestProgramRepeatContinue(t *testing.T) {
	// continue inside a begin(...) block skips the rest of the current repeat
	// iteration: the skipped assignment never runs, so s stays 0.
	prog := `s = 0
repeat(5, n, begin(n; continue; s = s + 100))
s
quit
`
	got := runProgram(t, prog)
	checkLine(t, got, "0")
}

func TestProgramRepeatBreak(t *testing.T) {
	// break inside a begin(...) block exits the repeat loop after iteration 1.
	prog := `s = 0
repeat(5, n, begin(s = s + 1; break))
s
quit
`
	got := runProgram(t, prog)
	checkLine(t, got, "1")
}

func TestProgramTopLevelBreakError(t *testing.T) {
	// break outside a loop is a clear program error, not a parse failure.
	prog := `begin(x = 1; break)
x
quit
`
	got := runProgram(t, prog)
	if !strings.Contains(got, "break outside a loop") {
		t.Fatalf("expected break-outside-loop error in program output:\n%s", got)
	}
}

func TestProgramNestedBreak(t *testing.T) {
	// inner begin break exits only the inner while; the outer loop continues.
	// inner runs once per outer step, so s = 1 + 2 + 3 = 6.
	prog := `i = 0
s = 0
while(i < 3, begin(i = i + 1; j = 0; while(j < 5, begin(j = j + 1; break)); s = s + i))
s
quit
`
	got := runProgram(t, prog)
	checkLine(t, got, "6")
}

func TestProgramNestedContinue(t *testing.T) {
	// inner begin continue skips the inner body's skipped statement every
	// inner iteration, but the outer accumulate still runs: s = 1 + 2 + 3.
	prog := `i = 0
s = 0
while(i < 3, begin(i = i + 1; j = 0; while(j < 5, begin(j = j + 1; continue; s = s + 100)); s = s + i))
s
quit
`
	got := runProgram(t, prog)
	checkLine(t, got, "6")
}

func TestProgramBreakValue(t *testing.T) {
	// break <value> returns an explicit result from a while loop on early
	// exit: the loop stops after the first iteration and returns i * 10.
	prog := `i = 0
r = while(i < 3, begin(i = i + 1; break i * 10))
r
quit
`
	got := runProgram(t, prog)
	checkLine(t, got, "10")
}

func TestProgramBreakInIfBranch(t *testing.T) {
	// break inside an if branch of a loop body exits the while loop.
	prog := `i = 0
while(i < 5, begin(i = i + 1; if(i == 3, break, 0)))
i
quit
`
	got := runProgram(t, prog)
	checkLine(t, got, "3")
}

func TestProgramForStep(t *testing.T) {
	// for with an explicit step accumulates odd numbers 1..9 -> 25.
	prog := `s = 0
for(i, 1, 9, 2, s = s + i)
s
quit
`
	got := runProgram(t, prog)
	checkLine(t, got, "25")
}
