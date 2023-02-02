package rate

type floating interface {
	Compute(yf float64) float64
}
