# ADR-0057: --help flag (deterministic usage, exit 0)

## Status
Accepted

## Context
The schema advertised `--help` but no such flag existed — a real agent hazard
(agents would call `--help` and get an unknown-flag error, exit 2). Go's flag
package also exits 2 on usage errors, violating the documented exit-code
contract (usage=1).

## Decision
Register a real `--help` flag that prints usage to stdout and exits 0. Handle
`flag.ErrHelp` explicitly (exit 0, not usage=1). Usage output documents the
machine-readable outputs and the exit-code contract.

## Consequences
Agents get deterministic help (exit 0, stdout-only) and the exit-code contract
holds for all usage paths.
