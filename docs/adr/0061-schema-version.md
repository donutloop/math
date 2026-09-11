# ADR-0061: Schema version pinning for agents

## Status
Accepted

## Context
Agents need a deterministic compatibility signal: when the machine surface
changes (new flags, formats, schema fields), a pinned version lets agents
detect breaking changes before they rely on new output.

## Decision
Bump `schemaVersion` to 1.1.0 to reflect the expanded machine surface
(--jsonl, --output, --verify --json). Export `calc.SchemaVersion()` so
--version and --schema report the same pinned version. --version's features
list advertises every machine format. Agents compare version before use.

## Consequences
One version number tracks surface changes; --version/--schema stay in
lockstep; one feature, one ADR (agents.md rule).
