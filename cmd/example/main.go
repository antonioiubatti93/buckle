package main

import (
	"flag"
	"fmt"

	"github.com/antonioiubatti93/buckle/pkg/person"
)

func main() {
	fmt.Println("Using a previous, deleted but tagged version")

	name := flag.String("name", "", "person name")
	age := flag.Int("age", 0, "person age")
	flag.Parse()

	fmt.Println("base:", person.New(*name, *age))
}
