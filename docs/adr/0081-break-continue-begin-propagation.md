# ADR-0081: break / continue propagate through begin(...) blocks

## Status
Accepted

## Context
ADR-0079 added direct `break`/`continue` to while-loop bodies, but deferred
propagation through `begin(...)` blocks: `while(x < 100, begin(x = x + 1;
break))` ignored the break and ran to completion.

## Decision
`expandBegin` raises internal sentinel strings (`breakSentinel`,
`continueSentinel`) when it encounters a direct `break`/`continue`. Loop
bodies evaluate statements at the string-expansion level (`substitute`) and
catch the sentinels before parsing: `breakSentinel` exits the loop with the
last value, `continueSentinel` skips the rest of the current iteration.
A new `evalExpanded` helper evaluates the already-expanded string.

## Consequences
break/continue now work inside begin blocks in loop bodies, giving
early-exit/skip control across structured bodies. Whole-program integration
tests cover both propagation paths. One ADR per feature per the agents.md
rule.
