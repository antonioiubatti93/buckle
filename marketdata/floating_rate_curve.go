package marketdata

import (
	"encoding/json"
	"fmt"

	"github.com/antonioiubatti93/buckle/curve"
	"github.com/antonioiubatti93/buckle/rate"
)

type Compounding string

const (
	Continuous Compounding = "Continuous"
	Simple     Compounding = "Simple"
)

type FloatingRateCurve struct {
	curve.FloatingRate
}

func (f *FloatingRateCurve) UnmarshalJSON(data []byte) error {
	var floatingRateData struct {
		TermStructure map[Tenor]float64 `json:"termStructure"`
		Spread        float64           `json:"spread"`
		Horizon       Tenor             `json:"horizon"`
		Compounding   Compounding       `json:"compounding"`
	}
	if err := json.Unmarshal(data, &floatingRateData); err != nil {
		return fmt.Errorf("could not unmarshal floating rate data: %w", err)
	}

	discrete, err := newDiscreteTermStructure(floatingRateData.TermStructure)
	if err != nil {
		return fmt.Errorf("could not build term structure: %w", err)
	}

	horizon, err := floatingRateData.Horizon.YearFraction()
	if err != nil {
		return fmt.Errorf("could not build horizon: %w", err)
	}

	compounding := rate.Continuous
	if floatingRateData.Compounding == Simple {
		compounding = rate.Simple
	}

	*f = FloatingRateCurve{
		rate.NewForward(discrete, horizon, floatingRateData.Spread, compounding),
	}

	return nil
}
