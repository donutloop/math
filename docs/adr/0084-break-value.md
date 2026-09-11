# ADR-0084: break <value> returns an explicit loop result

## Status
Accepted

## Context
Loops returned only their last body value on early exit via bare `break`.

## Decision
Add `break <expr>`: a loop exits and returns the value of `<expr>` (loop
variable already substituted). Direct body statements and begin(...) blocks
both support it — begin blocks raise a prefixed sentinel
(`__BREAK_VALUE__:<expr>`) that the enclosing loop catches before parsing.

## Consequences
Loops can return a computed result on early exit, useful for early-termination
search/accumulation. Unit and whole-program integration tests cover
while/for/repeat and the begin-wrapped form. One ADR per feature per the
agents.md rule.
