package curve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TermStructureFunc(t *testing.T) {
	t.Parallel()

	const (
		r  = 0.01
		yf = 1.0
	)

	assert.Equal(t, r, TermStructureFunc(func(_ float64) float64 {
		return r
	}).Value(1.0))
}

func Test_ShiftTermStructure(t *testing.T) {
	t.Parallel()

	const (
		r     = 0.01
		yf    = 1.0
		shift = 0.01
		tol   = 1.0e-15
	)

	ts := TermStructureFunc(func(_ float64) float64 {
		return r
	})

	assert.InDelta(t, r+shift, ShiftTermStructure(ts, shift).Value(1.0), tol)
}
