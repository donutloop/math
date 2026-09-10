# ADR 0041: lucas(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

fibonacci existed but the closely related Lucas sequence was missing.

## Decision

Add `lucas(n)` with L_0=2, L_1=1 and L_n = L_{n-1}+L_{n-2}.
Domain error for n < 0.

## Consequences

- lucas(0)=2, lucas(3)=4, lucas(4)=7, lucas(5)=11.
- Verified through the parser-level verify harness.
