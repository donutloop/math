# ADR 0008: let() local bindings

**Status:** Accepted
**Date:** 2026-09-10

## Context

The calculator already has global variable assignment (`x = 5`) and user
functions. Expressions that need intermediate values either pollute global
scope or force defining one-off functions. We want local composition without
global state, matching the mission's "assignments" language-level feature.

## Decision

Add a macro `let(x = e1, y = e2, body)` handled in the calculator's `expand`
pass (like `convert`, `countif`, range loops), not the parser. The last
argument is the body; every earlier argument is a binding `name = expr`.
Bindings may reference earlier bindings. Each binding value is parenthesized
on substitution to preserve precedence, so `let(a = 5, b = 10, a + b)` is
`(5) + (10)` = 15.

To avoid `let(x = 2, ...)` being misparsed as a global assignment, `findAssignEq`
now ignores `=` at paren depth > 0. `let` is added to the reserved-name list.

## Consequences

- Expressions get local scope and compose cleanly without global pollution.
- `let` is a macro, so it is verified by unit tests (calc/let_test.go) rather
  than the parser-level verify harness.
- Assignment detection is now paren-aware, which also future-proofs other
  paren-bearing macros.
