package rate

import "math"

type Compounding func(r float64) float64

var (
	Continuous Compounding = func(r float64) float64 {
		return r
	}

	Simple Compounding = func(r float64) float64 {
		return math.Exp(r) - 1.0
	}
)
