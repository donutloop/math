package calc

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestHarmonic(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want float64
	}{
		{"harmonic(1)", 1},
		{"harmonic(3)", 1.8333333333333333},
		{"harmonic(10)", 2.9289682539682538},
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
