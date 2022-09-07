package v2

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("name: %s, age: %d", p.name, p.age)
}

type Option func(*Person)

func WithName(name string) Option {
	return func(p *Person) {
		p.name = name
	}
}

func WithAge(age int) Option {
	return func(p *Person) {
		p.age = age
	}
}

func New(opts ...Option) *Person {
	p := &Person{}
	for _, opt := range opts {
		opt(p)
	}

	return p
}
