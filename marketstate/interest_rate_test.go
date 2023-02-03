package marketstate

import (
	"math"
	"testing"

	"github.com/antonioiubatti93/buckle/curve/test"
	"github.com/stretchr/testify/assert"
)

func Test_InterestRate(t *testing.T) {
	t.Parallel()

	const (
		rate = 0.01
		yf   = 1.0
		tol  = 1.0e-15
	)

	ir := NewInterestRate(test.NewTermStructure(rate))

	assert.InDelta(t, rate, ir.Spot(yf), tol)
	assert.InDelta(t, math.Exp(-rate), ir.Discount(yf), tol)
	assert.InDelta(t, math.Exp(rate), ir.Capitalize(yf), tol)
}

func Test_InterestRate_Shift(t *testing.T) {
	t.Parallel()

	const (
		rate  = 0.01
		yf    = 1.0
		shift = 0.001
		tol   = 1.0e-15
	)

	ir := NewInterestRate(test.NewTermStructure(rate))

	assert.InDelta(t, rate+shift, ir.Shift(shift).Spot(yf), tol)
}