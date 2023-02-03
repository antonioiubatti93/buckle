package marketdata

import (
	"fmt"

	"github.com/antonioiubatti93/buckle/termstructure"
)

type Tenor string

const (
	ON  Tenor = "ON"
	M1  Tenor = "M1"
	M6  Tenor = "M6"
	Y1  Tenor = "Y1"
	Y2  Tenor = "Y2"
	Y5  Tenor = "Y5"
	Y10 Tenor = "Y10"
)

func (t Tenor) YearFraction() (float64, error) {
	switch t {
	case ON:
		return 1.0 / 365.0, nil

	case M1:
		return 1.0 / 12.0, nil

	case M6:
		return 0.5, nil

	case Y1:
		return 1.0, nil

	case Y2:
		return 2.0, nil

	case Y5:
		return 5.0, nil

	case Y10:
		return 10.0, nil

	default:
		return 0.0, fmt.Errorf("tenor %s not recognized", t)
	}
}

type TermStructure map[Tenor]float64

func newDiscreteTermStructure(ts TermStructure) (termstructure.Discrete, error) {
	tenorValues := make([]termstructure.TenorValue, 0, len(ts))
	for tenor, value := range ts {
		yf, err := tenor.YearFraction()
		if err != nil {
			return termstructure.Discrete{}, fmt.Errorf("could not interpret tenor: %w", err)
		}

		tenorValues = append(tenorValues, termstructure.NewTenorValue(yf, value))
	}

	discrete, err := termstructure.NewDiscrete(tenorValues...)
	if err != nil {
		return termstructure.Discrete{}, fmt.Errorf("could not build discrete term structure: %w", err)
	}

	return discrete, nil
}
