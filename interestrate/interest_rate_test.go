package interestrate

import (
	"math"
	"testing"

	"github.com/antonioiubatti93/buckle/concepts/test"
	"github.com/stretchr/testify/assert"
)

func Test_InterestRate(t *testing.T) {
	t.Parallel()

	const (
		rate = 0.01
		yf   = 1.0
		tol  = 1.0e-15
	)

	ir := New(test.New(rate))

	assert.InDelta(t, rate, ir.Spot(yf), tol)
	assert.InDelta(t, math.Exp(-rate), ir.Discount(yf), tol)
	assert.InDelta(t, math.Exp(rate), ir.Capitalize(yf), tol)
}
