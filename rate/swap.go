package rate

import (
	"math"

	"github.com/antonioiubatti93/buckle/curve"
)

type Swap struct {
	ts       curve.TermStructure
	maturity float64
	period   float64
}

var _ curve.FloatingRate = Swap{}

func (f Swap) Compute(yf float64) float64 {
	return (f.discountFactor(yf) - f.discountFactor(yf+f.maturity)) / f.computeSwapFormulaDenominator(yf)
}

func (f Swap) computeSwapFormulaDenominator(yf float64) float64 {
	n := int(f.maturity / f.period)
	lastDelta := f.maturity - float64(n)*f.period

	value := lastDelta * f.discountFactor(yf+f.maturity)
	for i := 1; i <= n; i++ {
		value += f.period * f.discountFactor(yf+float64(i)*f.period)
	}

	return value
}

func (f Swap) discountFactor(yf float64) float64 {
	return math.Exp(-f.ts.Value(yf) * yf)
}

func NewSwap(ts curve.TermStructure, maturity, period float64) Swap {
	return Swap{
		ts:       ts,
		maturity: maturity,
		period:   period,
	}
}
