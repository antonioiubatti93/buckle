package termstructure

import "github.com/antonioiubatti93/buckle/curve"

// Linear is a linear term structure whose rate is defined
// by intercept and slope.
type Linear struct {
	intercept float64
	slope     float64
}

var _ curve.TermStructure = Linear{}

// Value evaluates the linear term structure.
func (l Linear) Value(yf float64) float64 {
	return l.intercept + l.slope*yf
}

// NewLinear returns a new linear term structure.
func NewLinear(intercept, slope float64) Linear {
	return Linear{
		intercept: intercept,
		slope:     slope,
	}
}
