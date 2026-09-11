# ADR-0083: clear loop-control errors + nested-loop break/continue

## Status
Accepted

## Context
Top-level `begin(...; break)` raised an internal sentinel that failed parsing
with a confusing "unknown function" error. Nested loops' break/continue scope
was unverified.

## Decision
`evalExpanded` detects sentinels before parsing and returns clear errors
("break outside a loop" / "continue outside a loop") when a break/continue
escapes to the top level. Whole-program tests verify nested semantics: an
inner begin break exits only the inner loop, and inner continue skips only
the inner body (outer accumulators still run).

## Consequences
Programmers get actionable diagnostics for misplaced break/continue, and
nested-loop control scope is locked by tests. One ADR per feature per the
agents.md rule.
