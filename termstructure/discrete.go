package termstructure

import (
	"fmt"

	"github.com/antonioiubatti93/buckle/curve"
	interp "github.com/edgelaboratories/interpolator"
)

type interpolator interface {
	Value(yf float64) float64
}

// Discrete is a term structure interpolated on a set
// of discrete points with thresholded extremes.
type Discrete struct {
	interpolator interpolator
}

var _ curve.TermStructure = Discrete{}

// Value evaluates the discrete term structure at a given year fraction.
func (d Discrete) Value(yf float64) float64 {
	return d.interpolator.Value(yf)
}

// NewDiscrete returns a new discrete term structure from a set of
// tenor-value points, or an error if it is invalid.
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
