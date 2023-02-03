package marketstate

import (
	"math"

	"github.com/antonioiubatti93/buckle/curve"
)

// InterestRate represents an interest rate.
type InterestRate struct {
	ts curve.TermStructure
}

// Spot returns the instant rate at a given year fraction.
func (ir InterestRate) Spot(yf float64) float64 {
	return ir.ts.Value(yf)
}

// Discount returns the discount factor at a given year fraction.
func (ir InterestRate) Discount(yf float64) float64 {
	return math.Exp(-ir.Spot(yf) * yf)
}

// Capitalize returns the capitalization factor at a given year fraction.
func (ir InterestRate) Capitalize(yf float64) float64 {
	return 1.0 / ir.Discount(yf)
}

// NewInterestRate returns an interest rate from an input term structure.
func NewInterestRate(ts curve.TermStructure) InterestRate {
	return InterestRate{
		ts: ts,
	}
}

// Shift applies an additive shift to the interest rate.
func (ir InterestRate) Shift(shift float64) InterestRate {
	return NewInterestRate(curve.ShiftTermStructure(ir.ts, shift))
}
