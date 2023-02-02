package termstructure

import (
	"fmt"

	interp "github.com/edgelaboratories/interpolator"
)

type interpolator interface {
	Value(yf float64) float64
}

type Discrete struct {
	interpolator interpolator
}

var _ termStructure = Discrete{}

func (d Discrete) Value(yf float64) float64 {
	return d.interpolator.Value(yf)
}

func NewDiscrete(tenorValues ...TenorValue) (Discrete, error) {
	xys := convertTenorValuesToXYs(tenorValues...)

	interpolator, err := interp.NewPiecewiseLinearThreshold(xys)
	if err != nil {
		return Discrete{}, fmt.Errorf("could not build interpolator: %w", err)
	}

	return Discrete{
		interpolator: interpolator,
	}, nil
}

func convertTenorValuesToXYs(tenorValues ...TenorValue) interp.XYs {
	xys := make(interp.XYs, 0, len(tenorValues))
	for _, tv := range tenorValues {
		xys = append(xys, interp.XY{
			X: tv.Tenor(),
			Y: tv.Value(),
		})
	}

	return xys
}
