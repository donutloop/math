# ADR 0042: digitproduct(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

digitsum and digitcount existed but the product of the digits was missing.

## Decision

Add `digitproduct(n)`: multiply the decimal digits (0 stays 0).
Domain error for n < 0.

## Consequences

- digitproduct(123)=6, digitproduct(50)=0, digitproduct(9)=9, digitproduct(0)=0.
- Verified through the parser-level verify harness.
