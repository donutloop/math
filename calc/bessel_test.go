package calc

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestBessel(t *testing.T) {
	// known values: J0(1)=0.7651976866, J1(1)=0.4400505857
	for _, tc := range []struct {
		in   string
		want float64
	}{
		{"besselj0(0)", 1},
		{"besselj1(0)", 0},
		{"besselj0(1)", 0.7651976865579666},
		{"besselj1(1)", 0.4400505857449335},
	} {
		got := strings.TrimSpace(runBatch(t, tc.in+"\nquit\n"))
		val, err := strconv.ParseFloat(got, 64)
		if err != nil {
			t.Fatalf("parse %q: %v", got, err)
		}
		if math.Abs(val-tc.want) > 1e-10 {
			t.Fatalf("%s = %v, want %v", tc.in, val, tc.want)
		}
	}
}
