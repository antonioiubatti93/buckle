package termstructure

import "github.com/antonioiubatti93/buckle/concepts"

type Func func(yf float64) float64

var _ concepts.TermStructure = Func(nil)

func (f Func) Value(yf float64) float64 {
	return f(yf)
}
