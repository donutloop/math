# ADR-0078: function-local variable scope (math programming language)

## Status
Accepted

## Context
Assignments inside user-function bodies wrote global variables, so internal
accumulators leaked between calls (f(5) then f(10) continued the same acc).

## Decision
Wrap user-function body expansion in a variable snapshot/restore: capture
c.vars, expand the body, then restore. Internal assignments are discarded on
return; recursion is safe because each call snapshots independently.

## Consequences
Functions are pure/local; accumulators do not leak. One ADR per feature per
the agents.md rule.
