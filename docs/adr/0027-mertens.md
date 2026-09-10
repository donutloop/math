# ADR 0027: mertens(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

mobius(n) gives a single value but the summatory Mertens function was missing.

## Decision

Add `mertens(n)` = sum of mobius(i) for i=1..n, using the existing mobius
helper. Domain error for n < 1.

## Consequences

- mertens(10)=-1 (matches the known sequence), mertens(1)=1, mertens(6)=-1.
- Verified through the parser-level verify harness.
