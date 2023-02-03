package termstructure

import "github.com/antonioiubatti93/buckle/curve"

// Func is a term structure functional.
type Func func(yf float64) float64

var _ curve.TermStructure = Func(nil)

// Value evaluates the term structure functional at a given year fraction.
func (f Func) Value(yf float64) float64 {
	return f(yf)
}
