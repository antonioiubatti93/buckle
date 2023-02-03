package termstructure

import "github.com/antonioiubatti93/buckle/curve"

// Constant is a flat term structure.
type Constant float64

var _ curve.TermStructure = Constant(0.0)

// Value evaluates the term structure at the input year fraction.
func (c Constant) Value(_ float64) float64 {
	return float64(c)
}

// NewConstant returns a new constant term structure.
func NewConstant(c float64) Constant {
	return Constant(c)
}
