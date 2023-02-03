package marketstate

import "github.com/antonioiubatti93/buckle/curve"

// FloatingRate is a floating rate with a spread.
type FloatingRate struct {
	rate   curve.FloatingRate
	spread float64
}

// Value evaluates the floating rate at a given year fraction
// as the sum of the rate value and the related spread.
func (f FloatingRate) Value(yf float64) float64 {
	return f.spread + f.rate.Compute(yf)
}

// NewFloatingRate returns a floating rate from input floating
// rate and spread.
func NewFloatingRate(rate curve.FloatingRate, spread float64) FloatingRate {
	return FloatingRate{
		rate:   rate,
		spread: spread,
	}
}

// Shift applies an additive shift to the floating rate.
func (f FloatingRate) Shift(shift float64) FloatingRate {
	return NewFloatingRate(curve.ShiftFloatingRate(f.rate, shift), f.spread)
}
