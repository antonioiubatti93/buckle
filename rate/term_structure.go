package rate

type TermStructure interface {
	Value(yf float64) float64
}
