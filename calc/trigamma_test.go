package calc

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestTrigamma(t *testing.T) {
	// trigamma(1)=zeta(2)=pi^2/6, trigamma(2)=pi^2/6-1
	for _, tc := range []struct {
		in   string
		want float64
	}{
		{"trigamma(1)", math.Pi * math.Pi / 6},
		{"trigamma(2)", math.Pi*math.Pi/6 - 1},
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
