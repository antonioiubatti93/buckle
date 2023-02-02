package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/antonioiubatti93/buckle/interestrate"
	"github.com/antonioiubatti93/buckle/marketdata"
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

	var curve marketdata.InterestRateCurve
	if err := json.Unmarshal([]byte(`{
		"M1": 0.01,
		"M6": 0.02,
		"Y1": 0.03
	}`), &curve); err != nil {
		log.Fatal(err)
	}

	ir := interestrate.New(curve.TermStructure())

	fmt.Println("discount at 1y:", ir.Discount(1.0))
}
