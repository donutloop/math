# ADR 0022: repeat(n, var, body) loop macro

**Status:** Accepted
**Date:** 2026-09-10

## Context

Range loops (sum/prod/count/min/max/avg/gcd/lcm) aggregate over a range but
there was no way to simply iterate and keep the final value.

## Decision

Add `repeat(n, var, body)`: a macro that substitutes var for each integer
1..n, re-expands body, and returns the last evaluated value. Count must be
>= 1; nested repeat calls are supported (the outer variable substitutes into
the inner call before inner expansion).

## Consequences

- repeat(5, i, i*i)=25 (last value).
- repeat(3, i, repeat(3, j, i+j))=6 (nested, last = 3+3).
- Verified by unit test (calc/repeat_test.go).
