# ADR 0036: nthdigit(n, k)

**Status:** Accepted
**Date:** 2026-09-10

## Context

digitsum and digitcount existed but there was no way to extract a single
digit at a given position.

## Decision

Add `nthdigit(n, k)`: divide by 10^k and take the units digit. Domain error
for n < 0 or k < 0.

## Consequences

- nthdigit(12345,0)=5, nthdigit(12345,1)=4, nthdigit(12345,4)=1.
- Verified through the parser-level verify harness.
