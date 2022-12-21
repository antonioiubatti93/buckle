package interestrate

type TermStructure interface {
	Value(yf float64) float64
}

type TermStructureFunc func(yf float64) float64

var _ TermStructure = TermStructureFunc(nil)

func (f TermStructureFunc) Value(yf float64) float64 {
	return f(yf)
}

func ShiftTermStructure(ts TermStructure, shift float64) TermStructure {
	return TermStructureFunc(func(yf float64) float64 {
		return ts.Value(yf) + shift
	})
}
