package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Currency string

type Country struct {
	currency Currency
}

func NewCountry(currency Currency) *Country {
	return &Country{
		currency: currency,
	}
}

func (c Country) Currency() Currency {
	return c.currency
}

const (
	USD = Currency("USD")
	EUR = Currency("EUR")
	CHF = Currency("CHF")
)

func IsUS(c *Country) bool {
	return c.Currency() == USD
}

func Test_Currency(t *testing.T) {
	t.Parallel()

	// Adding or modifying one currency requires two changes
	// e.g. CHF -> RUB must be replaced twice.

	assert.Equal(t, USD, NewCountry(USD).Currency())
	assert.Equal(t, EUR, NewCountry(EUR).Currency())
	assert.Equal(t, CHF, NewCountry(CHF).Currency())
}

func Test_Currency_WithConst(t *testing.T) {
	t.Parallel()

	const currency = USD

	// This test has an interesting property:
	// it works for ANY currency, because it's controlled
	// by just one test variable.
	// To change the currency one needs one move. This is because
	// this test is about a property that does not depend on
	// the currency itself.

	assert.Equal(t, currency, NewCountry(currency).Currency())
}

func Test_Currency_WithLoop(t *testing.T) {
	t.Parallel()

	// Bring it to the next level: the test above highlights
	// the working conditions of this property, so one can
	// factorize it more easily and repeat it.

	for _, currency := range []Currency{
		USD,
		EUR,
		CHF,
	} {
		assert.Equal(t, currency, NewCountry(currency).Currency())
	}
}

func Test_IsUS(t *testing.T) {
	t.Parallel()

	// In this case instead you can't factorize the test,
	// because the assertion does depend on the currency.
	// This property is currency-dependent.

	assert.True(t, IsUS(NewCountry(USD)))
	assert.False(t, IsUS(NewCountry(EUR)))

	// One may be able to do this at most.

	t.Run("us", func(t *testing.T) {
		t.Parallel()

		assert.True(t, IsUS(NewCountry(USD)))
	})

	t.Run("not us", func(t *testing.T) {
		t.Parallel()

		for _, currency := range []Currency{
			EUR,
			CHF,
		} {
			assert.False(t, IsUS(NewCountry(currency)))
		}
	})
}
