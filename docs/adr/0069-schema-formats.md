# ADR-0069: schema formats field

## Status
Accepted

## Context
Agents discovering output formats probed separate flags; the schema language
surface listed flags but not the machine formats themselves.

## Decision
Add `formats` to the Schema JSON: `["text", "json", "jsonl", "csv"]`, so
agents discover all output formats from `--schema` in one call. TestSchemaFormats
parses the list.

## Consequences
Format discovery is self-describing; one feature, one ADR (agents.md rule).
