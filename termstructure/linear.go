package termstructure

import "github.com/antonioiubatti93/buckle/concepts"

type Linear struct {
	intercept float64
	slope     float64
}

var _ concepts.TermStructure = Linear{}

func (l Linear) Value(yf float64) float64 {
	return l.intercept + l.slope*yf
}

func NewLinear(intercept, slope float64) Linear {
	return Linear{
		intercept: intercept,
		slope:     slope,
	}
}
