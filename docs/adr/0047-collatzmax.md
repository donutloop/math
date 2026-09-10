# ADR 0047: collatzmax(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

collatz(n) gave the step count to 1 but the peak value reached in the
sequence was missing.

## Decision

Add `collatzmax(n)`: track the maximum n encountered while applying the
Collatz map. Domain error for n < 1.

## Consequences

- collatzmax(3)=16, collatzmax(6)=16, collatzmax(1)=1.
- Verified through the parser-level verify harness.
