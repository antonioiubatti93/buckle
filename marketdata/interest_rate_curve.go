package marketdata

import (
	"encoding/json"
	"fmt"

	"github.com/antonioiubatti93/buckle/curve"
)

type InterestRateCurve struct {
	termStructure curve.TermStructure
}

func (c InterestRateCurve) TermStructure() curve.TermStructure {
	return c.termStructure
}

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
