package marketstate

import (
	"testing"

	"github.com/antonioiubatti93/buckle/curve"
	"github.com/stretchr/testify/assert"
)

func Test_FloatingRate_Shift(t *testing.T) {
	t.Parallel()

	const (
		rate   = 0.01
		spread = 0.005
		shift  = 0.0001
		yf     = 1.0
		tol    = 1.0e-15
	)

	f := NewFloatingRate(curve.FloatingRateFunc(func(_ float64) float64 {
		return rate
	}), spread)

	assert.InDelta(t, rate+spread, f.Value(yf), tol)
	assert.InDelta(t, rate+spread+shift, f.Shift(shift).Value(yf), tol)
}
