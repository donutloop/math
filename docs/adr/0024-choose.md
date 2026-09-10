# ADR 0024: choose(n, k) binomial coefficient

**Status:** Accepted
**Date:** 2026-09-10

## Context

fact and the combinatorics set lacked the binomial coefficient.

## Decision

Add `choose(n, k)` computed with the exact multiplicative formula
res * (n-k+i) / i, using min(k, n-k) for symmetry. Domain error unless
0 <= k <= n.

## Consequences

- choose(5,2)=10, choose(10,3)=120, choose(5,0)=1, choose(5,5)=1.
- Verified through the parser-level verify harness.
