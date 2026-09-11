// Command calculator is a full-featured interactive math calculator.
//
// Usage:
//
//	calculator                 start the interactive REPL
//	calculator --eval "expr"   evaluate one expression and print the result
//	calculator --help          show usage
//
// Expressions support arithmetic (+ - * /), parentheses, the constants pi/e,
// and a rich set of functions (trig, hyperbolic, logs, pow, min/max, fact).
// In the REPL you can also assign variables and recall the last result as
// "ans".
package main

import (
	"errors"
	"bytes"
	"flag"
	jsonenc "encoding/json"
	"fmt"
	"os"
	"strings"

	"prototype_kl/calc"
	"reflect"
)

// Version is the calculator release version.
const Version = "1.0.0"


// Exit-code contract for scripting agents and CI. Deterministic and documented
// in README/docs: 0=success, 1=usage/flag error, 2=IO error, 3=eval error.
const (
	ExitOK    = 0
	ExitUsage = 1
	ExitIO    = 2
	ExitEval  = 3
)
func main() {

	// Custom FlagSet with ContinueOnError so usage errors exit with the
	// documented contract code (1) instead of Go's default os.Exit(2).
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: calculator [flags]\n")
	}
	var evals []string
	output := flag.String("output", "", "output format: text|json|jsonl|csv (unified selector)")
	csv := flag.Bool("csv", false, "emit CSV results for --eval")
	quiet := flag.Bool("quiet", false, "suppress assignment echoes")
	json := flag.Bool("json", false, "emit JSON results for --eval")
	jsonl := flag.Bool("jsonl", false, "emit NDJSON: one JSON object per result line")
	calcHelp := flag.Bool("help", false, "print usage and exit 0")
	vars := flag.Bool("vars", false, "list defined variables after evaluation")
	snapshot := flag.Bool("snapshot", false, "dump REPL runtime state as JSON")
	base := flag.Int("base", 0, "output radix for integral results (2, 8, 16, or 0=decimal)")
	flag.Var(&multiFlag{&evals}, "eval", "evaluate an expression and print the result; may be given multiple times")
	state := flag.String("state", ".calc-state.json", "persist variables/history across sessions")
	noState := flag.Bool("no-state", false, "disable state persistence")
	prec := flag.Int("prec", 15, "significant digits for --eval output (1..17)")
	sci := flag.Bool("sci", false, "scientific notation for --eval output")
	demo := flag.Bool("demo", false, "run a guided tour")
	file := flag.String("file", "", "evaluate expressions from a file (batch)")
	deg := flag.Bool("deg", false, "trig in degrees for --eval")
	version := flag.Bool("version", false, "print version and exit")
	eng := flag.Bool("eng", false, "engineering notation for output")
	verify := flag.Bool("verify", false, "run the self-test battery")
	schemaFlag := flag.Bool("schema", false, "print the machine-readable JSON language schema and exit")
	rad := flag.Bool("rad", false, "trig in radians (default)")
	grad := flag.Bool("grad", false, "trig in gradians")
	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		if errors.Is(err, flag.ErrHelp) { os.Exit(ExitOK) }
		os.Exit(ExitUsage)
	}
	// --output is a unified format selector for agents: map it to the
	// individual json/jsonl/csv flags so one flag replaces flag-discovery.
	switch *output {
	case "", "text":
		// prose (default)
	case "json":
		*json = true
	case "jsonl":
		*jsonl = true
	case "csv":
		*csv = true
	default:
		fmt.Fprintf(os.Stderr, "error: unknown --output %q (want text|json|jsonl|csv)\n", *output)
		os.Exit(ExitUsage)
	}
	if *prec < 1 || *prec > 17 {
		fmt.Fprintln(os.Stderr, "error: --prec must be 1..17")
		os.Exit(ExitUsage)
	}
	if *noState {
		*state = ""
	}

	if *calcHelp {
		if *json {
			flags := []map[string]any{}
			flag.CommandLine.VisitAll(func(f *flag.Flag) {
				flags = append(flags, map[string]any{"name": f.Name, "type": reflect.TypeOf(f.Value).String(), "help": f.Usage})
			})
			b, err := jsonenc.MarshalIndent(map[string]any{
				"version":   calc.SchemaVersion(),
				"flags":     flags,
				"exit_codes": map[string]int{"ok": ExitOK, "usage": ExitUsage, "io": ExitIO, "eval": ExitEval},
			}, "", "  ")
			if err != nil {
				fmt.Fprintln(os.Stderr, "help: json:", err)
				os.Exit(ExitIO)
			}
			fmt.Println(string(b))
			return
		}
		flag.CommandLine.SetOutput(os.Stdout)
		flag.CommandLine.Usage()
		fmt.Println("Machine outputs: --schema --json --jsonl --csv --version")
		fmt.Println("Exit codes: 0=ok 1=usage 2=io 3=eval")
		return
	}
	if *version {
		// Machine-readable version for agent compatibility checks.
		b, err := jsonenc.MarshalIndent(map[string]any{
			"name":    "math-calculator",
			"version": calc.SchemaVersion(),
			"schema":  "1.0.0",
			"features": []string{"schema", "json", "jsonl", "csv", "output", "verify", "verify-json", "eval", "vars", "file", "help"},
			"exit_codes": map[string]int{
				"ok": 0, "usage": 1, "io": 2, "eval": 3,
			},
		}, "", "  ")
		if err != nil {
			os.Exit(ExitIO)
		}
		fmt.Println(string(b))
		return
	}


	if *schemaFlag {
		if err := calc.SchemaToWriter(os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, "error: schema:", err)
			os.Exit(ExitIO)
		}
		fmt.Println()
		return
	}

	if *verify {
		// --verify --json emits a machine-readable health report so agents can
		// parse passed/failed counts and each check without scraping prose.
		if *json {
			b, err := jsonenc.MarshalIndent(calc.VerifyJSON(), "", "  ")
			if err != nil {
				fmt.Fprintln(os.Stderr, "verify: json:", err)
				os.Exit(ExitIO)
			}
			fmt.Println(string(b))
			if calc.VerifyJSON()["failed"].(int) > 0 {
				os.Exit(ExitEval)
			}
			return
		}
		passed, failed := calc.Verify(os.Stdout)
		if failed > 0 {
			os.Exit(ExitEval)
		}
		fmt.Fprintf(os.Stderr, "verify: %d ok\n", passed)
		return
	}

	if *demo {
		calc.Demo(os.Stdout)
		return
	}

	if *file != "" {
		f, err := os.Open(*file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(ExitIO)
		}
		defer f.Close()
		c := calc.NewBatch(f, os.Stdout)
		if *state != "" {
			_ = c.LoadState(*state)
		}
		c.SetDisplay(*prec, *sci, *deg)
		if *rad {
			c.SetRad()
		} else if *grad {
			c.SetGrad()
		}
		if *eng {
			c.Eng()
		}
		if *base != 0 {
			if err := c.SetBase(*base); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(ExitUsage)
			}
		}
		c.SetQuietAssign(*quiet)
		c.Run()
		if *state != "" {
			_ = c.SaveState(*state)
		}
		return
	}

	if len(evals) == 1 && evals[0] == "-" {
		var out bytes.Buffer
		c := calc.NewBatch(os.Stdin, &out)
		c.SetDisplay(*prec, *sci, *deg)
		if *eng {
			c.Eng()
		}
		if *rad {
			c.SetRad()
		} else if *grad {
			c.SetGrad()
		}
		if *state != "" {
			_ = c.LoadState(*state)
		}
		c.SetQuietAssign(*quiet)
		c.Run()
		fmt.Print(out.String())
		if *state != "" {
			_ = c.SaveState(*state)
		}
		return
	}

	if len(evals) > 0 {
		var out bytes.Buffer
		c := calc.NewBatch(strings.NewReader(strings.Join(evals, ";")), &out)
		if *state != "" {
			_ = c.LoadState(*state)
		}
		c.SetDisplay(*prec, *sci, *deg)
		if *rad {
			c.SetRad()
		} else if *grad {
			c.SetGrad()
		}
		if *eng {
			c.Eng()
		}
		if *base != 0 {
			if err := c.SetBase(*base); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(ExitEval)
			}
		}
		c.SetQuietAssign(*quiet)
	if *csv {
		// Emit a single, valid CSV document with a header row so agents can parse
		// the whole stdout at once (csv.reader / pandas.read_csv).
		rows := []string{"kind,key,value"} // header row
		for _, e := range evals {
			if strings.Contains(e, "=") {
				name, v, err := c.AssignExpr(e)
				if err != nil {
					rows = append(rows, fmt.Sprintf("error,%s,%s", e, err.Error()))
					continue
				}
				rows = append(rows, fmt.Sprintf("assign,%s,%s", name, c.FormatValue(v)))
				continue
			}
			v, err := c.EvalExpr(e)
			if err != nil {
				rows = append(rows, fmt.Sprintf("error,%s,%s", e, err.Error()))
				continue
			}
			rows = append(rows, fmt.Sprintf("expr,%s,%s", e, c.FormatValue(v)))
		}
		if *vars {
			for name, v := range c.Vars() {
				rows = append(rows, fmt.Sprintf("var,%s,%s", name, c.FormatValue(v)))
			}
		}
		if *state != "" {
			_ = c.SaveState(*state)
		}
		for _, r := range rows {
			fmt.Println(r)
		}
		return
	}

	if *jsonl {
		for _, e := range evals {
			if strings.Contains(e, "=") {
				name, v, err := c.AssignExpr(e)
				if err != nil {
					printJSONL(e, "error", err.Error())
					continue
				}
				printJSONL(e, "assign", name+"="+c.FormatValue(v))
				continue
			}
			v, err := c.EvalExpr(e)
			if err != nil {
				printJSONL(e, "error", err.Error())
				continue
			}
			printJSONL(e, "value", v)
		}
		if *vars {
			for name, v := range c.Vars() {
				printJSONL("", "var", name+"="+c.FormatValue(v))
			}
		}
		if *state != "" {
			_ = c.SaveState(*state)
		}
		return
	}
		if *json {
	if *json {
		// Emit a single, parseable JSON document (array) so agents can consume
		// it without scraping line-delimited fragments.
		results := make([]map[string]any, 0, len(evals)+4)
		for _, e := range evals {
			if strings.Contains(e, "=") {
				name, v, err := c.AssignExpr(e)
				if err != nil {
					results = append(results, map[string]any{"expr": e, "error": err.Error()})
					continue
				}
				results = append(results, map[string]any{"expr": e, "assign": name, "value": v})
				continue
			}
			v, err := c.EvalExpr(e)
			if err != nil {
				results = append(results, map[string]any{"expr": e, "error": err.Error()})
				continue
			}
			results = append(results, map[string]any{"expr": e, "value": v})
		}
		if *vars {
			for name, v := range c.Vars() {
				results = append(results, map[string]any{"var": name, "value": v})
			}
		}
		if *state != "" {
			_ = c.SaveState(*state)
		}
		b, err := jsonenc.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Fprintln(os.Stderr, "error: json:", err)
			os.Exit(ExitIO)
		}
		fmt.Println(string(b))
		return
	}

		}
		c.Run()
		if strings.TrimSpace(out.String()) == "" {
			fmt.Fprintln(os.Stderr, "error: empty evaluation")
			os.Exit(ExitEval)
		}
		fmt.Print(out.String())
		if *vars {
			c.PrintVars()
		}
		if *snapshot {
			b, err := jsonenc.MarshalIndent(c.Snapshot(), "", "  ")
			if err != nil {
				fmt.Fprintln(os.Stderr, "snapshot: json:", err)
				os.Exit(ExitIO)
			}
			fmt.Println(string(b))
			return
		}
		if *state != "" {
			_ = c.SaveState(*state)
		}
		if c.ErrorCount() > 0 {
			os.Exit(ExitEval)
		}
		return
	}

	c, err := calc.NewPersistent(os.Stdin, os.Stdout, *state)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: could not load state %s: %v\n", *state, err)
		c = calc.New(os.Stdin, os.Stdout)
	}
	if *base != 0 {
		if err := c.SetBase(*base); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(ExitUsage)
		}
	}
	c.SetQuietAssign(*quiet)
	if *snapshot {
		b, err := jsonenc.MarshalIndent(c.Snapshot(), "", "  ")
		if err != nil {
			fmt.Fprintln(os.Stderr, "snapshot: json:", err)
			os.Exit(ExitIO)
		}
		fmt.Println(string(b))
		return
	}
	// Spawn a machine REPL session directly when a format flag is given
	// without --eval: presets the mode and suppresses the prose banner so a
	// driving agent gets clean parseable output from the first line.
	if len(evals) == 0 {
		switch {
		case *jsonl:
			if err := c.SetMachineMode("jsonl"); err != nil {
				fmt.Fprintln(os.Stderr, "repl:", err)
				os.Exit(ExitUsage)
			}
		case *json:
			if err := c.SetMachineMode("json"); err != nil {
				fmt.Fprintln(os.Stderr, "repl:", err)
				os.Exit(ExitUsage)
			}
		case *csv:
			if err := c.SetMachineMode("csv"); err != nil {
				fmt.Fprintln(os.Stderr, "repl:", err)
				os.Exit(ExitUsage)
			}
		case *output != "" && *output != "text":
			if err := c.SetMachineMode(*output); err != nil {
				fmt.Fprintln(os.Stderr, "repl:", err)
				os.Exit(ExitUsage)
			}
		}
	}
	c.Run()
}

// multiFlag accumulates a repeatable string flag.
type multiFlag struct{ list *[]string }

func (m multiFlag) String() string {
	return strings.Join(*m.list, ",")
}

func (m multiFlag) Set(v string) error {
	*m.list = append(*m.list, v)
	return nil
}

// printJSONL writes one NDJSON object per line for streaming agents.
func printJSONL(expr, kind, value any) {
	b, err := jsonenc.Marshal(map[string]any{
		"expr":  expr,
		"kind":  kind,
		"value": value,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: jsonl: %v\n", err)
		os.Exit(ExitIO)
	}
	fmt.Println(string(b))
}


// printJSONL writes one NDJSON object per line for streaming agents.
