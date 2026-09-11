# ADR-0058: --jsonl NDJSON output

## Status
Accepted

## Context
Single-doc JSON (ADR-0054) gives agents one parseable array, but streaming
agents and log-style consumers want one JSON object per line: they can ingest
each line incrementally without buffering a whole document.

## Decision
Add `--jsonl` which emits one JSON object per line, each with the stable
`{expr, kind, value}` shape (kind ∈ value|assign|var|error), mirroring the
JSON branch. A `printJSONL` helper marshals each result; `--vars` appends one
`var` line per defined variable. Schema CLI list and `--verify` self-check
advertise/validate the format.

## Consequences
Streaming agents get line-delimited, machine-consumable output. Format is
stable and self-checked; one feature, one ADR (agents.md rule).
