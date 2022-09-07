package main

import (
	"flag"
	"fmt"

	"github.com/antonioiubatti93/buckle/pkg/person"
	v2 "github.com/antonioiubatti93/buckle/pkg/person/v2"
	sdk "github.com/antonioiubatti93/buckle/sdk/person"
)

func main() {
	name := flag.String("name", "", "person name")
	age := flag.Int("age", 0, "person age")
	flag.Parse()

	fmt.Println("base:", person.New(*name, *age))

	fmt.Println("v2:", v2.New(
		v2.WithName(*name),
		v2.WithAge(*age),
	))

	fmt.Println("sdk:", sdk.New(*name, *age))
}
