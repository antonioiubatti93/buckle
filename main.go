package main

import (
	"flag"
	"fmt"

	"github.com/antonioiubatti93/buckle/pkg/person"
)

func main() {
	name := flag.String("name", "", "person name")
	age := flag.Int("age", 0, "person age")
	flag.Parse()

	p := person.New(*name, *age)

	fmt.Println(p)
}
