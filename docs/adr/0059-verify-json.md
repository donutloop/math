# ADR-0059: --verify --json structured health report

## Status
Accepted

## Context
Agents need deterministic, parseable health signals — exit codes alone carry
no detail. Prose --verify forces scraping. A structured JSON report lets an
agent confirm passed/failed counts and every check programmatically.

## Decision
When `--verify --json` is set, main emits `{"version", "passed", "failed",
"checks": [{check,pass,detail}]}` to stdout. verify.go refactors the battery
into runChecks() shared by both the prose Verify and VerifyJSON so reports
stay in lockstep. Exit contract is unchanged: failed>0 still exits 3.

## Consequences
Agents get a self-describing health report; prose output unchanged; one
feature, one ADR (agents.md rule).
