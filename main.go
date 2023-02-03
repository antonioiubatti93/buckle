package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/antonioiubatti93/buckle/marketdata"
	"github.com/antonioiubatti93/buckle/marketstate"
)

func main() {
	var rate marketdata.FloatingRateCurve
	if err := json.Unmarshal([]byte(`{
		"termStructure": {
			"ON": 0.01,
			"M1": 0.21,
			"Y1": 0.03
		},
		"rate": "Forward",
		"horizon": "Y1",
		"spread": 0.01,
		"compounding": "Simple"
	}`), &rate); err != nil {
		log.Fatal(err)
	}

	fr := marketstate.NewFloatingRate(rate, rate.Spread())

	fmt.Println("forward rate at 1y:", fr.Value(1.0))

	var curve marketdata.InterestRateCurve
	if err := json.Unmarshal([]byte(`{
		"M1": 0.01,
		"M6": 0.02,
		"Y1": 0.03
	}`), &curve); err != nil {
		log.Fatal(err)
	}

	ir := marketstate.NewInterestRate(curve)

	fmt.Println("discount at 1y:", ir.Discount(1.0))
}
