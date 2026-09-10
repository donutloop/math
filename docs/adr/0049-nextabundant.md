# ADR 0049: nextabundant(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

isabundant tested a single value but there was no way to find the next
abundant number.

## Decision

Add `nextabundant(n)`: scan upward testing sumProperDivisors(n) > n.
Domain error for n < 1.

## Consequences

- nextabundant(7)=12, nextabundant(12)=12, nextabundant(18)=18.
- Verified through the parser-level verify harness.
