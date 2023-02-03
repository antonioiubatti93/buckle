package interestrate

import (
	"math"

	"github.com/antonioiubatti93/buckle/curve"
)

type InterestRate struct {
	ts curve.TermStructure
}

func (ir InterestRate) Spot(yf float64) float64 {
	return ir.ts.Value(yf)
}

func (ir InterestRate) Discount(yf float64) float64 {
	return math.Exp(-ir.Spot(yf) * yf)
}

func (ir InterestRate) Capitalize(yf float64) float64 {
	return 1.0 / ir.Discount(yf)
}

func New(ts curve.TermStructure) InterestRate {
	return InterestRate{
		ts: ts,
	}
}

func (ir InterestRate) Shift(shift float64) InterestRate {
	return New(curve.ShiftTermStructure(ir.ts, shift))
}
