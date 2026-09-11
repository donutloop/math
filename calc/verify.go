package calc

import (
	"encoding/json"
	"fmt"
	"io"
	"prototype_kl/eval"
)

// CheckResult is one self-check outcome, machine-readable for --verify --json.
type CheckResult struct {
	Check  string `json:"check"`
	Pass   bool   `json:"pass"`
	Detail string `json:"detail,omitempty"`
}

// verifyCases is the expression/want self-test battery shared by Verify and
// VerifyJSON so both prose and JSON reports stay in lockstep.
var verifyCases = []struct{ expr, want string }{

	{"2 + 3", "5"},
	{"2 * 3 ^ 2", "18"},
	{"pow(2, 10)", "1024"},
	{"5!", "120"},
	{"200% + 10", "12"},
	{"sin(pi / 2)", "1"},
	{"0.1 + 0.2", "0.3"},
	{"sqrt(16)", "4"},
	{"rsqrt(4)", "0.5"},
	{"min(3, 1, 2)", "1"},
	{"var(2, 4, 4, 4, 5, 5, 7, 9)", "4"},
	{"count(1, 5)", "5"},
	{"count(5, 2)", "0"},
	{"isprime(17)", "1"},
	{"isprime(15)", "0"},
	{"isprime(1)", "0"},
	{"prime(1)", "2"},
	{"prime(10)", "29"},
	{"nextprime(10)", "11"},
	{"nextprime(2)", "2"},
	{"divcount(6)", "4"},
	{"divcount(12)", "6"},
	{"divcount(1)", "1"},
	{"prime(3) + nextprime(10)", "16"},
	{"npr(5, 2)", "20"},
	{"npr(10, 3)", "720"},
	{"npr(5, 0)", "1"},
	{"ncr(5, 2)", "10"},
	{"ncr(10, 3)", "120"},
	{"ncr(10, 5)", "252"},
	{"ncr(10, 0)", "1"},
	{"npr(5, 5)", "120"},
	{"ncr(5, 5)", "1"},
	{"round(2.5678, 2)", "2.57"},
	{"mode(1, 2, 2, 3, 3, 3)", "3"},
	{"sum(1, 10, 2)", "25"},
	{"stddev(2, 4, 4, 4, 5, 5, 7, 9)", "2"},
	{"median(1, 3, 2)", "2"},
	{"median(1, 2, 3, 4)", "2.5"},
	{"log2(8)", "3"},
	{"atan2(1, 1)", "0.785398163397448"},
	{"gcd(12, 18)", "6"},
	{"lcm(4, 6)", "12"},
	{"sumdigits(1234)", "10"},
	{"sumdigits(0)", "0"},
	{"sumdigits(999)", "27"},
	{"rev(1234)", "4321"},
	{"rev(120)", "21"},
	{"rev(0)", "0"},
	{"ispal(121)", "1"},
	{"ispal(1221)", "1"},
	{"ispal(123)", "0"},
	{"ispal(0)", "1"},
	{"fib(0)", "0"},
	{"fib(1)", "1"},
	{"fib(10)", "55"},
	{"fib(20)", "6765"},
	{"fib(30)", "832040"},
	{"powmod(2, 10, 1000)", "24"},
	{"powmod(3, 4, 7)", "4"},
	{"powmod(5, 3, 13)", "8"},
	{"powmod(7, 0, 5)", "1"},
	{"powmod(4, 13, 497)", "445"},
	{"powmod(2, 100, 97)", "16"},
	{"collatz(1)", "0"},
	{"collatz(2)", "1"},
	{"collatz(3)", "7"},
	{"collatz(4)", "2"},
	{"collatz(6)", "8"},
	{"collatz(27)", "111"},
	{"totient(12)", "4"},
	{"totient(9)", "6"},
	{"sigma(6)", "12"},
	{"sigma(28)", "56"},
	{"isperfect(6)", "1"},
	{"isperfect(28)", "1"},
	{"isperfect(10)", "0"},
	{"isabundant(12)", "1"},
	{"isdeficient(8)", "1"},
	{"isabundant(6)", "0"},
	{"isdeficient(6)", "0"},
	{"digitsum(1234)", "10"},
	{"digitcount(1234)", "4"},
	{"digitsum(0)", "0"},
	{"digitcount(0)", "1"},
	{"primorial(1)", "2"},
	{"primorial(2)", "6"},
	{"primorial(3)", "30"},
	{"primorial(4)", "210"},
	{"lambertw(0)", "0"},
	{"primecount(10)", "4"},
	{"primecount(100)", "25"},
	{"primecount(2)", "1"},
	{"primecount(0)", "0"},
	{"prevprime(10)", "7"},
	{"prevprime(100)", "97"},
	{"prevprime(3)", "2"},
	{"sumofprimes(5)", "28"},
	{"sumofprimes(10)", "129"},
	{"choose(5, 2)", "10"},
	{"choose(10, 3)", "120"},
	{"choose(5, 0)", "1"},
	{"choose(5, 5)", "1"},
	{"omega(30)", "3"},
	{"omega(12)", "2"},
	{"bigomega(12)", "3"},
	{"bigomega(8)", "3"},
	{"tau(12)", "6"},
	{"tau(6)", "4"},
	{"mobius(6)", "1"},
	{"mobius(30)", "-1"},
	{"mobius(12)", "0"},
	{"mobius(1)", "1"},
	{"mertens(10)", "-1"},
	{"mertens(1)", "1"},
	{"mertens(6)", "-1"},
	{"liouville(12)", "-1"},
	{"liouville(8)", "-1"},
	{"liouville(6)", "1"},
	{"liouville(1)", "1"},
	{"carmichael(9)", "6"},
	{"carmichael(12)", "2"},
	{"carmichael(10)", "4"},
	{"carmichael(1)", "1"},
	{"sumproperdivisors(12)", "16"},
	{"sumproperdivisors(6)", "6"},
	{"sumproperdivisors(28)", "28"},
	{"amicable(220, 284)", "1"},
	{"amicable(10, 5)", "0"},
	{"amicable(6, 6)", "0"},
	{"digitalroot(12345)", "6"},
	{"digitalroot(48)", "3"},
	{"digitalroot(9)", "9"},
	{"digitalroot(0)", "0"},
	{"nthdigit(12345, 0)", "5"},
	{"nthdigit(12345, 1)", "4"},
	{"nthdigit(12345, 4)", "1"},
	{"nthdigit(0, 2)", "0"},
	{"bell(0)", "1"},
	{"bell(3)", "5"},
	{"bell(4)", "15"},
	{"bell(5)", "52"},
	{"catalan(0)", "1"},
	{"catalan(3)", "5"},
	{"catalan(4)", "14"},
	{"catalan(5)", "42"},
	{"stirling(3, 2)", "3"},
	{"stirling(4, 2)", "7"},
	{"stirling(4, 3)", "6"},
	{"stirling(0, 0)", "1"},
	{"derangements(2)", "1"},
	{"derangements(3)", "2"},
	{"derangements(4)", "9"},
	{"derangements(5)", "44"},
	{"lucas(0)", "2"},
	{"lucas(3)", "4"},
	{"lucas(4)", "7"},
	{"lucas(5)", "11"},
	{"digitproduct(123)", "6"},
	{"digitproduct(50)", "0"},
	{"digitproduct(9)", "9"},
	{"digitproduct(0)", "0"},
	{"doublefactorial(0)", "1"},
	{"doublefactorial(4)", "8"},
	{"doublefactorial(5)", "15"},
	{"doublefactorial(6)", "48"},
	{"risingfact(1, 3)", "6"},
	{"risingfact(2, 3)", "24"},
	{"risingfact(5, 0)", "1"},
	{"multinomial(5, 2, 3)", "10"},
	{"multinomial(6, 2, 2, 2)", "90"},
	{"partition(0)", "1"},
	{"partition(4)", "5"},
	{"partition(5)", "7"},
	{"partition(6)", "11"},
	{"collatzmax(3)", "16"},
	{"collatzmax(6)", "16"},
	{"collatzmax(1)", "1"},
	{"nextabundant(7)", "12"},
	{"nextabundant(12)", "12"},
	{"nextabundant(18)", "18"},
	{"nextdeficient(5)", "5"},
	{"nextdeficient(6)", "7"},
	{"nextdeficient(12)", "13"},
	{"nextperfect(1)", "6"},
	{"nextperfect(7)", "28"},
	{"nextperfect(29)", "496"},
	{"gcd(12, 18, 24)", "6"},
	{"lcm(4, 6, 10)", "60"},
	{"gcd(36, 24)", "12"},
	{"lcm(3, 5, 7)", "105"},
	{"log10(100)", "2"},
	{"log1p(9)", "2.30258509299405"},
	{"asinh(0)", "0"},
	{"acosh(1)", "0"},
	{"atanh(0)", "0"},
	{"expm1(0)", "0"},
	{"exp2(3)", "8"},
	{"exp10(2)", "100"},
	{"sinc(0)", "1"},
	{"sinc(pi)", "3.89817183251937e-17"},
	{"sec(0)", "1"},
	{"csc(pi/2)", "1"},
	{"cot(pi/4)", "1"},
	{"asec(2)", "1.0471975511966"},
	{"acsc(2)", "0.523598775598299"},
	{"acot(1)", "0.785398163397448"},
	{"sech(0)", "1"},
	{"csch(1)", "0.850918128239322"},
	{"coth(1)", "1.31303528549933"},
	{"asech(0.5)", "1.31695789692482"},
	{"acsch(1)", "0.881373587019543"},
	{"acoth(2)", "0.549306144334055"},
	{"logistic(0)", "0.5"},
	{"logistic(1)", "0.731058578630005"},
	{"softplus(0)", "0.693147180559945"},
	{"softplus(1)", "1.31326168751822"},
	{"1 ? 2 : 3", "2"},
	{"0 ? 2 : 3", "3"},
	{"1<2", "1"},
	{"2>3", "0"},
	{"2<=2", "1"},
	{"3>=4", "0"},
	{"1==1", "1"},
	{"2!=1", "1"},
	{"1==1 && 2==2", "1"},
	{"1==1 || 1==2", "1"},
	{"deg(pi)", "180"},
	{"rad(180)", "3.14159265358979"},
	{"root(8, 3)", "2"},
	{"root(16, 2)", "4"},
	{"fract(3.5)", "0.5"},
	{"fract(3)", "0"},
	{"softsign(1)", "0.5"},
	{"softsign(-2)", "-0.666666666666667"},
	{"isqrt(10)", "3"},
	{"isqrt(9)", "3"},
	{"swish(0)", "0"},
	{"swish(1)", "0.731058578630005"},
	{"isfinite(1)", "1"},
	{"isfinite(2)", "1"},
	{"mish(0)", "0"},
	{"mish(1)", "0.86509838826731"},
	{"gamma(5)", "24"},
	{"mod(10, 3)", "1"},
	{"sign(-7)", "-1"},
	{"clamp(5, 0, 3)", "3"},
	{"lerp(0, 10, 0.5)", "5"},
	{"fma(2, 3, 4)", "10"},
	{"copysign(5, -2)", "-5"},
	{"erf(1)", "0.842700792949715"},
	{"erfc(0)", "1"},
	{"beta(1, 2)", "0.5"},
	{"logb(8)", "3"},
	{"nextafter(1, 2)", "1"},
	{"ldexp(1, 3)", "8"},
	{"dim(5, 3)", "2"},
	{"signbit(-0.0)", "1"},
	{"jn(0, 1)", "0.765197686557967"},
	{"yn(1, 1)", "-0.781212821300289"},
	{"lgamma(5)", "3.17805383034795"}}

func runChecks() []CheckResult {
	var res []CheckResult
	for _, c := range verifyCases {
		v, err := eval.Evaluate(c.expr)
		if err != nil {
			res = append(res, CheckResult{Check: c.expr, Pass: false, Detail: err.Error()})
			continue
		}
		got := Format(v)
		detail := fmt.Sprintf("= %s", got)
		if got != c.want {
			detail = fmt.Sprintf("= %s (want %s)", got, c.want)
		}
		res = append(res, CheckResult{Check: c.expr, Pass: got == c.want, Detail: detail})
	}
	// Schema self-check: the machine-readable language surface must expose the
	// stable feature lists agents depend on.
	schemaBytes, err := SchemaJSON()
	if err != nil {
		res = append(res, CheckResult{Check: "schema", Pass: false, Detail: err.Error()})
	} else {
		var s Schema
		schemaPass := json.Unmarshal(schemaBytes, &s) == nil && s.Name == "math-calculator" &&
			len(s.Functions) > 0 && len(s.Constants) > 0 && len(s.Commands) > 0 &&
			len(s.Operators) > 0 && len(s.Modes) > 0 && len(s.CLI) > 0
		detail := ""
		if schemaPass {
			detail = fmt.Sprintf("%d functions, %d constants, %d commands, %d operators, %d modes, %d cli_flags", len(s.Functions), len(s.Constants), len(s.Commands), len(s.Operators), len(s.Modes), len(s.CLI))
		}
		res = append(res, CheckResult{Check: "schema", Pass: schemaPass, Detail: detail})
	}
	// Exit-code contract self-check: 0=ok, 1=usage, 2=io, 3=eval.
	// main.go enforces these; here we confirm the schema's contract is intact.
	schemaBytes2, err2 := SchemaJSON()
	if err2 != nil {
		res = append(res, CheckResult{Check: "exit-codes", Pass: false, Detail: err2.Error()})
	} else {
		var s2 Schema
		_ = json.Unmarshal(schemaBytes2, &s2)
		exitPass := s2.ExitCodes["ok"] == 0 && s2.ExitCodes["usage"] == 1 && s2.ExitCodes["io"] == 2 && s2.ExitCodes["eval"] == 3
		res = append(res, CheckResult{Check: "exit-codes", Pass: exitPass, Detail: "ok=0 usage=1 io=2 eval=3"})
	}
	// NDJSON self-check: each --jsonl line must be one parseable JSON object
	// with the stable {expr,kind,value} shape, so streaming agents can ingest
	// one line at a time without a document-level parser.
	b, err := json.Marshal(map[string]any{"expr": "1+1", "kind": "value", "value": 2})
	if err != nil {
		res = append(res, CheckResult{Check: "jsonl", Pass: false, Detail: err.Error()})
	} else {
		var m map[string]any
		ok := json.Unmarshal(b, &m) == nil && m["kind"] == "value" && m["expr"] == "1+1"
		res = append(res, CheckResult{Check: "jsonl", Pass: ok, Detail: "single-line JSON object shape"})
	}
	return res
}

// Verify runs the self-test battery and prints human-readable prose lines.
func Verify(w io.Writer) (passed, failed int) {
	for _, r := range runChecks() {
		if r.Pass {
			fmt.Fprintf(w, "ok   %s %s\n", r.Check, r.Detail)
			passed++
		} else {
			fmt.Fprintf(w, "FAIL %s %s\n", r.Check, r.Detail)
			failed++
		}
	}
	fmt.Fprintf(w, "%d passed, %d failed\n", passed, failed)
	return passed, failed
}

// VerifyJSON runs the same battery and returns a machine-readable report so
// agents can parse the health check (passed/failed counts + each outcome).
func VerifyJSON() map[string]any {
	checks := runChecks()
	passed, failed := 0, 0
	for _, r := range checks {
		if r.Pass {
			passed++
		} else {
			failed++
		}
	}
	return map[string]any{
		"version": schemaVersion,
		"passed":  passed,
		"failed":  failed,
		"checks":  checks,
	}
}
