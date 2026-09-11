# ADR-0072: CLI --eval --jsonl structured error contract

## Status
Accepted

## Context
CLI `--eval --jsonl` on a failing expression emitted `{kind:"error",value:...}`
and exit 0 — the error key and exit code mismatched the REPL jsonl shape and
the documented 3=eval contract, so agents could miss failures.

## Decision
printJSONL emits the `error` key (not `value`) for `kind:"error"`, matching the
REPL jsonl shape. The eval jsonl branch tracks eval failures and exits 3
(ExitEval) when any expression failed. TestEvalJSONLStructuredError builds the
binary, asserts exit 3 and the {kind:"error",error} object.

## Consequences
CLI and REPL jsonl error objects share one shape; eval failures exit 3 per the
documented contract; one feature, one ADR (agents.md rule).
