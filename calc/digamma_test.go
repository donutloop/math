package calc

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestDigamma(t *testing.T) {
	// digamma(1) = -EulerMascheroni, digamma(1/2) = -gamma - 2*ln2
	gamma := 0.57721566490153286060651209008240243104215933593992
	for _, tc := range []struct {
		in   string
		want float64
	}{
		{"digamma(1)", -gamma},
		{"digamma(0.5)", -gamma - 2*math.Log(2)},
	} {
		got := strings.TrimSpace(runBatch(t, tc.in+"\nquit\n"))
		val, err := strconv.ParseFloat(got, 64)
		if err != nil {
			t.Fatalf("parse %q: %v", got, err)
		}
		if math.Abs(val-tc.want) > 1e-12 {
			t.Fatalf("%s = %v, want %v", tc.in, val, tc.want)
		}
	}
}
