package interestrate

import "math"

type InterestRate struct {
	ts TermStructure
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

func New(ts TermStructure) InterestRate {
	return InterestRate{
		ts: ts,
	}
}
