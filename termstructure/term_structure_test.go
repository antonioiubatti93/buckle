package termstructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_termStructure_Value(t *testing.T) {
	t.Parallel()

	const (
		rate = 0.01
		yf   = 1.0
		tol  = 1.0e-15
	)

	for name, ts := range map[string]termStructure{
		"constant":    NewConstant(rate),
		"function":    Func(func(_ float64) float64 { return rate }),
		"tenor value": newTenorValue(t, map[Tenor]float64{Tenor(12): rate}),
	} {
		ts := ts

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, rate, ts.Value(yf), tol)
		})
	}
}

func Test_termStructure_Gradient(t *testing.T) {
	t.Parallel()

	const (
		expected = 0.0
		yf       = 1.0
		tol      = 1.0e-15
	)

	for name, ts := range map[string]differentiableTermStructure{
		"constant": NewConstant(0.01),
	} {
		ts := ts

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, expected, ts.Gradient(yf), tol)
		})
	}
}
