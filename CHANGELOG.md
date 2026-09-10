## [135] trigamma(x)

- New `trigamma(x)`: the trigamma function psi'(x), via recurrence into
  the large-x regime plus asymptotic expansion to x^7.
  trigamma(1)=1.6449 (zeta(2)=pi^2/6), trigamma(2)=0.6449.
- Domain error at nonpositive integers.
- Updates parser dispatch, arity table, help, operations.md, unit test
  (calc/trigamma_test.go), ADR.
## [134] nextperfect(n)

- New `nextperfect(n)`: the smallest perfect number >= n, scanning upward
  with the aliquot-sum criterion. nextperfect(1)=6, nextperfect(7)=28,
  nextperfect(29)=496.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [133] nextdeficient(n)

- New `nextdeficient(n)`: the smallest deficient number >= n, scanning
  upward with the aliquot-sum criterion. nextdeficient(5)=5,
  nextdeficient(6)=7, nextdeficient(12)=13.
- Complements nextabundant. Updates parser dispatch, arity table, help,
  operations.md, verify, ADR.
## [132] nextabundant(n)

- New `nextabundant(n)`: the smallest abundant number >= n, scanning upward
  with the aliquot-sum criterion. nextabundant(7)=12, nextabundant(12)=12,
  nextabundant(18)=18.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [131] eta(x)

- New `eta(x)`: the Dirichlet eta function (1-2^(1-x))*zeta(x), with the
  analytic limit eta(1)=ln2. eta(2)=0.8225 (pi^2/12), eta(1)=0.6931.
- Updates parser dispatch, arity table, help, operations.md, unit test
  (calc/eta_test.go), ADR.
## [130] collatzmax(n)

- New `collatzmax(n)`: the maximum value reached in the Collatz sequence
  from n down to 1. collatzmax(3)=16, collatzmax(6)=16, collatzmax(1)=1.
- Complements collatz(n) (step count). Updates parser dispatch, arity
  table, help, operations.md, verify, ADR.
## [129] partition(n)

- New `partition(n)`: the integer partition count p(n) via the
  pentagonal-number recurrence. partition(0)=1, partition(4)=5,
  partition(5)=7, partition(6)=11.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [128] multinomial(n, a, b, ...)

- New variadic `multinomial(n, a, b, ...)`: n!/(a!b!...) via chained choose
  coefficients; the groups must sum to n. multinomial(5, 2, 3)=10,
  multinomial(6, 2, 2, 2)=90.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [127] risingfact(n, k)

- New `risingfact(n, k)`: the rising factorial (Pochhammer symbol),
  product n(n+1)...(n+k-1). risingfact(1,3)=6, risingfact(2,3)=24,
  risingfact(5,0)=1.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [126] doublefactorial(n)

- New `doublefactorial(n)`: the semi-factorial n!!, product of every other
  integer down from n. doublefactorial(0)=1, doublefactorial(4)=8,
  doublefactorial(5)=15, doublefactorial(6)=48.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [125] digitproduct(n)

- New `digitproduct(n)`: the product of the decimal digits.
  digitproduct(123)=6, digitproduct(50)=0, digitproduct(9)=9,
  digitproduct(0)=0.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [124] lucas(n)

- New `lucas(n)`: the Lucas numbers L_n with L_0=2, L_1=1 and the
  Fibonacci recurrence. lucas(0)=2, lucas(3)=4, lucas(4)=7, lucas(5)=11.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [123] derangements(n)

- New `derangements(n)`: the subfactorial !n, permutations of n elements
  with no fixed points, via the recurrence !n = (n-1)*(!(n-1)+!(n-2)).
  derangements(2)=1, derangements(3)=2, derangements(4)=9, derangements(5)=44.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [122] stirling(n, k)

- New `stirling(n, k)`: the Stirling number of the second kind, partitions
  of n into k non-empty subsets, via the recurrence S(n,k)=k*S(n-1,k)+S(n-1,k-1).
  stirling(3,2)=3, stirling(4,2)=7, stirling(4,3)=6, stirling(0,0)=1.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [121] catalan(n)

- New `catalan(n)`: the Catalan number C_n = choose(2n,n)/(n+1).
  catalan(0)=1, catalan(3)=5, catalan(4)=14, catalan(5)=42.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [120] bell(n)

- New `bell(n)`: the Bell number, the count of set partitions of n elements,
  via the recurrence B(i)=sum choose(i-1,k)*B(k). bell(0)=1, bell(3)=5,
  bell(4)=15, bell(5)=52.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [119] nthdigit(n, k)

- New `nthdigit(n, k)`: the k-th decimal digit of n, 0-indexed from the
  right. nthdigit(12345,0)=5, nthdigit(12345,1)=4, nthdigit(12345,4)=1.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [118] digitalroot(n)

- New `digitalroot(n)`: the iterated digit sum until a single digit.
  digitalroot(12345)=6, digitalroot(48)=3, digitalroot(9)=9,
  digitalroot(0)=0.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [117] harmonic(n)

- New `harmonic(n)`: the harmonic number H_n = sum 1/k for k=1..n.
  harmonic(1)=1, harmonic(3)=1.8333, harmonic(10)=2.9289.
- Updates parser dispatch, arity table, help, operations.md, unit test
  (calc/harmonic_test.go), ADR.
## [116] besselj0(x) and besselj1(x)

- New `besselj0(x)` and `besselj1(x)`: Bessel functions of the first kind
  via the power series to 40 terms. besselj0(0)=1, besselj1(0)=0,
  besselj0(1)=0.7652, besselj1(1)=0.4401.
- Updates parser dispatch, arity table, help, operations.md, unit test
  (calc/bessel_test.go), ADR.
## [115] loggamma(x)

- New `loggamma(x)`: ln(gamma(x)) via recurrence into the large-x regime
  plus Stirling expansion to the x^5 term. loggamma(1)=0, loggamma(5)=ln24
  (3.1781), loggamma(0.5)=ln(sqrt(pi)) (0.5724).
- Domain error at nonpositive integers.
- Updates parser dispatch, arity table, help, operations.md, unit test
  (calc/loggamma_test.go), ADR.
## [114] amicable(a, b)

- New `amicable(a, b)` predicate: 1 if a and b form an amicable pair
  (aliquot sums match), 0 otherwise; equal a=b pairs are rejected.
  amicable(220, 284)=1, amicable(10, 5)=0, amicable(6, 6)=0.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [113] sumproperdivisors(n)

- New `sumproperdivisors(n)`: the aliquot sum (sum of proper divisors).
  sumproperdivisors(12)=16, sumproperdivisors(6)=6, sumproperdivisors(28)=28
  (6 and 28 are perfect, so their aliquot sums equal themselves).
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [112] carmichael(n)

- New `carmichael(n)`: the Carmichael lambda function, the largest lambda
  with a^lambda = 1 mod n for all a coprime to n. lcm over prime-power
  factors; powers of 2 follow the special lambda(2^k)=2^(k-2) rule for k>=3.
  carmichael(9)=6, carmichael(12)=2, carmichael(10)=4, carmichael(1)=1.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [111] liouville(n)

- New `liouville(n)`: the Liouville function (-1)^bigomega(n).
  liouville(12)=-1, liouville(8)=-1, liouville(6)=1, liouville(1)=1.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [110] mertens(n)

- New `mertens(n)`: the Mertens function, sum of mobius(1..n).
  mertens(10)=-1, mertens(1)=1, mertens(6)=-1.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [109] tau(n) and mobius(n)

- New `tau(n)` (divisor count) and `mobius(n)` (Moebius function).
  tau(12)=6, tau(6)=4, mobius(6)=1, mobius(30)=-1, mobius(12)=0,
  mobius(1)=1.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [108] omega(n) and bigomega(n)

- New `omega(n)` (omega prime-factor count) and `bigomega(n)` (Omega with
  multiplicity). omega(30)=3, omega(12)=2, bigomega(12)=3, bigomega(8)=3.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [107] choose(n, k)

- New `choose(n, k)`: binomial coefficient via the multiplicative formula
  with min(k, n-k). choose(5,2)=10, choose(10,3)=120, choose(5,0)=1.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [106] sumofprimes(n)

- New `sumofprimes(n)`: sum of the first n primes, complementing primorial.
  sumofprimes(5)=28, sumofprimes(10)=129.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [105] repeat(n, var, body)

- New language-level loop macro `repeat(n, var, body)`: evaluates body with
  var bound to each integer 1..n and returns the LAST evaluated value.
  Useful for final-iteration computations and nested iteration.
- repeat(5, i, i*i)=25, repeat(3, j, j+j)=6, nested loops supported.
- Updates macro dispatch, help, operations.md, unit test
  (calc/repeat_test.go), ADR.
## [104] digamma(x)

- New `digamma(x)`: the psi (digamma) function via recurrence into the
  large-x regime plus an asymptotic expansion. digamma(1)=-0.5772 (the
  negative Euler-Mascheroni constant), digamma(0.5)=-1.9635.
- Domain error at nonpositive integers.
- Updates parser dispatch, arity table, help, operations.md, unit test
  (calc/digamma_test.go), ADR.
## [103] prevprime(n)

- New `prevprime(n)`: largest prime strictly below n. prevprime(10)=7,
  prevprime(100)=97, prevprime(3)=2. Domain error for n <= 2.
- Complements nextprime. Updates parser dispatch, arity table, help,
  operations.md, verify, ADR.
## [102] primecount(n)

- New `primecount(n)` = pi(n): number of primes <= n. primecount(10)=4,
  primecount(100)=25, primecount(2)=1, primecount(0)=0.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [101] zeta(x)

- New `zeta(x)`: Riemann zeta for x > 1 via a 2000-term series plus an
  Euler-Maclaurin tail correction. zeta(2)=1.644934, zeta(3)=1.202056.
  Domain error for x <= 1.
- Updates parser dispatch, arity table, help, operations.md, unit test
  (calc/zeta_test.go), ADR.
## [100] lambertw(x)

- New `lambertw(x)`: principal branch of the Lambert W function via Newton
  iteration. lambertw(0)=0, lambertw(1)=0.56714329 (omega constant),
  lambertw(e)=1. Domain error for x < -1/e.
- Updates parser dispatch, arity table, help, operations.md, verify, unit
  test (calc/lambertw_test.go), ADR.
## [99] Range gcd/lcm

- `gcd` and `lcm` now accept range-loop form alongside their variadic
  math form: gcd(i, 1, 10, i)=1, lcm(i, 1, 5, i)=60.
- Disambiguation mirrors min/max/avg: range only when the first arg is a
  loop variable and the call has 4-5 args.
- Updates dispatch, expandRangeLoop, help, operations.md, tests, ADR.
## [98] isabundant(n) / isdeficient(n)

- New `isabundant(n)`: 1 if sigma(n) > 2n (abundant), else 0. isabundant(12)=1.
- New `isdeficient(n)`: 1 if sigma(n) < 2n (deficient), else 0. isdeficient(8)=1.
- Domain error for n <= 0. Completes the abundant/perfect/deficient trio.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [97] Range min/max/avg

- `min`, `max`, and `avg` now accept range-loop form alongside their
  variadic math form: min(i, 1, 4, i*i)=1, max(i, 1, 4, i*i)=16,
  avg(i, 1, 3, i)=2.
- Disambiguation: a call is a range loop only when the first argument is a
  loop variable and the call has 4-5 args; otherwise it falls through to the
  variadic parser function.
- Updates dispatch, loop expansion, help, operations.md, tests, ADR.
## [96] primorial(n)

- New `primorial(n)`: product of the first n primes. primorial(1)=2,
  primorial(3)=30, primorial(4)=210. Domain error for n <= 0.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [95] digitsum / digitcount

- New `digitsum(n)`: sum of decimal digits (digitsum(1234)=10, digitsum(0)=0).
- New `digitcount(n)`: number of decimal digits (digitcount(1234)=4, digitcount(0)=1).
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [94] isperfect(n)

- New `isperfect(n)`: 1 if n is a perfect number (sigma(n) == 2n), else 0.
  isperfect(6)=1, isperfect(28)=1, isperfect(10)=0. Domain error for n <= 0.
- Updates parser dispatch, arity table, help, operations.md, verify, ADR.
## [93] Number theory: totient and sigma

- New `totient(n)`: Euler\u2019s phi — count of integers 1..n coprime to n
  (totient(12)=4, totient(9)=6).
- New `sigma(n)`: sum of positive divisors of n (sigma(6)=12, sigma(28)=56).
- Both reject non-positive / zero inputs with a typed domain error.
- Updates parser dispatch, arity table, help text, docs, verify cases, ADR.
## [92] Variadic gcd/lcm

- `gcd` and `lcm` now accept two or more arguments (like `min`/`max`/`avg`),
  folding left-to-right: `gcd(12, 18, 24)` -> 6, `lcm(4, 6, 10)` -> 60.
- Updates parser dispatch, arity table, help text, docs, verify cases, and ADR.
## [91] Local bindings: let

- New language-level `let(x = e1, y = e2, body)` expression: the last argument
  is the body; earlier arguments bind names that stay local and never touch
  global variables. Bindings can reference earlier bindings, and substituted
  values are parenthesized to preserve precedence.
  - `let(x = 2, x^2 + x)` -> 6
  - `let(x = 1, y = x + 1, x * y)` -> 2
  - `let(a = 5, b = 10, c = a + b, c * 2)` -> 30
- `findAssignEq` now only treats `=` at paren depth 0 as an assignment, so
  `let(x = 2, ...)` is not misparsed as a global assignment.
- `let` is a reserved name so it cannot be shadowed by user definitions.
- Adds unit tests, help text, and documentation.
## [90] Collatz stopping time: collatz

- `collatz(n)` returns the number of Collatz steps to reach 1 (the stopping
  time): collatz(1)=0, collatz(3)=7, collatz(27)=111.
- Requires a positive integer n <= 100000; n <= 0, non-integer, or oversized
  inputs are rejected as typed domain errors.
- Uses a step guard (10000) so non-terminating inputs are reported rather than
  looping forever.

## [89] modular exponentiation: powmod

- `powmod(a, b, m)` computes (a^b) mod m using fast modular exponentiation
  (powmod(2, 10, 1000)=24, powmod(4, 13, 497)=445).
- Requires a, b >= 0 and m > 0 as integers; negative, non-integer, or m <= 0
  inputs are rejected as typed domain errors.
- Keeps intermediate products reduced mod m, so large exponents stay fast.

## [88] Fibonacci: fib

- `fib(n)` returns the n-th Fibonacci number (fib(0)=0, fib(1)=1,
  fib(10)=55, fib(30)=832040), computed iteratively.
- Exact for n <= 97 (fib(97) < 2^63); larger n is rejected as a typed
  domain error, as is any negative or non-integer input.

## [87] digit functions: sumdigits / rev / ispal

- `sumdigits(n)` returns the sum of the decimal digits of n (sumdigits(1234)=10).
- `rev(n)` returns the decimal digits reversed, dropping leading zeros
  (rev(1234)=4321, rev(120)=21).
- `ispal(n)` returns 1 if n reads the same forward and backward, else 0
  (ispal(121)=1, ispal(123)=0).
- All accept non-negative integers; negative or non-integer inputs are
  rejected as typed domain errors.

## [86] combinatorics: npr / ncr

- `npr(n, r)` computes permutations n!/(n-r)! — ways to pick r ordered items
  from n (e.g. npr(5, 2)=20).
- `ncr(n, r)` computes combinations n!/(r!·(n-r)!) — ways to pick r unordered
  items from n (e.g. ncr(5, 2)=10).
- Both accept non-negative integers with r <= n; negative, non-integer, and
  r > n inputs are rejected as typed domain errors.
- Computed with multiplicative loops (no huge intermediate factorials) and
  the symmetric ncr reduction min(r, n-r), so larger n/r values stay accurate.
- Examples: npr(10, 3)=720, ncr(10, 5)=252, ncr(n, 0)=1, ncr(n, n)=1.

## [85] number theory functions

- `isprime(x)` returns 1 if x is prime, else 0 (x < 2 is not prime).
- `prime(n)` returns the n-th prime (1-indexed: prime(1)=2, prime(2)=3, ...).
- `nextprime(n)` returns the smallest prime >= n.
- `divcount(n)` returns the number of positive divisors of n.
- Implemented as parser-level built-ins (deterministic, verifyable).
- Examples: isprime(17)=1, isprime(15)=0, prime(10)=29, nextprime(10)=11,
  divcount(6)=4, divcount(12)=6.

## [84] generalized range loops (for-loop over integers)

- `sum(i, lo, hi, expr)` sums `expr` over integer `i` in [lo, hi].
- `sum(i, lo, hi, step, expr)` same, stepping by `step`.
- `prod(i, lo, hi[, step], expr)` product of `expr` over integer `i`.
- `count(i, lo, hi[, step], cond)` counts `i` where `cond` is nonzero.
- The loop variable is a real identifier; bounds, step, and body may reference
  user variables, ans, or other functions. Implemented as a calc-level macro
  expansion (generalizes the plain numeric sum/prod/count forms).
- Examples: `sum(i, 1, 10, i^2)` = 385, `prod(i, 1, 6, i)` = 6! = 720,
  `sum(i, 1, 10, 2, i)` = 25, `count(i, 1, 10, i % 2 == 0)` = 5.

# Changelog

All notable changes to the Math Calculator. Entries are grouped per commit,
newest first. Each feature ships code, tests, help text, verify coverage, and
docs updates.

## [83] bitwise operators

- Single `&` (AND), `|` (OR), `~` (NOT), `<<` (left shift), `>>` (right shift).
- Precedence: bind tighter than comparison and looser than addition.
- `12 & 10` = 8, `12 | 10` = 14, `5 << 2` = 20, `16 >> 3` = 2, `~5` = -6.

## [82] recursive user-defined functions

- `if(cond, then, else)` now evaluates `cond` at expansion time and expands only
  the taken branch, so recursive user functions terminate instead of killing the
  process with exponential macro expansion.
- Adds an expansion-depth guard (500) that turns runaway recursion into a
  `recursion too deep` error rather than a stack overflow.
- Verified: `fib(n) = if(n < 2, n, fib(n-1) + fib(n-2))` -> `fib(6)`=8, `fib(10)`=55.
- `if(1, 0, 1/0)` now yields 0 (untaken branch not evaluated).

## [81] tree command: AST inspection
- New `tree <expr>` command prints the abstract syntax tree (S-expression) of an
  expression after calc-macro expansion, e.g. `tree 7 % 3` -> `(mod 7 3)`.
- Adds `parser.Parse` (lex+parse without eval) and `parser.Dump` AST pretty-printer.
- Supports Number, BinaryOp, UnaryOp, Function, Postfix (%%/!), Ternary nodes.

## [80] binary %% modulo operator
- New binary modulo operator `%%` (e.g. `7 %% 3` -> 1), same precedence as `*` and `/`.
- Postfix percent (e.g. `50%%` -> 0.5) still works; `%%` is treated as modulo when an operand follows.
- `%%` works inside countif/sumif conditions: `countif(x %% 2 == 0, 1, 10)` counts evens.
- Added OpMod evaluation (math.Mod), lexer disambiguation, parser mapping, and tests.

## [79] conditional countif / sumif range loops
- New `countif(cond, lo, hi)` counts integers `x` in `[lo, hi]` satisfying `cond`.
- New `sumif(cond, lo, hi)` sums those `x` values.
- `x` is a placeholder substituted into `cond` (e.g. `countif(x > 3, 1, 10)`).
- Fixed assignment detection so `==`, `>=`, `<=`, `!=` are no longer mistaken
  for `=` assignments, enabling comparisons inside expressions and calls.

## [78] range-loop step parameter
- `sum(a, b)`, `prod(a, b)`, `count(a, b)` now accept an optional step:
  `sum(1, 10, 2)` sums every 2nd integer (25).

## [77] mode(a, b, ...) most frequent value
- Added `mode(a, b, ...)` returning the most frequent value (first on a tie),
  completing the statistics family (avg, var, stddev, median, mode).

## [76] round(x, n) decimal precision
- Made `round` variadic: `round(x)` nearest integer, `round(x, n)` rounds to n
  decimal places (e.g. `round(2.5678, 2)` -> 2.57).

## [75] count range loop
- Added `count(a, b)` returning the number of integers from floor(a) to floor(b)
  (0 when b < a), completing the integer-range loop family (sum, prod, count).

## [74] temperature units
- Added affine temperature units to `convert`: `c`/`celsius`, `f`/`fahrenheit`,
  `k`/`kelvin` (e.g. `convert(100, c, f)` -> 212).
- Generalized unitInfo with an `offset` field; conversion formula is now
  `((value * scale) + offset - to.offset) / scale`.

## [73] gradians trig mode
- Added a `grad` command and `--grad` CLI flag for gradian trig mode (200 in a
  full circle), completing the degrees/radians/gradians triad.
- Trig conversion generalized in degrees.go to `applyTrig(expr, factor)` for
  degrees (180) and gradians (200).
- New `SetGrad()` API; `deg`/`rad`/`grad` commands and prompt/status reflect
  the active mode; gradian mode persists across sessions.

## [72] statistics functions
- Added `var(a, b, ...)` population variance, `stddev(a, b, ...)` population
  standard deviation, and `median(a, b, ...)` middle value, all variadic with
  at least two arguments.

## [71] floor(x, n) / ceil(x, n)
- `floor(x, n)` and `ceil(x, n)` round down/up to n decimal places.
## [70] remap(x, lo, hi, nlo, nhi)
- Added `remap(x, lo, hi, nlo, nhi)` mapping a value between ranges.
## [69] round(x, n) to n decimal places
- `round(x, n)` rounds x to n decimal places; `round(x)` stays integer.
## [68] smoothstep(x, e0, e1) easing
- Added `smoothstep(x, e0, e1)` Hermite easing via min/max clamp.
## [67] diff(a, b) and pct(x, total) utilities
- Added `diff(a, b)` = |a - b| and `pct(x, total)` = (x / total) * 100.
## [66] avg(...) variadic arithmetic mean
- `avg` now accepts any number of values (>= 2), returning their mean.
## [65] parser: comparisons compose in sub-expressions
- Comparisons now parse inside parentheses and function arguments.
- Added `and(a, b)`, `or(a, b)`, `not(a)` boolean helpers.
## [64] step(x, edge) Heaviside step
- Added `step(x, edge)` = 1 when x >= edge, else 0.
## [63] lerp(a, b, t) linear interpolation
- Added `lerp(a, b, t)` = a + (b - a) * t.
## [62] clamp(x, lo, hi)
- Added `clamp(x, lo, hi)` bounding x into [lo, hi] via min/max.
## [61] avg(a, b) arithmetic mean
- Added `avg(a, b)` returning (a + b) / 2.
## [60] fix(x, n) decimal rounding
- Added `fix(x, n)` rounding x to n decimal places.
## [59] if conditional
- Added `if(cond, then, else)` with lazy evaluation of the selected branch.
- Nonzero cond selects then; zero selects else.
## [58] range loops (sum/prod)
- Added `sum(a, b)` and `prod(a, b)` integer range loops (floor(a) to floor(b)).
- Reversed ranges yield 0 (sum) and 1 (prod).
## [57] unit conversion
- Added `convert(value, from, to)` for length, mass, and time units.
- Supported units: m km cm mm nm mi ft in yd; kg g mg lb oz; s min h hr day.
- Cross-dimension and unknown-unit conversions are rejected; `convert` is reserved.

## [56] user-defined functions
- Added user-defined functions: `f(x) = expr` defines a function; calls use normal syntax `f(2)`.
- Parameters bind to evaluated arguments; bodies may reference variables, built-ins, and other user functions.
- Arity mismatches and reserved-name redefinition are rejected.
- Function definitions persist across saves and survive undo/redo.

## [55] mish activation
- Added `mish(x)` = x·tanh(ln(1+e^x)).

## [54] isfinite predicate
- Added `isfinite(x)` returning 1 (finite) or 0 (infinite/NaN).

## [53] swish activation
- Added `swish(x)` = x/(1+e^-x).

## [52] isqrt integer square root
- Added `isqrt(x)` = floor(sqrt(x)) for x >= 0.

## [51] softsign activation
- Added `softsign(x)` = x/(1+|x|).

## [50] fract(x) fractional part
- Added `fract(x)` = x - floor(x), the nonnegative fractional part.

## [49] root(x, n) n-th root
- Added `root(x, n)` = x^(1/n).

## [48] deg/rad conversion functions
- Added `deg(x)` radians-to-degrees and `rad(x)` degrees-to-radians.

## [47] logical && and || operators
- Added `&&` and `||` returning 1 (true) or 0 (false).
- Precedence: ternary < logical < comparison < arithmetic.

## [46] == and != equality comparisons
- Added `==` and `!=` with lexer peeking on '=' and '!'.
- Factorial `!` preserved by peeking only when followed by '='.

## [45] <= and >= comparisons
- Added two-character `<=` and `>=` with lexer peeking.
- Distinct op values and token-type mapping fix precedence.

## [44] comparison operators
- Added `<` and `>` comparisons returning 1 (true) or 0 (false).
- New tokens/ops, parseComparison precedence between ternary and arithmetic.

## [43] ternary conditional operator
- Added `cond ? a : b` with lowest precedence; nonzero cond selects a.
- New tokens `?` and `:`, `TernaryNode` AST, and evaluator support.

## [42] softplus
- Added `softplus(x)` = ln(1+e^x).

## [41] logistic sigmoid
- Added `logistic(x)` = 1/(1+e^-x).
- Prompts relocated from setup/prompts to root agents.md with new requirements.

## [40] asech/acsch/acoth inverse hyperbolic reciprocal
- Added `asech(x)`, `acsch(x)`, `acoth(x)`.
- Domain errors for out-of-range inputs.
- Added agents.md documenting the perpetual feature loop.

## [39] sech/csch/coth hyperbolic reciprocal
- Added `sech(x)`, `csch(x)`, `coth(x)`.
- Domain errors at csch(0) and coth(0).

## [38] asec/acsc/acot inverse reciprocal trig
- Added `asec(x)`, `acsc(x)`, `acot(x)`.

## [37] sec/csc/cot reciprocal trig
- Added `sec(x)`, `csc(x)`, `cot(x)`.
- Domain errors at undefined points instead of silent infinities.
- Added docs/operations.md listing all supported operations.

## [36] sinc cardinal sine
- Added `sinc(x)` = sin(x)/x with sinc(0)=1.

## [35] exp10 base-10 exponent
- Added `exp10(x)` = 10^x.

## [34] rsqrt reciprocal square root
- Added `rsqrt(x)` = 1/sqrt(x), domain x > 0.

## [33] state-aware one-shot eval/file
- `--eval`/`--file` load and save `--state` so scripts share variables.

## [32] --verify self-test battery
- Runs known-good expressions through parser + formatting; non-zero exit on failure.

## [31] '^' exponent operator
- Right-associative `^` binds tighter than `*` (2 * 3^2 == 18).

## [30] persist display settings
- deg/sci/eng/prec survive restarts in the state file.

## [29] --eng engineering notation
- One-shot output uses multiples-of-3 exponents.

## [28] --prec significant digits
- Sets display precision for one-shot eval output.

## [27] --sci scientific notation
- Toggles scientific formatting for one-shot eval.

## [26] --deg degree-mode trig
- `--eval --deg` evaluates trig in degrees.

## [25] 'last' command
- Prints the most recent expression and result.

## [24] 'redo' paired with undo
- Undo pushes a redo stack; 'redo' restores; new statements clear it.

## [23] 'undo' command
- Reverts the last statement.

## [22] 'history' shows results
- `history` lists `expr = result` pairs.

## [21] 'vars' command
- Lists defined variables with values.

## [20] 'clear' command
- Resets variables and ans.

## [19] gamma function
- Added `gamma(x)`.

## [18] erf/erfc
- Added `erf(x)` and `erfc(x)`.

## [17] atan2
- Added two-argument `atan2(y, x)`.

## [16] Bessel jn/yn
- Added `jn(n, x)` and `yn(n, x)`.

## [15] gcd/lcm
- Added `gcd(a, b)` and `lcm(a, b)`.

## [14] pow
- Added `pow(x, y)`.

## [13] min/max variadic
- Added variadic `min(...)` and `max(...)`.

## [12] log2/log10
- Added `log2(x)` and `log10(x)`.

## [11] round/trunc
- Added `round(x)` and `trunc(x)`.

## [10] floor/ceil
- Added `floor(x)` and `ceil(x)`.

## [9] inverse hyperbolic functions
- Added `asinh(x)`, `acosh(x)`, `atanh(x)`.

## [8] hyperbolic functions
- Added `sinh(x)`, `cosh(x)`, `tanh(x)`.

## [7] inverse trig functions
- Added `asin(x)`, `acos(x)`, `atan(x)`.

## [6] trig functions
- Added `sin(x)`, `cos(x)`, `tan(x)`.

## [5] log family
- Added `ln(x)`, `log(x)`, `log1p(x)`, `exp(x)`.

## [4] math helpers
- Added `abs(x)`, `sign(x)`, `sqrt(x)`, `cbrt(x)`.

## [3] constants
- Added `pi` and `e`.

## [2] variables and assignments
- Added `x = expr` variable assignment and reuse.

## [1] initial parser + REPL
- Expression parser, evaluator, interactive shell.
