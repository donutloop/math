# ADR 0050: nextdeficient(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

isdeficient tested a single value but there was no way to find the next
deficient number.

## Decision

Add `nextdeficient(n)`: scan upward testing sumProperDivisors(n) < n.
Domain error for n < 1.

## Consequences

- nextdeficient(5)=5, nextdeficient(6)=7, nextdeficient(12)=13.
- Verified through the parser-level verify harness.
