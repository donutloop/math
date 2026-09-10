# ADR 0038: catalan(n)

**Status:** Accepted
**Date:** 2026-09-10

## Context

Bell numbers were added but Catalan numbers, the other standard combinatorics
sequence, were missing.

## Decision

Add `catalan(n)` = choose(2n,n)/(n+1) using the existing choose helper.
Domain error for n < 0.

## Consequences

- catalan(0)=1, catalan(3)=5, catalan(4)=14, catalan(5)=42.
- Verified through the parser-level verify harness.
