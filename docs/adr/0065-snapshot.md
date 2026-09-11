# ADR-0065: --snapshot runtime state dump

## Status
Accepted

## Context
Agents introspecting a session need one consolidated, parseable view of
runtime state (variables, angle mode, precision, radix) rather than probing
separate commands.

## Decision
Add `--snapshot`: `calc.Snapshot()` returns `{vars, mode, prec, base}` and
main emits it as JSON. It works in the shared construction path (before the
eval/REPL split) so `--snapshot` is valid regardless of other flags. The
schema CLI list advertises it; TestSnapshot parses the JSON keys.

## Consequences
Agents get one-call runtime introspection; one feature, one ADR
(agents.md rule).
