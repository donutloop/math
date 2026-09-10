# ADR 0025: omega(n) and bigomega(n) prime-factor counts

**Status:** Accepted
**Date:** 2026-09-10

## Context

The number-theory set lacked the standard omega and Omega prime-factor
counting functions.

## Decision

Add `omega(n)` = number of distinct prime factors and `bigomega(n)` = total
prime factors with multiplicity, via trial-division factoring. Domain error
for n < 1.

## Consequences

- omega(30)=3, omega(12)=2.
- bigomega(12)=3, bigomega(8)=3.
- Verified through the parser-level verify harness.
