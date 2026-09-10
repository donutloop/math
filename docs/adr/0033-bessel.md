# ADR 0033: besselj0(x) and besselj1(x)

**Status:** Accepted
**Date:** 2026-09-10

## Context

The special-function set lacked Bessel functions of the first kind.

## Decision

Add `besselj0(x)` and `besselj1(x)` with the convergent power series
(40 terms): J0(x)=sum (-1)^k/(k!)^2 (x/2)^(2k) and
J1(x)=sum (-1)^k/(k!(k+1)!) (x/2)^(2k+1).

## Consequences

- besselj0(0)=1, besselj1(0)=0, besselj0(1)=0.7651976866,
  besselj1(1)=0.4400505857 (1e-10 error).
- Verified by tolerance unit test (calc/bessel_test.go).
