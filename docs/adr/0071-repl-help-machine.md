# ADR-0071: Structured REPL help in machine modes

## Status
Accepted

## Context
Driving agents in a machine REPL typed `help` and got prose, breaking
session parseability. Machine modes need a structured command list.

## Decision
In jsonl/json/csv modes, printHelp emits `{kind:"help",commands:[...]}` (or
CSV rows) so agents discover REPL commands programmatically; prose stays for
text mode. TestREPLHelpMachine drives jsonl+help and parses the object.

## Consequences
REPL machine sessions are fully parseable including help; one feature, one
ADR (agents.md rule).
