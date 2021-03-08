package techan

import (
	"math"
	"testing"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

func TestFastStochasticRSIIndicator(t *testing.T) {
	indicator := NewFastStochasticRSIIndicator(NewClosePriceIndicator(mockedTimeSeries), 4)

	expectedValues := []float64{
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		100,
		95.0049,
		47.5256,
		74.0436,
		0,
		22.5376,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestSlowStochasticRSIIndicator(t *testing.T) {
	indicator := NewSlowStochasticRSIIndicator(NewFastStochasticRSIIndicator(NewClosePriceIndicator(mockedTimeSeries),
		4), 2)

	expectedValues := []float64{
		0,
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		97.5024,
		71.2653,
		60.7846,
		37.0218,
		11.2688,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestFastStochasticRSIIndicatorNoPriceChange(t *testing.T) {
	close := NewClosePriceIndicator(mockTimeSeries("42.0", "42.0"))
	rsInd := NewRelativeStrengthIndicator(close, 2)
	assert.Equal(t, big.NewDecimal(math.Inf(1)).FormattedString(2), rsInd.Calculate(1).FormattedString(2))
}
