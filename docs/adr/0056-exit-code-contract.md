# ADR-0056: Deterministic exit codes + machine-readable version

## Status
Accepted

## Context
Scripting agents need deterministic, documented exit codes and a parseable
version endpoint for compatibility checks. Previously eval/IO/usage errors all
collapsed to exit 1, and `--version` printed a prose string.

## Decision
Define a stable exit-code contract: `0` success, `1` usage/flag error, `2` IO
error (file-not-found, marshal/write), `3` evaluation error. Classify every
`os.Exit` site accordingly. Expose the contract in the schema (`exit_codes`),
and make `--version` emit a single valid JSON document (name, version, schema,
features, exit_codes).

## Consequences
Agents can branch deterministically on exit codes and parse version JSON for
compatibility checks. Contract documented in README/docs and self-checked by
`--verify`.
