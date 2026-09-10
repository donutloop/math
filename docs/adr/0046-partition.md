# ADR 0046: partition(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

bell(n) counts set partitions but the integer partition count p(n) was
missing.

## Decision

Add `partition(n)` = number of integer partitions of n, computed with the
pentagonal-number recurrence
p(n) = sum_k (-1)^(k+1) p(n - k(3k-1)/2) + p(n - k(3k+1)/2).
Domain error for n < 0.

## Consequences

- partition(0)=1, partition(4)=5, partition(5)=7, partition(6)=11.
- Verified through the parser-level verify harness.
