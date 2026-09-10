# ADR 0032: loggamma(x)

**Status:** Accepted
**Date:** 2026-09-10

## Context

gamma(x) existed via math.Gamma but the log-gamma function ln(gamma(x))
was missing and useful in its own right (e.g. for beta/likelihood work).

## Decision

Add `loggamma(x)` with the recurrence lnGamma(x)=lnGamma(x+1)-ln(x) into
the large-x regime plus the Stirling expansion to the x^5 term. Domain error
at nonpositive integers.

## Consequences

- loggamma(1)=0, loggamma(5)=ln24 (~1e-10 error), loggamma(0.5)=ln(sqrt(pi)).
- Verified by tolerance unit test (calc/loggamma_test.go).
