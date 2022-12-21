package termstructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTenorValue(t require.TestingT, tenorValues map[Tenor]float64) *TenorValue {
	ts, err := NewTenorValue(tenorValues)
	require.NoError(t, err)

	return ts
}

func Test_NewTenorValue_NotEnoughPoints(t *testing.T) {
	t.Parallel()

	_, err := NewTenorValue(make(map[Tenor]float64))
	assert.Error(t, err)
}
