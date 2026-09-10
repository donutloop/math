# ADR 0023: sumofprimes(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

primorial(n) gives the product of the first n primes but there was no way to
sum them.

## Decision

Add `sumofprimes(n)`: sum the first n primes via trial division. Domain error
for n < 1.

## Consequences

- sumofprimes(5)=28, sumofprimes(10)=129.
- Verified through the parser-level verify harness.
