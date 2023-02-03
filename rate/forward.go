package rate

import (
	"math"

	"github.com/antonioiubatti93/buckle/curve"
)

type Forward struct {
	ts          curve.TermStructure
	horizon     float64
	spread      float64
	compounding Compounding
}

var _ curve.FloatingRate = Forward{}

func (f Forward) Compute(yf float64) float64 {
	yfStart, yfEnd := yf, yf+f.horizon

	return f.spread + f.computeRate(yfStart, yfEnd)
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

func NewForward(ts curve.TermStructure, horizon, spread float64, compounding Compounding) Forward {
	return Forward{
		ts:          ts,
		horizon:     horizon,
		spread:      spread,
		compounding: compounding,
	}
}
