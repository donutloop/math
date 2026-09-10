# ADR 0009: Variadic gcd/lcm

**Status:** Accepted
**Date:** 2026-09-10

## Context

`min`/`max`/`avg` already accept many arguments, but `gcd`/`lcm` were
binary-only. Number-theory composition (e.g. a shared divisor across several
terms) required nesting `gcd(gcd(a, b), c)`.

## Decision

Make `gcd` and `lcm` variadic (2+ args), folding left-to-right over the
argument list in the parser dispatch, and mark their arity as -1 in the
arity table so the parser accepts any number of arguments.

## Consequences

- `gcd(12, 18, 24)` = 6 and `lcm(4, 6, 10)` = 60 work directly.
- Binary calls like `gcd(36, 24)` remain valid.
- Verified through the parser-level verify harness (calc/verify.go).
