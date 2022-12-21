package pricing

type Bond struct {
	notional float64
	maturity float64
}

func NewBond(notional, maturity float64) *Bond {
	return &Bond{
		notional: notional,
		maturity: maturity,
	}
}

func (b Bond) Price(ir InterestRate) float64 {
	return b.notional * ir.DiscountAt(b.maturity)
}
