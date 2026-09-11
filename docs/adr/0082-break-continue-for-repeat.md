# ADR-0082: break / continue propagate through for and repeat loops

## Status
Accepted

## Context
ADR-0081 added break/continue propagation through `begin(...)` blocks for
`while` loops. `for` (range) and `repeat` loops still ignored them: their
bodies were treated as a single expression (no statement splitting), so a
direct `break`/`continue` or a sentinel from a begin block was mis-parsed.

## Decision
`expandFor` and `expandRepeat` now split loop bodies into statements
(paren-aware), substitute the loop variable into each statement, and process
each one like `expandWhile`: direct `break` exits the loop, `continue` skips
the rest of the current iteration, and sentinels raised by `begin(...)` blocks
are caught at string-expansion level before parsing. Both loops return the
last statement value.

## Consequences
break/continue work uniformly across while/for/repeat, including inside
begin(...) blocks. Unit tests (calc/for_test.go) and whole-program integration
tests cover both loops and both propagation paths. One ADR per feature per the
agents.md rule.
