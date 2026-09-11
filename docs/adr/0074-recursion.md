# ADR-0074: recursion in user-defined functions (math programming language)

## Status
Accepted

## Context
User-function arguments were bound to raw expanded strings (f(n-1) became
f(4-1)), so recursive bodies misparsed (4-1 <= 1 read as 4 - (1<=1)). This
blocked recursion — a core programming-language capability.

## Decision
Evaluate each user-function argument to its numeric value (c.eval + %g)
before parameter substitution. Recursive calls then bind params to numbers,
keeping conditions and arithmetic in bodies well-formed. The existing
500-depth guard bounds runaway recursion.

## Consequences
Recursion works: factorial f(5)=120, f(10)=3628800; fibonacci verified.
One ADR per feature per the agents.md rule.
