# ADR 0021: digamma(x)

**Status:** Accepted
**Date:** 2026-09-10

## Context

The calculator lacked the psi (digamma) function, fundamental to special
function identities and harmonic-number generalizations.

## Decision

Add `digamma(x)` implemented with the recurrence psi(x) = psi(x+1) - 1/x
into the large-x regime plus an asymptotic expansion to the x^12 Bernoulli
term. Domain error at nonpositive integers.

## Consequences

- digamma(1) = -0.577215664901531 (~ -gamma, 2e-15 error).
- digamma(0.5) = -1.96351002602142 (~ -gamma-2ln2, 3e-15 error).
- Verified by tolerance unit test (calc/digamma_test.go).
