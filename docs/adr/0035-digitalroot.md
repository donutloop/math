# ADR 0035: digitalroot(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

digitsum and digitcount existed but the iterated digital root was missing.

## Decision

Add `digitalroot(n)`: repeatedly sum decimal digits until a single digit
remains (0 stays 0). Domain error for n < 0.

## Consequences

- digitalroot(12345)=6, digitalroot(48)=3, digitalroot(9)=9, digitalroot(0)=0.
- Verified through the parser-level verify harness.
