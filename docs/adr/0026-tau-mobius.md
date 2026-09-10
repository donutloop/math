# ADR 0026: tau(n) and mobius(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

The number-theory set lacked the divisor-count and Moebius functions.

## Decision

Add `tau(n)` = count of positive divisors (pairing d and n/d) and
`mobius(n)` = 1/-1/0 depending on squarefreeness and parity of distinct prime
factors. Domain error for n < 1.

## Consequences

- tau(12)=6, tau(6)=4.
- mobius(6)=1, mobius(30)=-1, mobius(12)=0, mobius(1)=1.
- Verified through the parser-level verify harness.
