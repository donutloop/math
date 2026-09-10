# ADR 0045: multinomial(n, ...)

**Status:** Accepted
**Date:** 2026-09-10

## Context

choose(n,k) gave binomial coefficients but the multinomial generalization
was missing.

## Decision

Add variadic `multinomial(n, a, b, ...)` computed as the chained product of
binomial coefficients choose(rem, a) for each group; groups must sum to n.
Domain error unless groups are non-negative and sum to n.

## Consequences

- multinomial(5, 2, 3)=10, multinomial(6, 2, 2, 2)=90.
- Verified through the parser-level verify harness.
