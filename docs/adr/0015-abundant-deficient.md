# ADR 0015: isabundant(n) / isdeficient(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

sigma(n) and isperfect(n) exist. The abundant (>2n) and deficient (<2n)
classifications complete the divisor-sum classification trio.

## Decision

Add `isabundant(n)` (sigma(n) > 2n) and `isdeficient(n)` (sigma(n) < 2n) as
parser-level math functions, one integer argument each, domain error for n<=0.

## Consequences

- isabundant(12)=1, isdeficient(8)=1, isabundant(6)=0 (6 is perfect).
- Verified through the parser-level verify harness.
