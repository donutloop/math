# ADR-0060: --output unified format selector

## Status
Accepted

## Context
Agents discover output formats by probing separate flags (--json, --jsonl,
--csv). A single `--output <format>` selector collapses discovery: one flag
names the machine format (text|json|jsonl|csv).

## Decision
Add `--output text|json|jsonl|csv` which maps onto the existing json/jsonl/csv
branches (text = default prose). Unknown formats exit 1 (usage). The schema
CLI list advertises it; TestOutputSelector proves it equals the individual
flags. The per-format flags remain for back-compat.

## Consequences
Agents use one flag; behavior identical to existing flags; one feature, one
ADR (agents.md rule).
