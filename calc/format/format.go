package format

import (
	"math"
	"strconv"
	"strings"
)

func FormatBase(v float64, base, prec int, sci bool) string {
	if v != math.Trunc(v) || math.IsNaN(v) || math.IsInf(v, 0) {
		// Non-integral values fall back to decimal formatting with active precision.
		return FormatPrec(v, prec, sci)
	}
	var prefix string
	switch base {
	case 2:
		prefix = "0b"
	case 8:
		prefix = "0o"
	case 16:
		prefix = "0x"
	}
	iv := int64(v)
	if iv == 0 {
		return prefix + "0"
	}
	var neg bool
	if iv < 0 {
		neg = true
		iv = -iv
	}
	digits := "0123456789abcdef"
	var sb strings.Builder
	for iv > 0 {
		sb.WriteByte(digits[iv%int64(base)])
		iv /= int64(base)
	}
	out := sb.String()
	// reverse
	var rev strings.Builder
	for i := len(out) - 1; i >= 0; i-- {
		rev.WriteByte(out[i])
	}
	res := rev.String()
	if neg {
		return "-" + prefix + res
	}
	return prefix + res
}

func Format(v float64) string {
	return FormatPrec(v, 15, false)
}

// format renders v according to the calculator's display settings.
func FormatPrec(v float64, prec int, sci bool) string {
	if math.IsNaN(v) {
		return "NaN"
	}
	if math.IsInf(v, 1) {
		return "+Inf"
	}
	if math.IsInf(v, -1) {
		return "-Inf"
	}

	if sci {
		s := strconv.FormatFloat(v, 'e', prec-1, 64)
		if i := strings.IndexByte(s, 'e'); i >= 0 {
			mantissa := strings.TrimRight(s[:i], "0")
			mantissa = strings.TrimRight(mantissa, ".")
			s = mantissa + s[i:]
		}
		return s
	}

	s := strconv.FormatFloat(v, 'g', prec, 64)
	// Trim trailing zeros that arise from float noise, e.g. "0.30000000000000004"
	// -> "0.3". Only touch plain decimal forms, not exponent forms.
	if i := strings.IndexByte(s, '.'); i >= 0 && !strings.ContainsAny(s, "eE") {
		s = strings.TrimRight(s, "0")
		s = strings.TrimRight(s, ".")
	}
	return s
}

// FormatPrec renders v with the given significant digits, optionally in
// scientific notation. It backs the one-shot CLI display flags.
func FormatEng(v float64, prec int) string {
	if v == 0 {
		return "0"
	}
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return Format(v)
	}
	abs := math.Abs(v)
	exp := int(math.Floor(math.Log10(abs)))
	// round exp down to a multiple of 3
	exp3 := exp - (exp % 3)
	if exp3 < 0 {
		exp3 -= 3
	}
	mant := v / math.Pow10(exp3)
	s := strconv.FormatFloat(mant, 'f', prec-1, 64)
	if i := strings.IndexByte(s, '.'); i >= 0 {
		s = strings.TrimRight(s, "0")
		s = strings.TrimRight(s, ".")
	}
	return s + "e" + strconv.Itoa(exp3)
}
