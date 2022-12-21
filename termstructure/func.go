package termstructure

type Func func(yf float64) float64

var _ termStructure = Func(nil)

func (f Func) Value(yf float64) float64 {
	return f(yf)
}
