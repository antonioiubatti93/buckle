package curve

// TermStructure defines a term structure.
type TermStructure interface {
	// Value evaluates the term structure at the input
	// tenor yf, expressed in year fraction units.
	Value(yf float64) float64
}

// TermStructureFunc is a term structure functional.
type TermStructureFunc func(yf float64) float64

var _ TermStructure = TermStructureFunc(nil)

// Value evaluates the term structure functional.
func (f TermStructureFunc) Value(yf float64) float64 {
	return f(yf)
}

// ShiftTermStructure applies an additive shift to the input
// term structure.
func ShiftTermStructure(ts TermStructure, shift float64) TermStructureFunc {
	return func(yf float64) float64 {
		return ts.Value(yf) + shift
	}
}
