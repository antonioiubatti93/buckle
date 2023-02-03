package curve

type FloatingRate interface {
	Compute(yf float64) float64
}

type FloatingRateFunc func(yf float64) float64

var _ FloatingRate = FloatingRateFunc(nil)

func (f FloatingRateFunc) Compute(yf float64) float64 {
	return f(yf)
}

func ShiftFloatingRate(f FloatingRate, shift float64) FloatingRateFunc {
	return func(yf float64) float64 {
		return f.Compute(yf) + shift
	}
}
