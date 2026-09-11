# ADR-0085: break / continue inside if branches of loop bodies

## Status
Accepted

## Context
`if(cond, break, val)` in a loop body failed ("unknown function break"),
because expandIf expanded the branch as an ordinary expression.

## Decision
`expandIf` routes its selected branch through `expandBranch`: a bare
`break`/`continue` branch returns the sentinel, and `break <expr>` returns the
break-value prefix. `expandBegin`'s value-statement path now expands at string
level and forwards any nested sentinel up to the enclosing loop instead of
parsing it as an error.

## Consequences
Loop control works conditionally through if/else branches inside loop bodies.
Unit and whole-program integration tests cover break, continue, and
break-value in if branches. One ADR per feature per the agents.md rule.
