# ADR 0020: prevprime(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

nextprime(n) finds the smallest prime above n but there was no way to find
the largest prime below n.

## Decision

Add `prevprime(n)`: scan downward from n-1 for the first prime. Domain error
for n <= 2 (no prime strictly below 2).

## Consequences

- prevprime(10)=7, prevprime(100)=97, prevprime(3)=2.
- Verified through the parser-level verify harness.
