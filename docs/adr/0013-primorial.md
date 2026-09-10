# ADR 0013: primorial(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

Number theory already covers primes (isprime, prime, nextprime) and divisor
arithmetic. The primorial (product of the first n primes) is a natural
companion used in proofs and bounds.

## Decision

Add `primorial(n)` as a parser-level math function: product of the first n
primes, with a domain error for n <= 0.

## Consequences

- primorial(1)=2, primorial(2)=6, primorial(3)=30, primorial(4)=210.
- Verified through the parser-level verify harness.
