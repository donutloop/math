# ADR 0011: isperfect(n) perfect-number check

**Status:** Accepted
**Date:** 2026-09-10

## Context

sigma(n) (sum of divisors) was added in ADR 0010. Perfect numbers are the
classic application (sigma(n) == 2n), but composing that check inline is
awkward.

## Decision

Add `isperfect(n)` returning 1 when sigma(n) == 2n and 0 otherwise, with a
domain error for n <= 0. Mirrors sigma's dispatch/arity pattern.

## Consequences

- isperfect(6)=1, isperfect(28)=1, isperfect(10)=0.
- Verified through the parser-level verify harness.
