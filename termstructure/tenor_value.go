package termstructure

import (
	"fmt"

	interp "github.com/edgelaboratories/interpolator"
	"golang.org/x/exp/slices"
)

type Tenor int

func (t Tenor) Value() float64 {
	return float64(t) / 12.0
}

type interpolator interface {
	Value(x float64) float64
}

type TenorValue struct {
	interpolator interpolator
}

var _ termStructure = TenorValue{}

func (tv TenorValue) Value(yf float64) float64 {
	return tv.interpolator.Value(yf)
}

func NewTenorValue(tenorValues map[Tenor]float64) (*TenorValue, error) {
	interpolator, err := interp.NewPiecewiseLinearThreshold(convertTenorValuesIntoDataPoints(tenorValues))
	if err != nil {
		return nil, fmt.Errorf("could not create interpolator on data points: %w", err)
	}

	return &TenorValue{
		interpolator: interpolator,
	}, nil
}

func convertTenorValuesIntoDataPoints(tenorValues map[Tenor]float64) interp.XYs {
	xys := make(interp.XYs, len(tenorValues))
	for tenor, value := range tenorValues {
		xys = append(xys, interp.XY{
			X: tenor.Value(),
			Y: value,
		})
	}

	slices.SortFunc(xys, func(a, b interp.XY) bool {
		return a.X < b.X
	})

	return xys
}
