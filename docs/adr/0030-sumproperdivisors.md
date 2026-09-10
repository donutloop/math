# ADR 0030: sumproperdivisors(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

sigma(n) sums all divisors but the aliquot sum (proper divisors only), the
basis of amicable and perfect number sequences, was missing.

## Decision

Add `sumproperdivisors(n)`: sum of proper divisors via the d/other pairing
loop, excluding n itself. Domain error for n < 2.

## Consequences

- sumproperdivisors(12)=16, sumproperdivisors(6)=6, sumproperdivisors(28)=28.
- Verified through the parser-level verify harness.
