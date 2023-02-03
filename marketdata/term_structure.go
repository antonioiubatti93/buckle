package marketdata

import (
	"fmt"

	"github.com/antonioiubatti93/buckle/termstructure"
)

// Tenor represents an encoded time period.
type Tenor string

const (
	// ON is the overnight (daily) tenor.
	ON Tenor = "ON"
	// M1 is one month.
	M1 Tenor = "M1"
	// M6 is six months.
	M6 Tenor = "M6"
	// Y1 is one year.
	Y1 Tenor = "Y1"
	// Y2 is two years.
	Y2 Tenor = "Y2"
	// Y5 is five years.
	Y5 Tenor = "Y5"
	// Y10 is ten years.
	Y10 Tenor = "Y10"
)

// YearFraction returns the year fraction associated
// with the tenor, or an error if the tenor is invalid.
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

// TermStructure is a set of tenor-value data points
// representing averaged continuously compounded rates.
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
