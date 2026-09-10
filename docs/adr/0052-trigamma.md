# ADR 0052: trigamma(x)

**Status:** Accepted
**Date:** 2026-09-10

## Context

digamma(x) gave psi but the next polygamma derivative psi' was missing.

## Decision

Add `trigamma(x)` = psi'(x) with the recurrence psi'(x)=psi'(x+1)+1/x^2 into
the large-x regime plus the asymptotic expansion to the x^7 term.
Domain error at nonpositive integers.

## Consequences

- trigamma(1)=pi^2/6 (1.6449), trigamma(2)=pi^2/6-1 (0.6449).
- Verified by tolerance unit test (calc/trigamma_test.go).
