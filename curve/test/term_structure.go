package test

import "github.com/antonioiubatti93/buckle/curve"

// NewTermStructure returns a constant term structure
// for testing purposes.
func NewTermStructure(c float64) curve.TermStructureFunc {
	return func(_ float64) float64 {
		return c
	}
}
