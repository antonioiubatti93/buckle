package interestrate

import (
	"math"

	"github.com/antonioiubatti93/buckle/concepts"
)

type InterestRate struct {
	ts concepts.TermStructure
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

func New(ts concepts.TermStructure) InterestRate {
	return InterestRate{
		ts: ts,
	}
}
