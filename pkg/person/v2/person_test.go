package v2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Person_String(t *testing.T) {
	const (
		name = "name"
		age  = 10
	)

	assert.Equal(t,
		fmt.Sprintf("name: %s, age: %d", name, age),
		New(
			WithName(name),
			WithAge(age),
		).String(),
	)
}
