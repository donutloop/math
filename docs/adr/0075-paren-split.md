# ADR-0075: paren-aware statement splitting (math programming language)

## Status
Accepted

## Context
Statements split on ';' at the top level, so multi-step loop bodies
(acc = acc + x; x = x + 1) were broken apart before the loop saw them.

## Decision
splitStatements tracks paren depth and splits only at depth 0. ';' inside
parentheses stays within the enclosing statement, letting while bodies and
future begin/end blocks sequence statements.

## Consequences
Multi-step while bodies work. Statement separation now respects block
structure. One ADR per feature per the agents.md rule.
