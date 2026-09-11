package schema

import (
	"encoding/json"
	"strings"
	"testing"
)

// TestSchemaJSON verifies the schema renders as valid, self-describing JSON
// that an agent can parse without scraping prose help.
func TestSchemaJSON(t *testing.T) {
	b, err := SchemaJSON()
	if err != nil {
		t.Fatalf("SchemaJSON: %v", err)
	}
	var s Schema
	if err := json.Unmarshal(b, &s); err != nil {
		t.Fatalf("schema is not parseable JSON: %v", err)
	}
	if s.Name != "math-calculator" {
		t.Errorf("name = %q, want math-calculator", s.Name)
	}
	if s.Version == "" {
		t.Errorf("version empty")
	}
	if len(s.Functions) == 0 {
		t.Errorf("no functions in schema")
	}
	if len(s.Constants) == 0 {
		t.Errorf("no constants in schema")
	}
	if len(s.Commands) == 0 {
		t.Errorf("no commands in schema")
	}
	if len(s.Operators) == 0 {
		t.Errorf("no operators in schema")
	}
	if len(s.Modes) == 0 {
		t.Errorf("no modes in schema")
	}
	// Every built-in function must carry arity metadata.
	for name, fn := range s.Functions {
		if fn.Arity == 0 && name != "sum" && name != "count" && name != "min" && name != "max" && name != "avg" && name != "gcd" && name != "lcm" && name != "round" && name != "multinomial" {
			t.Errorf("function %q has zero arity", name)
		}
		if fn.Help == "" {
			t.Errorf("function %q has no help", name)
		}
	}
	// Constants must carry numeric values.
	if c, ok := s.Constants["pi"]; !ok || c.Value == 0 {
		t.Errorf("pi constant missing or zero")
	}
	if c, ok := s.Constants["e"]; !ok || c.Value == 0 {
		t.Errorf("e constant missing or zero")
	}
	// The schema must advertise the machine-facing surface.
	if !contains(s.Commands, "schema") {
		t.Errorf("schema command not advertised")
	}
	if !contains(s.CLI, "--schema") {
		t.Errorf("--schema CLI flag not advertised")
	}
}

// TestPrintSchema exercises the REPL printer.
func TestPrintSchema(t *testing.T) {
	var out strings.Builder
	if err := PrintSchema(&out); err != nil {
		t.Fatalf("PrintSchema: %v", err)
	}
	if !strings.Contains(out.String(), "\"name\": \"math-calculator\"") {
		t.Errorf("schema output missing name: %s", out.String())
	}
	if !strings.Contains(out.String(), "\"functions\"") {
		t.Errorf("schema output missing functions section")
	}
}

func contains(items []string, want string) bool {
	for _, it := range items {
		if it == want {
			return true
		}
	}
	return false
}

// TestSchemaExitCodes verifies the deterministic exit-code contract is exposed
// to agents in the schema.
func TestSchemaExitCodes(t *testing.T) {
	b, err := SchemaJSON()
	if err != nil {
		t.Fatalf("SchemaJSON: %v", err)
	}
	var s Schema
	if err := json.Unmarshal(b, &s); err != nil {
		t.Fatalf("schema not parseable: %v", err)
	}
	if s.ExitCodes["ok"] != 0 || s.ExitCodes["usage"] != 1 || s.ExitCodes["io"] != 2 || s.ExitCodes["eval"] != 3 {
		t.Errorf("exit-code contract wrong: %v", s.ExitCodes)
	}
}
