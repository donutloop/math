# ADR 0017: lambertw(x)

**Status:** Accepted
**Date:** 2026-09-10

## Context

The calculator has many special functions but lacked the Lambert W function,
which is fundamental for solving exponential/log equations.

## Decision

Add `lambertw(x)` computing the principal branch via Newton iteration with
up to 50 iterations and a 1e-14 convergence tolerance. Domain error for
x < -1/e.

## Consequences

- lambertw(0)=0, lambertw(1)=0.56714329 (omega), lambertw(e)=1.
- Verified by a tolerance unit test (calc/lambertw_test.go).
