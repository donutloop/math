# ADR-0080: whole-program integration tests (math programming language)

## Status
Accepted

## Context
Single-expression parser tests did not exercise the language as a complete
program (variable setup, function definitions, loops, begin blocks, scope).

## Decision
Add `tests/integration/programs_test.go`: each case runs a whole math program
through the batch calculator runtime and asserts exact final result lines.
Programs cover factorial loops, range for-loops, recursive functions,
function-local scope, paren-aware multi-statement bodies, and composition.

## Consequences
The language is validated end-to-end as a runtime, not just expression by
expression. One ADR per feature per the agents.md rule.
