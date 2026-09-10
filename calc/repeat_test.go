package calc

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	for _, tc := range []struct{ in, want string }{
		{"repeat(5, i, i*i)", "25"},
		{"repeat(3, j, j+j)", "6"},
		{"repeat(4, k, k+10)", "14"},
		{"repeat(3, i, repeat(3, j, i+j))", "6"}, // nested: last = 3+3
	} {
		got := strings.TrimSpace(runBatch(t, tc.in+"\nquit\n"))
		if got != tc.want {
			t.Fatalf("%s = %q, want %q", tc.in, got, tc.want)
		}
	}
}
