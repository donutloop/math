package schema

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"

	"prototype_kl/parser"
)

// Schema is the machine-readable, self-describing description of the full
// calculator language surface. An agent or script can consume it (as JSON) to
// discover every function, constant, command, operator, display mode, and CLI
// flag — without scraping prose help.
type Schema struct {
	Version   string                    `json:"version"`
	Name      string                    `json:"name"`
	Functions map[string]FuncSchema     `json:"functions"`
	Constants map[string]ConstantSchema `json:"constants"`
	Commands  []string                  `json:"commands"`
	Operators []string                  `json:"operators"`
	Modes     []string                  `json:"modes"`
	CLI       []string                  `json:"cli_flags"`
	Formats   []string                  `json:"formats"`
	ExitCodes map[string]int            `json:"exit_codes"`
}

// FuncSchema describes a built-in function: its arity and one-line help.
type FuncSchema struct {
	Arity int    `json:"arity"`
	Help  string `json:"help"`
}

// ConstantSchema describes a built-in constant and its numeric value.
type ConstantSchema struct {
	Value float64 `json:"value"`
	Help  string  `json:"help"`
}

// schemaVersion is the calculator release version exposed to agents.
const schemaVersion = "1.1.0"

// SchemaVersion returns the machine-surface version so agents can pin
// compatibility and detect breaking changes across releases.
func SchemaVersion() string { return schemaVersion }

// schemaCommands is the full interactive REPL command surface.
var schemaCommands = []string{
	"help", "?", "quit", "exit", "q", "vars", "clear", "reset", "history",
	"ans", "ms", "m+", "m-", "mr", "mc", "mem", "deg", "rad", "grad", "sci",
	"fix", "eng", "std", "prec <n>", "base hex|dec|oct|bin", "quiet", "json",
	"csv", "status", "last", "undo", "redo", "func f(x)=...", "schema",
	"tree <expr>", "@N",
}

// schemaOperators is the operator/notation surface.
var schemaOperators = []string{
	"+", "-", "*", "/", "^", "(", ")", "!", "%", "%%", "&", "|", "~", "<<", ">>",
	"<", ">", "<=", ">=", "==", "!=", "&&", "||", "?", ":", ",", "=",
}

// schemaModes is the display/angle-mode surface.
var schemaModes = []string{
	"degrees", "radians", "gradians", "scientific", "engineering",
	"fixed", "precision", "base", "quiet", "json", "csv",
}

// schemaExitCodes documents the deterministic scripting exit-code contract.
var schemaExitCodes = map[string]int{
	"ok":    0,
	"usage": 1,
	"io":    2,
	"eval":  3,
}

// schemaCLI is the stable non-interactive CLI surface for agents and scripts.
var schemaCLI = []string{
	"--eval <expr>", "--file <path>", "--verify", "--version", "--help",
	"--json", "--jsonl", "--csv", "--output <format>", "--vars", "--snapshot", "--base <radix>", "--prec <n>", "--sci",
	"--eng", "--deg", "--rad", "--grad", "--demo", "--state <path>",
	"--no-state", "--quiet", "--schema",
}

// BuildSchema returns the full language surface as a Schema value.
func BuildSchema() Schema {
	fns := make(map[string]FuncSchema, len(parser.SupportedFunctions))
	for name, arity := range parser.SupportedFunctions {
		help := Topics[strings.ToLower(name)]
		if help == "" {
			help = fmt.Sprintf("%s(...): built-in function", name)
		}
		fns[name] = FuncSchema{Arity: arity, Help: help}
	}
	consts := make(map[string]ConstantSchema, len(parser.SupportedConstants))
	for name, v := range parser.SupportedConstants {
		help := Topics[strings.ToLower(name)]
		if help == "" {
			help = fmt.Sprintf("%s: constant", name)
		}
		consts[name] = ConstantSchema{Value: v, Help: help}
	}
	return Schema{
		Version:   schemaVersion,
		Name:      "math-calculator",
		Functions: fns,
		Constants: consts,
		Commands:  schemaCommands,
		Operators: schemaOperators,
		Modes:     schemaModes,
		Formats:   []string{"text", "json", "jsonl", "csv"},
		CLI:       schemaCLI,
		ExitCodes: schemaExitCodes,
	}
}

// SchemaJSON renders the schema as indented, machine-readable JSON.
func SchemaJSON() ([]byte, error) {
	s := BuildSchema()
	return json.MarshalIndent(s, "", "  ")
}

// PrintSchema writes the schema JSON to w.
func PrintSchema(w io.Writer) error {
	out, err := SchemaJSON()
	if err != nil {
		return err
	}
	fmt.Fprintln(w, string(out))
	return nil
}

// SchemaNames returns the sorted list of function names for tests/docs.
func SchemaNames() []string {
	names := make([]string, 0, len(parser.SupportedFunctions))
	for name := range parser.SupportedFunctions {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// SchemaCount is the number of built-in functions in the schema.
func SchemaCount() int { return len(parser.SupportedFunctions) }

// SchemaToWriter writes the schema JSON to any io.Writer (used by main).
func SchemaToWriter(w io.Writer) error {
	out, err := SchemaJSON()
	if err != nil {
		return err
	}
	_, err = w.Write(out)
	return err
}
