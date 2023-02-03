package curve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShiftFloatingRate(t *testing.T) {
	t.Parallel()

	const (
		r     = 0.01
		yf    = 1.0
		shift = 0.01
		tol   = 1.0e-15
	)

	f := FloatingRateFunc(func(_ float64) float64 {
		return r
	})

	assert.InDelta(t, r+shift, ShiftFloatingRate(f, shift).Compute(1.0), tol)
}
