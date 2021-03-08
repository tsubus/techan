package techan

import (
	"math"
	"testing"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

func TestFastStochasticRSIIndicator(t *testing.T) {
	indicator := NewStochasticRSIIndicator(NewClosePriceIndicator(mockedTimeSeries), 5)

	expectedValues := []float64{
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		100,
		95.9481,
		54.5245,
		93.1791,
		0,
		21.6754,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestSlowStochasticRSIIndicator(t *testing.T) {
	indicator := NewSlowStochasticRSIIndicator(NewStochasticRSIIndicator(NewClosePriceIndicator(mockedTimeSeries),
		5), 3)

	expectedValues := []float64{
		0,
		0,
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		83.4909,
		81.2173,
		49.2346,
		38.2848,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestFastStochasticRSIIndicatorNoPriceChange(t *testing.T) {
	close := NewClosePriceIndicator(mockTimeSeries("42.0", "42.0"))
	rsInd := NewRelativeStrengthIndicator(close, 2)
	assert.Equal(t, big.NewDecimal(math.Inf(1)).FormattedString(2), rsInd.Calculate(1).FormattedString(2))
}
