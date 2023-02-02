package termstructure

type termStructure interface {
	Value(yf float64) float64
}
