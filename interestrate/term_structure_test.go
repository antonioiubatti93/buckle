package interestrate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShiftTermStructure(t *testing.T) {
	t.Parallel()

	const (
		rate     = 0.01
		shift    = 0.01 * rate
		yf       = 1.0
		expected = rate + shift
	)

	ts := TermStructureFunc(func(_ float64) float64 { return rate })

	assert.Equal(t, expected, ShiftTermStructure(ts, shift).Value(yf))
}

type constantTermStructure struct {
	rate float64
}

var _ TermStructure = constantTermStructure{}

func (ts constantTermStructure) Value(yf float64) float64 {
	return ts.rate
}

func newConstantTermStructure(rate float64) *constantTermStructure {
	return &constantTermStructure{
		rate: rate,
	}
}
