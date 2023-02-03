package rate

import (
	"math"
	"testing"

	"github.com/antonioiubatti93/buckle/curve"
	"github.com/antonioiubatti93/buckle/curve/test"
	"github.com/stretchr/testify/assert"
)

func Test_Compute(t *testing.T) {
	t.Parallel()

	const (
		yf  = 1.0
		tol = 1.0e-15
	)

	ts := test.NewTermStructure(0.01)

	for _, tc := range []struct {
		name     string
		f        curve.FloatingRate
		expected float64
	}{
		{
			"forward/continuous",
			NewForward(ts, 1.0, Continuous),
			0.01,
		},
		{
			"forward/simple",
			NewForward(ts, 0.5, Simple),
			(math.Exp(0.01*0.5) - 1.0) / 0.5,
		},
		{
			"forward/instant",
			NewForward(ts, 0.0, Continuous),
			0.01,
		},
		{
			"swap",
			NewSwap(ts, 5.0, 1.0),
			math.Exp(0.01) - 1.0,
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, tc.expected, tc.f.Compute(yf), tol)
		})
	}
}
