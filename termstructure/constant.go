package termstructure

import "github.com/antonioiubatti93/buckle/curve"

type Constant float64

var _ curve.TermStructure = Constant(0.0)

func (c Constant) Value(_ float64) float64 {
	return float64(c)
}

func NewConstant(c float64) Constant {
	return Constant(c)
}
