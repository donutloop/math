# ADR 0016: Range gcd/lcm

**Status:** Accepted
**Date:** 2026-09-10

## Context

gcd/lcm are variadic parser functions and range loops cover sum/prod/count/
min/max/avg. Extending iteration to gcd/lcm completes the range-family.

## Decision

Extend expandRangeLoop with gcd/lcm kinds that join the terms as variadic
gcd/lcm calls. Disambiguation mirrors min/max/avg: a call is a range loop only
when the first argument is a loop variable and the call has 4-5 args.

## Consequences

- gcd(i, 1, 10, i)=1, lcm(i, 1, 5, i)=60, lcm(i, 2, 4, i)=12.
- Variadic calls like gcd(12, 18)=6 still work.
- Verified by unit tests (calc/calculator_test.go TestRangeLoops).
