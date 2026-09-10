# ADR 0014: Range min/max/avg

**Status:** Accepted
**Date:** 2026-09-10

## Context

sum/prod/count already iterate over a range. min/max/avg only existed as
variadic parser functions, so computing the extreme/mean of a range required
manual listing.

## Decision

Extend the range-loop macro (expandRangeLoop) to accept kinds min/max/avg:
min/max join the terms as variadic calls, avg joins the sum divided by the
term count. Disambiguation: a min/max/avg call is a range loop only when the
first argument is a valid identifier (loop variable) and the call has 4-5
arguments; otherwise it falls through to the variadic parser form.

## Consequences

- min(i, 1, 4, i*i)=1, max(i, 1, 4, i*i)=16, avg(i, 1, 3, i)=2.
- Variadic calls like avg(1, 2, 3, 4)=2.5 still work.
- Verified by unit tests (calc/calculator_test.go TestRangeLoops).
