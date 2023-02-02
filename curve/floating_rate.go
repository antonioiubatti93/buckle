package curve

type FloatingRate interface {
	Compute(yf float64) float64
}
