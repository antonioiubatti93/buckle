package marketstate

import "github.com/antonioiubatti93/buckle/curve"

type FloatingRate struct {
	rate   curve.FloatingRate
	spread float64
}

func (f FloatingRate) Value(yf float64) float64 {
	return f.spread + f.rate.Compute(yf)
}

func NewFloatingRate(rate curve.FloatingRate, spread float64) FloatingRate {
	return FloatingRate{
		rate:   rate,
		spread: spread,
	}
}

func (f FloatingRate) Shift(shift float64) FloatingRate {
	return NewFloatingRate(curve.ShiftFloatingRate(f.rate, shift), f.spread)
}
