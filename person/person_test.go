package person

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Person(t *testing.T) {
	t.Parallel()

	const (
		name = "Bob"
		age  = 30
	)

	p := NewPerson(name, age)
	assert.Equal(t, name, p.Name())
	assert.Equal(t, age, p.Age())
}
