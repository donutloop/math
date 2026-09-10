# ADR 0040: derangements(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

The combinatorics set lacked the subfactorial (derangement) sequence.

## Decision

Add `derangements(n)` = !n with the rolling recurrence
!n = (n-1)*(!(n-1)+!(n-2)), seeded !0=1, !1=0. Domain error for n < 0.

## Consequences

- derangements(2)=1, derangements(3)=2, derangements(4)=9, derangements(5)=44.
- Verified through the parser-level verify harness.
