# ADR 0019: primecount(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

prime/nthprime/nextprime exist but there was no way to count primes up to
a bound (the prime-counting function pi(n)).

## Decision

Add `primecount(n)` = pi(n): a simple trial-division loop over primes <= n.
Domain error for n < 0.

## Consequences

- primecount(10)=4, primecount(100)=25, primecount(2)=1, primecount(0)=0.
- Verified through the parser-level verify harness.
