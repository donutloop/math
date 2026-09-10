# ADR 0010: totient and sigma number-theory functions

**Status:** Accepted
**Date:** 2026-09-10

## Context

Number theory already covers primality (isprime), divisors (divcount),
factorization-ish helpers, and sequences (collatz). Missing the two most
fundamental arithmetic functions: Euler's totient and the divisor-sum.

## Decision

Add `totient(n)` (count of integers 1..n coprime to n) and `sigma(n)` (sum of
positive divisors of n) as parser-level math functions, one integer argument
each, with a typed domain error for n <= 0 (totient) and n == 0 (sigma).

## Consequences

- `totient(12)` = 4 and `sigma(6)` = 12 work directly.
- Verified through the parser-level verify harness (calc/verify.go).
