package calc

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestLambertW(t *testing.T) {
	omega := 0.56714329040978387299996866221035554975381578718699
	got := runBatch(t, "lambertw(1)\nquit\n")
	got = strings.TrimSpace(got)
	val, err := strconv.ParseFloat(got, 64)
	if err != nil {
		t.Fatalf("could not parse lambertw(1) = %q: %v", got, err)
	}
	if math.Abs(val-omega) > 1e-10 {
		t.Fatalf("lambertw(1) = %v, want omega %v", val, omega)
	}
}
