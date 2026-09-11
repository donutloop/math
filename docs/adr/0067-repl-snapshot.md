# ADR-0067: REPL snapshot command

## Status
Accepted

## Context
Driving agents in an interactive machine REPL need a runtime state dump
without leaving the session. The CLI has --snapshot but the REPL had no
equivalent command.

## Decision
Add a `snapshot` REPL command in the calculator's handle switch: it emits
`calc.Snapshot()` ({vars, mode, prec, base}) as one JSON object, mirroring
CLI --snapshot. TestREPLSnapshot drives the REPL and parses the state.

## Consequences
REPL and CLI state introspection stay in lockstep; one feature, one ADR
(agents.md rule).
