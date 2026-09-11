# ADR-0073: while loop construct (math programming language)

## Status
Accepted

## Context
The language had range-based loops (sum/prod/count/repeat) but no general
control flow with mutable state. Steering toward a math programming language
requires a general loop construct.

## Decision
Add while(cond, body): evaluate cond with current variable bindings; while
nonzero, evaluate body as a statement so assignments (x = x + 1) mutate
variables. Reuse c.assign / c.eval so mutation works. A hard 10000-iteration
bound guards against non-progressing bodies.

## Consequences
General mutable-state loops become available (counters, accumulation,
search). The language moves from pure range-expansion toward real control
flow. One ADR per feature per the agents.md rule.
