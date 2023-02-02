package marketdata

import (
	"encoding/json"
	"fmt"

	"github.com/antonioiubatti93/buckle/curve"
	"github.com/antonioiubatti93/buckle/termstructure"
)

type InterestRateCurve struct {
	termStructure curve.TermStructure
}

func (c InterestRateCurve) TermStructure() curve.TermStructure {
	return c.termStructure
}

func (c *InterestRateCurve) UnmarshalJSON(data []byte) error {
	var dataPoints map[Tenor]float64
	if err := json.Unmarshal(data, &dataPoints); err != nil {
		return fmt.Errorf("could not unmarshal data points: %w", err)
	}

	curve, err := newInterestRateCurve(dataPoints)
	if err != nil {
		return fmt.Errorf("could not build interest rate curve from data points: %w", err)
	}
	*c = curve

	return nil
}

func newInterestRateCurve(dataPoints map[Tenor]float64) (InterestRateCurve, error) {
	tenorValues := make([]termstructure.TenorValue, 0, len(dataPoints))
	for tenor, value := range dataPoints {
		yf, err := tenor.YearFraction()
		if err != nil {
			return InterestRateCurve{}, fmt.Errorf("could not interpret tenor: %w", err)
		}

		tenorValues = append(tenorValues, termstructure.NewTenorValue(yf, value))
	}

	discrete, err := termstructure.NewDiscrete(tenorValues...)
	if err != nil {
		return InterestRateCurve{}, fmt.Errorf("could not build discrete term structure: %w", err)
	}

	return InterestRateCurve{
		termStructure: discrete,
	}, nil
}
