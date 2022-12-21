package pricing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type simpleInterestRate struct {
	rate float64
}

var _ InterestRate = simpleInterestRate{}

func (ir simpleInterestRate) DiscountAt(yf float64) float64 {
	return 1.0 / (1.0 + ir.rate*yf)
}

func newSimpleInterestRate(rate float64) *simpleInterestRate {
	return &simpleInterestRate{
		rate: rate,
	}
}

func Test_Bond_Price(t *testing.T) {
	t.Parallel()

	const (
		notional = 1.0
		maturity = 1.0
		rate = 0.01
		expected = 1.0 / 1.01
		tol = 1.0e-15
	)

	bond := NewBond(notional, maturity)
	ir := newSimpleInterestRate(rate)

	assert.InDelta(t, expected, bond.Price(ir), tol)
}
