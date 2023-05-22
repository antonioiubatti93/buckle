package main

import (
	"fmt"

	"github.com/antonioiubatti93/buckle/person"
)

func main() {
	fmt.Println("Create a person - local package")

	person := person.NewPerson("Bob", 30)
	fmt.Println("Hello, I'm", person.Name(), "and I'm", person.Age(), "years old")
}
