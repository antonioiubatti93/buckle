package rate

import (
	"math"
	"testing"

	"github.com/antonioiubatti93/buckle/concepts"
	"github.com/antonioiubatti93/buckle/concepts/test"
	"github.com/stretchr/testify/assert"
)

func Test_Compute(t *testing.T) {
	t.Parallel()

	const (
		yf  = 1.0
		tol = 1.0e-15
	)

	ts := test.New(0.01)

	for _, tc := range []struct {
		name     string
		f        concepts.FloatingRate
		expected float64
	}{
		{
			"forward/continuous",
			NewForward(ts,
				WithCompounding(Continuous),
				WithSpread(0.01),
				WithHorizon(1.0),
			),
			0.02,
		},
		{
			"forward/simple",
			NewForward(ts,
				WithCompounding(Simple),
				WithSpread(0.01),
				WithHorizon(0.5),
			),
			0.01 + (math.Exp(0.01*0.5)-1.0)/0.5,
		},
		{
			"forward/instant",
			NewForward(ts,
				WithHorizon(0.0),
			),
			0.01,
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, tc.expected, tc.f.Compute(yf), tol)
		})
	}
}
