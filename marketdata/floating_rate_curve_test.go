package marketdata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_FloatingRateData_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	var f FloatingRateCurve
	require.NoError(t, json.Unmarshal([]byte(`{
		"termStructure": {
			"M1": 0.01,
			"Y1": 0.02
		},
		"rate": "Forward",
		"spread": 0.01,
		"compounding": "Simple",
		"horizon": "Y5"
	}`), &f))
	assert.Equal(t, 0.01, f.Spread())
}

func Test_FloatingRateData_UnmarshalJSON_InvalidInput(t *testing.T) {
	t.Parallel()

	var f FloatingRateCurve
	assert.Error(t, json.Unmarshal([]byte(`{
		"termStructure": {
			"M1": "not a number"
		}
	}`), &f))
}

func Test_FloatingRateData_UnmarshalJSON_InvalidTermStructure(t *testing.T) {
	t.Parallel()

	var f FloatingRateCurve
	assert.Error(t, json.Unmarshal([]byte(`{
		"spread": 0.01,
		"compounding": "Simple",
		"horizon": "Y5"
	}`), &f))
}

func Test_FloatingRateData_UnmarshalJSON_InvalidHorizon(t *testing.T) {
	t.Parallel()

	var f FloatingRateCurve
	assert.Error(t, json.Unmarshal([]byte(`{
		"termStructure": {
			"M1": 0.01,
			"Y1": 0.02
		},
		"spread": 0.01,
		"compounding": "Simple",
		"horizon": "M3"
	}`), &f))
}
