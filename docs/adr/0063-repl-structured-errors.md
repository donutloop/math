# ADR-0063: Structured errors in REPL machine modes

## Status
Accepted

## Context
Driving agents parse REPL machine output, but eval failures still printed
prose `error: ...`, forcing scraping. Machine modes (jsonl/json/csv) need
parseable error objects so failures are first-class machine output.

## Decision
In the REPL, when jsonl/json/csv mode is active, an eval error emits a
structured object: jsonl `{expr,kind:"error",error}`, json `{error}`,
csv `expr,error,<msg>`; prose stays for text mode. TestREPLJSONLError drives
1/0 and asserts a structured error object appears.

## Consequences
REPL machine modes are fully parseable including failures; one feature, one
ADR (agents.md rule).
