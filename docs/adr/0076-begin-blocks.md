# ADR-0076: begin(...) statement blocks (math programming language)

## Status
Accepted

## Context
The language needed a general sequencing construct so multi-step logic can
appear anywhere (standalone or inside loop bodies), not only at top level.

## Decision
Add begin(...): the statements inside are evaluated in sequence (paren-aware
split on ';'), each as an assignment or expression, and the last value is
returned. Reuses the same assign/eval path as while bodies.

## Consequences
Statement blocks are a first-class language construct: usable standalone,
as loop bodies, and nested. One ADR per feature per the agents.md rule.
