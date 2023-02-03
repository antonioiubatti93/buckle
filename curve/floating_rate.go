package curve

// FloatingRate represents a floating rate.
type FloatingRate interface {
	// Compute evaluates the floating rate at the input
	// tenor yf, expressed in year fraction units.
	Compute(yf float64) float64
}

// FloatingRateFunc is a floating rate functional.
type FloatingRateFunc func(yf float64) float64

var _ FloatingRate = FloatingRateFunc(nil)

// Compute evaluates the floating rate functional.
func (f FloatingRateFunc) Compute(yf float64) float64 {
	return f(yf)
}

// ShiftFloatingRate applies an additive shift to the input
// floating rate.
func ShiftFloatingRate(f FloatingRate, shift float64) FloatingRateFunc {
	return func(yf float64) float64 {
		return f.Compute(yf) + shift
	}
}
