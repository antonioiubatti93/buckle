package interestrate

import (
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

func Test_InterestRate_DiscountAt_PositiveRateImpliesDiscounting(t *testing.T) {
	t.Parallel()

	var rate float64
	fuzz.NewWithSeed(1234).Fuzz(&rate)

	ir := New(newConstantTermStructure(rate))

	const yf = 1.0

	assert.True(t, rate < 0 || ir.DiscountAt(yf) < 1.0)
}
