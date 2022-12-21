package pricing

type InterestRate interface {
	DiscountAt(yf float64) float64
}
