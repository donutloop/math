# ADR 0028: liouville(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

bigomega(n) gives the prime-factor count but the parity-based Liouville
function was missing.

## Decision

Add `liouville(n)` = (-1)^bigomega(n) using the existing bigomega helper.
Domain error for n < 1.

## Consequences

- liouville(12)=-1, liouville(8)=-1, liouville(6)=1, liouville(1)=1.
- Verified through the parser-level verify harness.
