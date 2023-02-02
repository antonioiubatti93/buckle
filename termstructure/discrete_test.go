package termstructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewDiscrete_ShouldHaveAtLeastOnePoint(t *testing.T) {
	t.Parallel()

	_, err := NewDiscrete()
	assert.Error(t, err)
}
