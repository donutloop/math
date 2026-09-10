# ADR 0034: harmonic(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

The discrete-special-function set lacked the harmonic numbers.

## Decision

Add `harmonic(n)` = sum 1/k for k=1..n as a float loop. Domain error for n<1.

## Consequences

- harmonic(1)=1, harmonic(3)=1.8333, harmonic(10)=2.9289.
- Verified by tolerance unit test (calc/harmonic_test.go).
