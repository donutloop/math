# ADR 0029: carmichael(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

totient gives Euler's phi but the related Carmichael lambda function was
missing.

## Decision

Add `carmichael(n)`: lcm of lambda over the prime-power factors p^k, where
lambda(p^k)=phi(p^k)=p^(k-1)(p-1) except for 2^k with the special values
lambda(2)=1, lambda(4)=2, lambda(2^k)=2^(k-2) for k>=3. Uses the existing
lcm helper. Domain error for n < 1.

## Consequences

- carmichael(9)=6, carmichael(12)=2, carmichael(10)=4, carmichael(1)=1.
- Verified through the parser-level verify harness.
