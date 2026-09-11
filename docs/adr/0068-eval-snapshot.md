# ADR-0068: --eval --snapshot clean post-eval state

## Status
Accepted

## Context
Agents wanting evaluate-then-introspect got mixed output: eval results
printed prose/JSON THEN the state JSON appended, so `--eval --snapshot` was
not one parseable document.

## Decision
When `--snapshot` is set with `--eval`, main suppresses the eval result print
and emits ONLY the clean post-eval `Snapshot()` JSON ({vars, mode, prec,
base}). Agents get one parseable document per call. TestEvalSnapshot drives
`--eval x=5 --snapshot` and asserts the state JSON contains x=5.

## Consequences
Eval-and-state introspection is a single clean machine call; one feature,
one ADR (agents.md rule).
