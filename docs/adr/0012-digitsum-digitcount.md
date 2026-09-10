# ADR 0012: digitsum(n) and digitcount(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

Digit-level arithmetic (cross-sums, digit lengths) is common in number
theory and recreational math but requires manual loops.

## Decision

Add `digitsum(n)` (sum of decimal digits) and `digitcount(n)` (number of
decimal digits) as parser-level math functions with one integer argument.
digitsum(0)=0 and digitcount(0)=1.

## Consequences

- digitsum(1234)=10, digitcount(1234)=4, digitcount(0)=1.
- Verified through the parser-level verify harness.
