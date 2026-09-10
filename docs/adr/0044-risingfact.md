# ADR 0044: risingfact(n, k)

**Status:** Accepted
**Date:** 2026-09-10

## Context

perm(n,k) gave the falling factorial but the rising (Pochhammer) factorial
was missing.

## Decision

Add `risingfact(n, k)` = n(n+1)...(n+k-1). Domain error for k < 0.

## Consequences

- risingfact(1,3)=6, risingfact(2,3)=24, risingfact(5,0)=1.
- Verified through the parser-level verify harness.
