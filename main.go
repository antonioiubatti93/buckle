package main

import (
	"flag"
	"fmt"

	"github.com/antonioiubatti93/buckle/interestrate"
	"github.com/antonioiubatti93/buckle/pricing"
	"github.com/antonioiubatti93/buckle/termstructure"
)

func main() {
	fmt.Println("Application: price a zero-coupon bond")

	notional := flag.Float64("n", 1.0, "zero-coupon bond notional")
	rate := flag.Float64("r", 0.01, "constant interest rate")
	maturity := flag.Int("m", 1, "maturity [years]")
	flag.Parse()

	ts := termstructure.NewConstant(*rate)
	ir := interestrate.New(ts)
	bond := pricing.NewBond(*notional, float64(*maturity))

	fmt.Printf("notional: %0.3f\n", *notional)
	fmt.Printf("interest rate: %0.3f%%\n", *rate*100.0)
	fmt.Printf("maturity: %d years\n", *maturity)

	fmt.Printf("price: %0.6f\n", bond.Price(ir))
}
