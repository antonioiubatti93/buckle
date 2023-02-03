package rate

import (
	"math"

	"github.com/antonioiubatti93/buckle/curve"
)

// Forward represents a forward rate of some horizon
// and given compounding.
type Forward struct {
	ts          curve.TermStructure
	horizon     float64
	compounding Compounding
}

var _ curve.FloatingRate = Forward{}

// Compute evaluates the forward rate at a given year fraction.
func (f Forward) Compute(yf float64) float64 {
	yfStart, yfEnd := yf, yf+f.horizon

	return f.computeRate(yfStart, yfEnd)
}

func (f Forward) computeRate(yfStart, yfEnd float64) float64 {
	yfDelta := yfEnd - yfStart

	const tol = 1.0e-15
	if math.Abs(yfDelta) < tol {
		return f.ts.Value(yfStart)
	}

	return f.compounding(f.integralContinuousRate(yfStart, yfEnd)) / yfDelta
}

func (f Forward) integralContinuousRate(yfStart, yfEnd float64) float64 {
	return f.ts.Value(yfEnd)*yfEnd - f.ts.Value(yfStart)*yfStart
}

// NewForward returns a new forward rate.
func NewForward(ts curve.TermStructure, horizon float64, compounding Compounding) Forward {
	return Forward{
		ts:          ts,
		horizon:     horizon,
		compounding: compounding,
	}
}
