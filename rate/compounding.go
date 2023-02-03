package rate

import "math"

// Compounding defines the compounding transformation
// from continuous rate to the desired rate.
type Compounding func(r float64) float64

var (
	// Continuous is the identity function that maps
	// a continuously compounded rate into itself.
	Continuous Compounding = func(r float64) float64 {
		return r
	}

	// Simple maps continuously compounded into simply
	// compounded rates.
	Simple Compounding = func(r float64) float64 {
		return math.Exp(r) - 1.0
	}
)
