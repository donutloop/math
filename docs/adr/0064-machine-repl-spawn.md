# ADR-0064: Spawn a machine REPL from the CLI

## Status
Accepted

## Context
Driving agents want an interactive machine session in one command: pass a
format flag and get parseable output from the first line, with no prose
banner to scrape around.

## Decision
When a format flag (`--jsonl`, `--json`, `--csv`, or `--output`) is given
WITHOUT `--eval`, main presets the REPL via `calc.SetMachineMode(mode)`:
it sets the output mode and suppresses the prose banner. `SetMachineMode`
returns an error for unknown modes (exit 1). TestSetMachineMode drives a
preset REPL and asserts clean NDJSON results.

## Consequences
Agents spawn an interactive machine session with one command; banner stays
clean; one feature, one ADR (agents.md rule).
