package main

import (
	"fmt"
	"log"

	"github.com/antonioiubatti93/buckle/rate"
	"github.com/antonioiubatti93/buckle/termstructure"
)

func main() {
	ts, err := termstructure.NewDiscrete(
		termstructure.NewTenorValue(1.0/365.0, 0.01),
		termstructure.NewTenorValue(1.0/12.0, 0.21),
		termstructure.NewTenorValue(1.0, 0.03),
	)
	if err != nil {
		log.Fatal(err)
	}

	rate := rate.NewForward(ts,
		rate.WithCompounding(rate.Simple),
		rate.WithHorizon(1.0),
		rate.WithSpread(0.01),
	)

	fmt.Println("forward rate at 1y:", rate.Compute(1.0))
}
