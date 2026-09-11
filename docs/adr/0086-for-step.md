# ADR-0086: for loop with an explicit step

## Status
Accepted

## Context
Range loops could only step by 1; summing evens/odds required manual guards.

## Decision
`for(var, lo, hi, step, body)` accepts an optional 5th argument `step`. The
dispatch reads 5 args (var, lo, hi, step, body) or falls back to the 4-arg
form (step defaults to 1). `expandFor` evaluates step, rejects zero, and
advances the loop variable by `|step|`.

## Consequences
Range loops support arithmetic progression with an explicit stride. Unit and
whole-program integration tests cover even and odd stepping. One ADR per
feature per the agents.md rule.
