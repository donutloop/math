# ADR-0077: for(var, lo, hi, body) range loop (math programming language)

## Status
Accepted

## Context
The loop family (sum/prod/count/repeat/while) lacked an explicit range
counter loop with mutable body state.

## Decision
Add for(var, lo, hi, body): bind var to each integer in [lo,hi] inclusive,
evaluate body as a statement (assignment or expression), return the last
value. lo/hi are evaluated up front; a descending range yields 0.

## Consequences
Range iteration with mutable state is a first-class construct. One ADR per
feature per the agents.md rule.
