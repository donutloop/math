# ADR-0062: REPL jsonl (NDJSON) mode

## Status
Accepted

## Context
Driving agents interact with the REPL and need line-delimited, parseable
results. The REPL already had json (one object per result) and csv modes but
no streaming NDJSON mode, so agents could not ingest one line at a time.

## Decision
Add a `jsonl` REPL command and a `jsonlMode` flag on the calculator. When
active, every result is emitted as one `{expr,kind,value}` JSON object per
line, matching the CLI `--jsonl` shape. `jsonl` disables json/csv modes.
TestREPLJSONL drives the REPL and parses every `{`-prefixed line.

## Consequences
REPL and CLI machine formats stay in lockstep; one feature, one ADR
(agents.md rule).
