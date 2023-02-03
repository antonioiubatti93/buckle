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

	spread float64
}

func (c FloatingRateCurve) Spread() float64 {
	return c.spread
}

func (f *FloatingRateCurve) UnmarshalJSON(data []byte) error {
	var floatingRateData struct {
		TermStructure TermStructure `json:"termStructure"`
		Spread        float64       `json:"spread"`
		Rate          string        `json:"rate"`
	}
	if err := json.Unmarshal(data, &floatingRateData); err != nil {
		return fmt.Errorf("could not unmarshal floating rate data: %w", err)
	}

	rate, err := newFloatingRate(data, floatingRateData.TermStructure, floatingRateData.Rate)
	if err != nil {
		return fmt.Errorf("could not build floating rate: %w", err)
	}

	*f = FloatingRateCurve{
		rate,
		floatingRateData.Spread,
	}

	return nil
}

func newFloatingRate(data []byte, ts TermStructure, rate string) (curve.FloatingRate, error) {
	switch rate {
	case "Forward":
		return newForwardRate(data, ts)

	case "Swap":
		return newSwapRate(data, ts)

	default:
		return nil, fmt.Errorf("rate %s not supported", rate)
	}
}

func newForwardRate(data []byte, ts TermStructure) (rate.Forward, error) {
	var forwardData struct {
		Horizon     Tenor       `json:"horizon"`
		Compounding Compounding `json:"compounding"`
	}
	if err := json.Unmarshal(data, &forwardData); err != nil {
		return rate.Forward{}, fmt.Errorf("could not unmarshal forward data: %w", err)
	}

	discrete, err := newDiscreteTermStructure(ts)
	if err != nil {
		return rate.Forward{}, fmt.Errorf("could not build term structure: %w", err)
	}

	horizon, err := forwardData.Horizon.YearFraction()
	if err != nil {
		return rate.Forward{}, fmt.Errorf("could not build horizon: %w", err)
	}

	compounding := rate.Continuous
	if forwardData.Compounding == Simple {
		compounding = rate.Simple
	}

	return rate.NewForward(discrete, horizon, compounding), nil
}

func newSwapRate(data []byte, ts TermStructure) (rate.Swap, error) {
	var swapData struct {
		Maturity Tenor `json:"maturity"`
		Period   Tenor `json:"period"`
	}
	if err := json.Unmarshal(data, &swapData); err != nil {
		return rate.Swap{}, fmt.Errorf("could not unmarshal forward data: %w", err)
	}

	discrete, err := newDiscreteTermStructure(ts)
	if err != nil {
		return rate.Swap{}, fmt.Errorf("could not build term structure: %w", err)
	}

	maturity, err := swapData.Maturity.YearFraction()
	if err != nil {
		return rate.Swap{}, fmt.Errorf("could not build maturity: %w", err)
	}

	period, err := swapData.Period.YearFraction()
	if err != nil {
		return rate.Swap{}, fmt.Errorf("could not build maturity: %w", err)
	}

	return rate.NewSwap(discrete, maturity, period), nil
}
