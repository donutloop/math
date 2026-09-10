# ADR 0043: doublefactorial(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

fact existed but the double (semi-) factorial was missing.

## Decision

Add `doublefactorial(n)` = n!! = product of n, n-2, ... down to 2 (or 1),
with 0!! = 1. Domain error for n < 0.

## Consequences

- doublefactorial(0)=1, doublefactorial(4)=8, doublefactorial(5)=15,
  doublefactorial(6)=48.
- Verified through the parser-level verify harness.
