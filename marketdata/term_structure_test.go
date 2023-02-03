package marketdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Tenor_YearFraction(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		t        Tenor
		expected float64
	}{
		{
			ON,
			1.0 / 365.0,
		},
		{
			M1,
			1.0 / 12.0,
		},
		{
			M6,
			0.5,
		},
		{
			Y1,
			1.0,
		},
		{
			Y2,
			2.0,
		},
		{
			Y5,
			5.0,
		},
		{
			Y10,
			10.0,
		},
	} {
		tc := tc

		t.Run(string(tc.t), func(t *testing.T) {
			t.Parallel()

			yf, err := tc.t.YearFraction()
			require.NoError(t, err)
			assert.Equal(t, tc.expected, yf)
		})
	}
}

func Test_Tenor_YearFraction_Invalid(t *testing.T) {
	t.Parallel()

	_, err := Tenor("M3").YearFraction()
	assert.Error(t, err)
}

func Test_newInterestRateCurve(t *testing.T) {
	t.Parallel()

	curve, err := newInterestRateCurve(TermStructure{
		M1:  0.01,
		Y1:  0.02,
		Y10: 0.03,
	})
	require.NoError(t, err)
	assert.NotNil(t, curve.TermStructure())
}

func Test_newInterestRateCurve_Invalid(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name       string
		dataPoints TermStructure
	}{
		{
			"no points",
			nil,
		},
		{
			"unsupported tenor",
			map[Tenor]float64{
				Tenor("M4"): 0.01,
			},
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := newInterestRateCurve(tc.dataPoints)
			assert.Error(t, err)
		})
	}
}
