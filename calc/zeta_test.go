package calc

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestZeta(t *testing.T) {
	pi2 := 1.64493406684822643647241516664602518921894990120680
	got := runBatch(t, "zeta(2)\nquit\n")
	got = strings.TrimSpace(got)
	val, err := strconv.ParseFloat(got, 64)
	if err != nil {
		t.Fatalf("parse zeta(2) %q: %v", got, err)
	}
	if math.Abs(val-pi2) > 1e-6 {
		t.Fatalf("zeta(2) = %v, want %v", val, pi2)
	}
	apery := 1.202056903159594285399738161511449990764986292345
	got = runBatch(t, "zeta(3)\nquit\n")
	got = strings.TrimSpace(got)
	val, err = strconv.ParseFloat(got, 64)
	if err != nil {
		t.Fatalf("parse zeta(3) %q: %v", got, err)
	}
	if math.Abs(val-apery) > 1e-9 {
		t.Fatalf("zeta(3) = %v, want %v", val, apery)
	}
}
