# ADR-0054: Single-document JSON output

## Status
Accepted

## Context
`--json` emitted line-delimited fragments, forcing agents to scrape stdout.

## Decision
Emit one valid, parseable JSON array (expr/assign/var/error objects) so agents
can `json.load` the whole stdout at once. Errors are captured inline per
expression rather than aborting the run.

## Consequences
Deterministic, machine-parseable output; agents get a stable JSON contract.
