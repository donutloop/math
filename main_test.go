package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"os"
	"strings"
	"testing"
)

// runMain executes main() with the given args and returns captured stdout.
func runMain(statePath string, args ...string) string {
	flag.CommandLine = flag.NewFlagSet("calculator-test", flag.ExitOnError)
	os.Args = append([]string{"calculator", "-state", statePath}, args...)

	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return ""
	}
	os.Stdout = w
	main()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestMultiEvalSharesState(t *testing.T) {
	got := runMain(t.TempDir()+"/s.json", "-eval", "x=5", "-eval", "x*2")
	if !strings.Contains(got, "10") {
		t.Errorf("repeatable --eval should share variables:\n%s", got)
	}
}

func TestMultiEvalAnsChaining(t *testing.T) {
	got := runMain(t.TempDir()+"/s.json", "-eval", "1+1", "-eval", "ans*10")
	if !strings.Contains(got, "20") {
		t.Errorf("repeatable --eval should chain ans:\n%s", got)
	}
}

func TestStdinEvalMode(t *testing.T) {
	old := os.Stdin
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdin = r
	flag.CommandLine = flag.NewFlagSet("calculator-test2", flag.ExitOnError)
	os.Args = []string{"calculator", "-state", t.TempDir() + "/stdin.json", "--eval", "-"}
	io.WriteString(w, "1+1\nans*10\nquit\n")
	w.Close()
	oldOut := os.Stdout
	pr, pw, _ := os.Pipe()
	os.Stdout = pw
	main()
	pw.Close()
	os.Stdout = oldOut
	os.Stdin = old
	var buf bytes.Buffer
	io.Copy(&buf, pr)
	if !strings.Contains(buf.String(), "20") {
		t.Errorf("stdin eval mode should chain ans:\n%s", buf.String())
	}
}

func TestVarsFlagListsVariables(t *testing.T) {
	got := runMain(t.TempDir()+"/v.json", "-eval", "a=3", "-eval", "b=4", "--vars")
	if !strings.Contains(got, "a = 3") || !strings.Contains(got, "b = 4") {
		t.Errorf("--vars should list variables:\n%s", got)
	}
}

func TestNoStateDisablesPersistence(t *testing.T) {
	got := runMain(t.TempDir()+"/ns.json", "-eval", "x=5", "-eval", "x*2", "--no-state")
	if !strings.Contains(got, "10") {
		t.Errorf("--no-state should still evaluate:\n%s", got)
	}
	// The temp state file should not be created.
	if _, err := os.Stat(t.TempDir() + "/ns.json"); err == nil {
		t.Errorf("state file should not exist with --no-state")
	}
}

func TestJSONEvalOutput(t *testing.T) {
	got := runMain(t.TempDir()+"/j.json", "-eval", "1+1", "-eval", "gcd(12,18)", "--json", "--no-state")
	if !strings.Contains(got, `"value": 2`) || !strings.Contains(got, `"value": 6`) {
		t.Errorf("JSON eval output missing values:\n%s", got)
	}
}

func TestJSONEvalAssignment(t *testing.T) {
	got := runMain(t.TempDir()+"/ja.json", "-eval", "x=5", "-eval", "x*2", "--json", "--no-state")
	if !strings.Contains(got, `"assign": "x"`) || !strings.Contains(got, `"value": 10`) {
		t.Errorf("JSON assignment output missing:\n%s", got)
	}
}

func TestJSONVars(t *testing.T) {
	got := runMain(t.TempDir()+"/jv.json", "-eval", "a=3", "-eval", "b=4", "--json", "--vars", "--no-state")
	if !strings.Contains(got, `"var": "a"`) || !strings.Contains(got, `"var": "b"`) {
		t.Errorf("JSON vars output missing:\n%s", got)
	}
}

func TestCSVEvalOutput(t *testing.T) {
	got := runMain(t.TempDir()+"/c.json", "-eval", "x=5", "-eval", "x*2", "--csv", "--no-state")
	if !strings.Contains(got, "assign,x,5") || !strings.Contains(got, "expr,x*2,10") {
		t.Errorf("CSV eval output missing:\n%s", got)
	}
}

func TestJSONLEvalOutput(t *testing.T) {
	got := runMain(t.TempDir()+"/jl.json", "-eval", "1+1", "-eval", "x=5", "--jsonl", "--no-state")
	lines := strings.Split(strings.TrimSpace(got), "\n")
	if len(lines) != 2 {
		t.Fatalf("JSONL should emit one object per line, got %d lines:\n%s", len(lines), got)
	}
	// every line must be a single parseable JSON object with the stable shape.
	for i, ln := range lines {
		var m map[string]any
		if err := json.Unmarshal([]byte(ln), &m); err != nil {
			t.Errorf("line %d not parseable JSON: %v", i+1, err)
		}
		if _, ok := m["kind"]; !ok {
			t.Errorf("line %d missing kind:\n%s", i+1, ln)
		}
	}
	if !strings.Contains(got, `"kind":"value"`) || !strings.Contains(got, `"kind":"assign"`) {
		t.Errorf("JSONL kinds missing:\n%s", got)
	}
}

func TestVerifyJSON(t *testing.T) {
	got := runMain(t.TempDir()+"/v.json", "--verify", "--json")
	var d map[string]any
	if err := json.Unmarshal([]byte(got), &d); err != nil {
		t.Fatalf("--verify --json not parseable JSON: %v\n%s", err, got)
	}
	if d["failed"].(float64) != 0 || d["passed"].(float64) <= 0 {
		t.Errorf("verify report should have 0 failures, got passed=%v failed=%v", d["passed"], d["failed"])
	}
	checks, ok := d["checks"].([]any)
	if !ok || len(checks) == 0 {
		t.Errorf("verify report missing checks array")
	}
	// every check must be a structured {check,pass} object.
	for i, c := range checks {
		m, ok := c.(map[string]any)
		if !ok || m["pass"] == nil || m["check"] == nil {
			t.Errorf("check %d not structured {check,pass}: %v", i, c)
		}
	}
}

func TestOutputSelector(t *testing.T) {
	// --output json must equal --json for the same eval.
	want := runMain(t.TempDir()+"/o.json", "--eval", "1+1", "--json")
	got := runMain(t.TempDir()+"/o2.json", "--eval", "1+1", "--output", "json")
	if got != want {
		t.Errorf("--output json != --json:\n got %q\nwant %q", got, want)
	}
	// --output jsonl must equal --jsonl.
	wantL := runMain(t.TempDir()+"/ol.json", "--eval", "1+1", "--jsonl")
	gotL := runMain(t.TempDir()+"/ol2.json", "--eval", "1+1", "--output", "jsonl")
	if gotL != wantL {
		t.Errorf("--output jsonl != --jsonl:\n got %q\nwant %q", gotL, wantL)
	}
	// --output csv must equal --csv.
	wantC := runMain(t.TempDir()+"/oc.json", "--eval", "1+1", "--csv")
	gotC := runMain(t.TempDir()+"/oc2.json", "--eval", "1+1", "--output", "csv")
	if gotC != wantC {
		t.Errorf("--output csv != --csv:\n got %q\nwant %q", gotC, wantC)
	}
}

func TestSchemaVersion(t *testing.T) {
	got := runMain(t.TempDir()+"/v.json", "--schema")
	var d map[string]any
	if err := json.Unmarshal([]byte(got), &d); err != nil {
		t.Fatalf("--schema not parseable JSON: %v\n%s", err, got)
	}
	if d["version"] != "1.1.0" {
		t.Errorf("schema version should be 1.1.0, got %v", d["version"])
	}
	// --version features must advertise the full machine surface.
	gotV := runMain(t.TempDir()+"/v2.json", "--version")
	var d2 map[string]any
	if err := json.Unmarshal([]byte(gotV), &d2); err != nil {
		t.Fatalf("--version not parseable JSON: %v\n%s", err, gotV)
	}
	feat, ok := d2["features"].([]any)
	if !ok {
		t.Fatalf("--version missing features array: %v", gotV)
	}
	for _, want := range []string{"jsonl", "output", "verify-json"} {
		found := false
		for _, f := range feat {
			if f == want {
				found = true
			}
		}
		if !found {
			t.Errorf("--version features missing %q", want)
		}
	}
}
