package calc

import (
	"fmt"
	"prototype_kl/calc/format"
)

// SetBase selects an output radix (0 = decimal, or 2/8/16) for integral
// results. Returns an error for unsupported bases.
func (c *Calculator) SetBase(base int) error {
	if base != 0 && base != 2 && base != 8 && base != 16 {
		return fmt.Errorf("unsupported base %d", base)
	}
	c.base = base
	return nil
}

// format renders v according to the calculator's display settings.
func (c *Calculator) format(v float64) string {
	if c.base != 0 {
		return format.FormatBase(v, c.base, c.prec, c.sci)
	}
	if c.eng {
		return format.FormatEng(v, c.prec)
	}
	return format.FormatPrec(v, c.prec, c.sci)
}
