package termstructure

type termStructure interface {
	Value(yf float64) float64
}

type differentiableTermStructure interface {
	termStructure

	Gradient(yf float64) float64
}
