package calc

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestEta(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want float64
	}{
		{"eta(1)", math.Log(2)},
		{"eta(2)", math.Pi * math.Pi / 12},
	} {
		got := strings.TrimSpace(runBatch(t, tc.in+"\nquit\n"))
		val, err := strconv.ParseFloat(got, 64)
		if err != nil {
			t.Fatalf("parse %q: %v", got, err)
		}
		if math.Abs(val-tc.want) > 1e-6 {
			t.Fatalf("%s = %v, want %v", tc.in, val, tc.want)
		}
	}
}
