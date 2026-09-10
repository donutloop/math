package calc

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestLogGamma(t *testing.T) {
	// loggamma(1)=0, loggamma(5)=ln(24), loggamma(0.5)=ln(sqrt(pi))
	for _, tc := range []struct {
		in   string
		want float64
	}{
		{"loggamma(1)", 0},
		{"loggamma(5)", math.Log(24)},
		{"loggamma(0.5)", math.Log(math.Sqrt(math.Pi))},
	} {
		got := strings.TrimSpace(runBatch(t, tc.in+"\nquit\n"))
		val, err := strconv.ParseFloat(got, 64)
		if err != nil {
			t.Fatalf("parse %q: %v", got, err)
		}
		if math.Abs(val-tc.want) > 1e-9 {
			t.Fatalf("%s = %v, want %v", tc.in, val, tc.want)
		}
	}
}
