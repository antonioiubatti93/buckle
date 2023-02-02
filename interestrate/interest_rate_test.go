package interestrate

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type termStructure float64

var _ TermStructure = termStructure(0.0)

func (t termStructure) Value(_ float64) float64 {
	return float64(t)
}

func newTermStructure(c float64) termStructure {
	return termStructure(c)
}

func Test_InterestRate(t *testing.T) {
	t.Parallel()

	const (
		rate = 0.01
		yf   = 1.0
		tol  = 1.0e-15
	)

	ir := New(newTermStructure(rate))

	assert.InDelta(t, rate, ir.Spot(yf), tol)
	assert.InDelta(t, math.Exp(-rate), ir.Discount(yf), tol)
	assert.InDelta(t, math.Exp(rate), ir.Capitalize(yf), tol)
}
