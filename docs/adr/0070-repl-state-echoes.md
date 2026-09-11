# ADR-0070: Structured REPL state echoes

## Status
Accepted

## Context
Driving agents parsing REPL machine output hit prose for state commands:
`prec`, `base`, `deg`/`rad`/`grad` printed `prec = 5`, `base = 2`, `deg`.
State changes were not machine-parseable.

## Decision
Add `emitState(name, value)` on the calculator: in jsonl/json/csv modes it
emits a structured object (`{state,value}` for json/jsonl, `state,<name>,<v>`
for csv); prose stays for text mode. prec/base/mode commands use it.
TestREPLStateEchoes drives prec/base/deg in jsonl and parses the objects.

## Consequences
REPL machine sessions are fully parseable including state changes; one
feature, one ADR (agents.md rule).
