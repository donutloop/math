# ADR 0051: nextperfect(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

isperfect tested a single value but there was no way to find the next
perfect number.

## Decision

Add `nextperfect(n)`: scan upward testing sumProperDivisors(n) == n.
Domain error for n < 1.

## Consequences

- nextperfect(1)=6, nextperfect(7)=28, nextperfect(29)=496.
- Verified through the parser-level verify harness.
