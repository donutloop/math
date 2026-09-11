# ADR-0079: break / continue in while loops (math programming language)

## Status
Accepted

## Context
A math programming language needs loop control flow beyond the condition
check, so loops can exit early or skip a body step.

## Decision
Add direct `break` and `continue` statements to while-loop bodies. When a
body statement (after paren-aware splitting) is exactly `break`, the loop
exits immediately with the last body value; `continue` skips the rest of the
current body iteration. `begin(...)` blocks also honor them for their own
statement sequence. The iteration bound guard still applies.

## Consequences
Loops get early-exit/skip control. Nested begin-block propagation of sentinels
is deferred; direct body statements are the supported form. One ADR per
feature per the agents.md rule.
