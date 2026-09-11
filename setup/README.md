# How This Project Was Generated

This project was generated entirely by an AI assistant  using
the following prompt.

The Prompt is located in `code_agent_prompt.md`.

## What the AI Produced

A fully working Go **math programming language** structured by language-design
concern into focused packages:

- `parser`  — tokenizing, parsing, and the AST only (`Parse`, node types, op
  constants). No evaluation lives here.
- `eval`    — AST evaluation and all built-in math functions (`Evaluate`),
  importing only `parser`. Error and AST types are aliased so the evaluator body
  compiles unchanged.
- `calc`    — the language runtime, itself subdivided by concern:
  - `calc/format`  — pure number formatting (`Format`, `FormatPrec`, radix/eng).
  - `calc/schema`  — machine-readable schema + help topics (`BuildSchema`,
    `SchemaJSON`, `PrintSchema`).
  - `calc/verify`  — self-check battery (`Verify`, `VerifyJSON`).
  - `calc/*.go`    — state, assignment, control-flow expansion (loops and
    conditionals), user functions, the REPL, and machine-readable output.
- `tests/`  — unit tests (`tests/unit`) and whole-program integration tests
  (`tests/integration`).

## Key Properties

- **Parsing and evaluation are cleanly separated**: `parser.Parse` builds an AST
  and `eval.Evaluate` interprets it, so the two concerns never share a package.
- **calc is divided into sub-packages** by language-design concern: the runtime
  (`calc`), pure formatting (`calc/format`), machine schema (`calc/schema`),
  and the verification battery (`calc/verify`).
- Typed errors, a strict lexer/parser pipeline, float64 throughout, and a
  deterministic CLI exit-code contract.
- 20+ unit tests and whole-program integration tests, all passing.
- A Makefile, `go.mod`, and a root `README.md` documenting package structure.

## Layout

```
parser/      AST + parsing (no evaluation)
eval/        AST evaluation + math functions
calc/        language runtime
  calc/format/  pure number formatting
  calc/schema/  machine-readable schema + help topics
  calc/verify/  self-check battery
tests/
  unit/        per-package unit tests
  integration/ whole-program integration tests
main.go       CLI entry point
setup/        this document
```
