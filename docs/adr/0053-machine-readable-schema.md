# ADR-0053: Machine-readable JSON schema

## Status
Accepted

## Context
Agents and scripts need to discover the full language surface (functions,
constants, commands, operators, modes, CLI flags) without scraping prose help.

## Decision
Add a `schema` REPL command and `--schema` CLI flag that emit a single
self-describing JSON document built from `parser.SupportedFunctions` (arity +
help), `parser.SupportedConstants` (value + help), plus commands/operators/
modes/CLI lists.

## Consequences
Agents get deterministic discovery; the schema is the source of truth for the
interface contract. Schema self-check wired into `--verify`.
