package marketdata

import (
	"encoding/json"
	"fmt"

	"github.com/antonioiubatti93/buckle/curve"
)

// InterestRateCurve is a term structure.
type InterestRateCurve struct {
	curve.TermStructure
}

// UnmarshalJSON deserializes an interest rate curve from JSON.
func (c *InterestRateCurve) UnmarshalJSON(data []byte) error {
	var ts map[Tenor]float64
	if err := json.Unmarshal(data, &ts); err != nil {
		return fmt.Errorf("could not unmarshal data points: %w", err)
	}

	curve, err := newInterestRateCurve(ts)
	if err != nil {
		return fmt.Errorf("could not build interest rate curve from data points: %w", err)
	}
	*c = curve

	return nil
}

func newInterestRateCurve(ts TermStructure) (InterestRateCurve, error) {
	discrete, err := newDiscreteTermStructure(ts)
	if err != nil {
		return InterestRateCurve{}, err
	}

	return InterestRateCurve{
		discrete,
	}, nil
}
