package test

import "github.com/antonioiubatti93/buckle/concepts"

type Constant float64

var _ concepts.TermStructure = Constant(0.0)

func (c Constant) Value(_ float64) float64 {
	return float64(c)
}

func New(c float64) Constant {
	return Constant(c)
}
