# ADR 0039: stirling(n, k)

**Status:** Accepted
**Date:** 2026-09-10

## Context

Bell numbers count all set partitions; the finer Stirling numbers by subset
count were missing.

## Decision

Add `stirling(n, k)` = number of partitions of n into k non-empty subsets,
via a rolling DP of the recurrence S(n,k)=k*S(n-1,k)+S(n-1,k-1).
Domain error unless 0 <= k <= n.

## Consequences

- stirling(3,2)=3, stirling(4,2)=7, stirling(4,3)=6, stirling(0,0)=1.
- Verified through the parser-level verify harness.
