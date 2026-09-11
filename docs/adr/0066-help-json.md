# ADR-0066: --help --json structured flag contract

## Status
Accepted

## Context
Agents need to discover the CLI contract programmatically — flag names,
types, help text, exit codes — without scraping prose usage output.

## Decision
When `--help --json` is set, main emits a structured document:
`{version, flags:[{name,type,help}], exit_codes:{ok,usage,io,eval}}` by
introspecting flag.CommandLine via VisitAll and reflect on flag Value types.
Prose --help is unchanged. TestHelpJSON parses it.

## Consequences
Agents discover the full CLI surface in one call; one feature, one ADR
(agents.md rule).
