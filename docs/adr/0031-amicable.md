# ADR 0031: amicable(a, b)

**Status:** Accepted
**Date:** 2026-09-10

## Context

sumproperdivisors enables aliquot analysis but no direct amicable-pair
predicate existed.

## Decision

Add `amicable(a, b)` returning 1/0: a and b are amicable iff a != b and
sumProperDivisors(a)=b and sumProperDivisors(b)=a. Domain error for a, b < 2.

## Consequences

- amicable(220, 284)=1, amicable(10, 5)=0, amicable(6, 6)=0.
- Verified through the parser-level verify harness.
