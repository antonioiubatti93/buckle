package main

import (
	"math"
	"testing"

	"github.com/antonioiubatti93/buckle/interestrate"
	"github.com/antonioiubatti93/buckle/termstructure"
	"github.com/stretchr/testify/assert"
)

func Test_priceZeroCouponBond(t *testing.T) {
	t.Parallel()

	const (
		notional = 1.0
		rate     = 0.01
		yf       = 1.0
		tol      = 1.0e-15
	)

	ir := interestrate.New(termstructure.NewConstant(rate))

	assert.InDelta(t, notional*math.Exp(-rate*yf), priceZeroCouponBond(ir, notional, yf), tol)
}
