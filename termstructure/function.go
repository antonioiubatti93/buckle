package termstructure

import "github.com/antonioiubatti93/buckle/curve"

type Func func(yf float64) float64

var _ curve.TermStructure = Func(nil)

func (f Func) Value(yf float64) float64 {
	return f(yf)
}
