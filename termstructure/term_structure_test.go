package termstructure

import (
	"testing"

	"github.com/antonioiubatti93/buckle/concepts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newDiscrete(t *testing.T, tenorValues ...TenorValue) Discrete {
	d, err := NewDiscrete(tenorValues...)
	require.NoError(t, err)

	return d
}

func Test_termStructure_Value(t *testing.T) {
	t.Parallel()

	const (
		yf  = 1.0
		tol = 1.0e-15
	)

	for _, tc := range []struct {
		name     string
		ts       concepts.TermStructure
		expected float64
	}{
		{
			"constant",
			NewConstant(0.01),
			0.01,
		},
		{
			"function",
			Func(func(yf float64) float64 {
				return 0.01 * (1.0 + yf)
			}),
			0.02,
		},
		{
			"discrete",
			newDiscrete(t,
				NewTenorValue(0.0, 0.01),
				NewTenorValue(2.0, 0.02),
			),
			0.015,
		},
		{
			"linear",
			NewLinear(0.01, 0.01),
			0.02,
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, tc.expected, tc.ts.Value(yf), tol)
		})
	}
}
