package marketdata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_InterestRateCurve_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	var c InterestRateCurve
	assert.NoError(t, json.Unmarshal([]byte(`{
		"M1": 0.01,
		"Y1": 0.02,
		"Y10": 0.03
	}`), &c))

	assert.Equal(t, 0.02, c.TermStructure().Value(1.0))
}

func Test_InterestRateCurve_UnmarshalJSON_InvalidInput(t *testing.T) {
	t.Parallel()

	var c InterestRateCurve
	assert.Error(t, json.Unmarshal([]byte(`{
		"M1": "not a number"
	}`), &c))
}

func Test_InterestRateCurve_UnmarshalJSON_InvalidTenors(t *testing.T) {
	t.Parallel()

	var c InterestRateCurve
	assert.Error(t, json.Unmarshal([]byte(`{
		"M1": 0.01,
		"M2": 0.02,
		"D1": 0.03
	}`), &c))
}

func Test_newInterestRateCurve(t *testing.T) {
	t.Parallel()

	curve, err := newInterestRateCurve(map[Tenor]float64{
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
		dataPoints map[Tenor]float64
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
