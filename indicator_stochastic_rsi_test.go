package techan

import (
	"math"
	"testing"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

func TestStochasticRSIIndicator(t *testing.T) {
	indicator := NewStochasticRSIIndicator(NewClosePriceIndicator(mockedTimeSeries), 3)

	expectedValues := []float64{
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		math.Inf(1),
		100,
		93.2406,
		0,
		71.4899,
		0,
		23.27,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestStochasticRSIIndicatorNoPriceChange(t *testing.T) {
	close := NewClosePriceIndicator(mockTimeSeries("42.0", "42.0"))
	rsInd := NewRelativeStrengthIndicator(close, 2)
	assert.Equal(t, big.NewDecimal(math.Inf(1)).FormattedString(2), rsInd.Calculate(1).FormattedString(2))
}
