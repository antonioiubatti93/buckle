package interestrate

import "math"

type InterestRate struct {
	ts TermStructure
}

func New(ts TermStructure) *InterestRate {
	return &InterestRate{
		ts: ts,
	}
}

func (ir InterestRate) ShortRate(yf float64) float64 {
	return ir.ts.Value(yf)
}

func (ir InterestRate) IntegralRate(yf float64) float64 {
	return ir.ShortRate(yf) * yf
}

func (ir InterestRate) DiscountAt(yf float64) float64 {
	return math.Exp(-ir.IntegralRate(yf))
}
