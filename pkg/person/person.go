package person

import "fmt"

type Person struct {
	name string
	age  int
}

func New(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func (p Person) String() string {
	return fmt.Sprintf("name: %s, age: %d", p.name, p.age)
}
