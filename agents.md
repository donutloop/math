# Agent Workflow

This is the root-level prompt for the coding agent building the Math Calculator.
The agent must follow these rules indefinitely — this is a loop, not a one-off.

> **This calculator is built for BOTH humans and agentic workflows.** The
> interface must reflect that duality at every layer: humans get a friendly
> REPL (help, discoverable commands, readable formatting), while agents and
> scripts get structured, predictable, self-describing access (machine-readable
> JSON/CSV, a JSON schema, stable CLI flags, self-describing commands, and
> deterministic exit codes). No feature ships until it has a machine
> consumption path as well as a human one.

## Mission

Build the **best calculator language ever known by humanity**. Think like a
true programming-language expert. Every feature must consider the whole system:

- **Language surface** — statements, expressions, operators, variables,
  functions, control flow, composition.
- **Internal components** — lexer, parser, AST, evaluator, error model,
  formatting. Keep them clean, layered, and extensible.
- **User experience** — a great REPL for humans: clear help, discoverable
  commands, readable formatting, sensible errors, degrees/radians/gradians
  modes.
- **Agentic workflows** — this calculator is not just for humans: it is also
  a computation engine for agents and scripts. The interface must expose
  machine-readable output (JSON/Csv/schema), stable CLI flags (`--eval`,
  `--file`, `--verify`, `--version`), and self-describing commands so an
  agent can discover the full language surface, plan its computation, and
  consume results without guessing.
- **Correctness** — every feature ships with unit tests, domain checks, and
  verify coverage. `go test ./...` must pass before commit.

Because we build for **both humans and agents**, the interface must reflect
both: humans get a friendly REPL; agents get structured, predictable,
self-describing access. Every new feature should consider its machine
consumption path (JSON/Csv output, exit codes, schema) as well as its human
one.

Be creative: prefer language-level features (conditionals, loops, user
functions, assignments, modes, units, formatting) over plain math functions.
Add math functions only when they genuinely expand the language.

Think like somebody writing a brand-new calculator language in 2026: modern
ergonomics, clean error messages, discoverable commands, and a REPL that
feels great to use.

## The loop

1. Pick a new feature.
2. Implement across the stack:
   - lexer/parser/evaluator for syntax and semantics.
   - `calc/help.go` — one-line help.
   - `calc/verify.go` — known-good verify case.
   - docs — `docs/operations.md`, `CHANGELOG.md`, `README.md`, ADR if needed.
   - **Machine path** — if the feature is language/interface, make sure it is
     discoverable via `schema`/help and consumable as JSON/Csv where it makes
     sense (agents must not need to scrape prose).
3. **Add tests for the feature** before committing.
4. **Run tests before committing**: `go test ./...` must pass.
5. Commit with a clear message (`feat(parser): ...`).
6. **Always push your commits**: `git push origin HEAD`.
7. If the remote diverged, `git pull --rebase origin HEAD` then push again.
8. Return to step 1 and repeat — never stop evolving the language.

## New requirements

- Keep `agents.md` at the repo root; do not store prompts under `setup/prompts/`.
- Every feature commit updates `docs/operations.md`, `CHANGELOG.md`,
  `README.md`, and the ADR if relevant.
- `docs/operations.md` is the single source of truth for the language surface.
- Always push every commit; never leave a feature unpushed.
- Preserve the CHANGELOG as a clean single list (deduplicate if it grows).
- Follow this document's direction on every cycle; it is the contract.

## Rules

- One commit per feature; do not bundle unrelated features.
- Never commit failing tests; fix or drop them first.
- Continue the loop even after pushes.

## Iteration contract (never stop)

- Keep executing forever: after each feature is done, immediately begin the
  next one from the mission list (conditionals, loops, user functions,
  assignments, modes, units, formatting, agentic/JSON/schema/schema paths) or
  any remaining gap.
- **Push every completed feature to the remote before starting the next one**:
  `git add -A && git commit -m "<feature>" && git push origin HEAD`.
- If the push fails or the remote diverged, `git pull --rebase origin HEAD`
  and push again.
- There is no terminal state; the loop continues indefinitely.
- Do **not** stop after step 8 or after a single feature — the loop must keep
  running. Every cycle must also keep the agentic interface honest: keep
  `schema`/JSON output current with every new feature added.

## Architecture Decision Records (ADR)

Every feature ships with an ADR in `docs/adr/` (e.g. `0004-json-single-doc.md`,
`0006-exit-code-contract.md`). Each ADR records the decision, the agentic
rationale, and the alternatives rejected — so future cycles and agents can
reconstruct *why* the interface is shaped the way it is. Write the ADR for the
feature *before* committing, and keep it one per feature.
