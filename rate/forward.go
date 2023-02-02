package rate

import (
	"math"

	"github.com/antonioiubatti93/buckle/concepts"
)

type Forward struct {
	ts          concepts.TermStructure
	horizon     float64
	spread      float64
	compounding Compounding
}

var _ concepts.FloatingRate = Forward{}

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

type ForwardOption func(Forward) Forward

func WithHorizon(horizon float64) ForwardOption {
	return func(f Forward) Forward {
		f.horizon = horizon
		return f
	}
}

func WithSpread(spread float64) ForwardOption {
	return func(f Forward) Forward {
		f.spread = spread
		return f
	}
}

func WithCompounding(compounding Compounding) ForwardOption {
	return func(f Forward) Forward {
		f.compounding = compounding
		return f
	}
}

func applyOpts(f Forward, opts ...ForwardOption) Forward {
	for _, opt := range opts {
		f = opt(f)
	}

	return f
}

func NewForward(ts concepts.TermStructure, opts ...ForwardOption) Forward {
	f := Forward{
		ts: ts,
	}
	f = applyOpts(f,
		WithSpread(0.0),
		WithHorizon(1.0),
		WithCompounding(Continuous),
	)

	return applyOpts(f, opts...)
}
