package main

import (
	"fmt"

	"github.com/antonioiubatti93/buckle/people"
	"github.com/antonioiubatti93/buckle/person"
)

func main() {
	fmt.Println("Create a person - local package")

	person := person.New("Bob", 30)
	fmt.Println("Hello, I'm", person.Name(), "and I'm", person.Age(), "years old")

	fmt.Println("Create a person - people module")

	anotherPerson := people.New("Alice", 35)
	fmt.Println("Hello, I'm", anotherPerson.Name(), "and I'm", anotherPerson.Age(), "years old")
}
