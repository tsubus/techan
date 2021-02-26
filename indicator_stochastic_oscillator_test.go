package techan

import (
	"math"
	"testing"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

var fastStochValues = []float64{
	1,
	1,
	12.0 / 16.0,
	2.0 / 16.0,
	6.0 / 16.0,
	2.0 / 16.0,
	3.0 / 15.0,
	2.0 / 16.0,
	4.0 / 13.0,
	11.0 / 17.0,
	24.0 / 49.0,
}

func TestFastStochasticIndicator(t *testing.T) {
	ts := mockTimeSeriesOCHL(
		[]string{"10", "12", "12", "8"},
		[]string{"11", "14", "14", "9"},
		[]string{"10", "20", "24", "10"},
		[]string{"9", "10", "11", "9"},
		[]string{"11", "14", "14", "9"},
		[]string{"9", "10", "11", "9"},
		[]string{"10", "12", "12", "10"},
		[]string{"9", "10", "11", "8"},
		[]string{"6", "5", "8", "1"},
		[]string{"15", "12", "18", "9"},
		[]string{"35", "25", "50", "20"},
	)

	window := 6

	closePrice := NewClosePriceIndicator(ts)
	minValue := NewMinimumValueIndicator(NewLowPriceIndicator(ts), window)
	maxValue := NewMaximumValueIndicator(NewHighPriceIndicator(ts), window)

	k := NewFastStochasticIndicator(closePrice, minValue, maxValue, window)

	decimalEquals(t, fastStochValues[0], k.Calculate(0))
	decimalEquals(t, fastStochValues[1], k.Calculate(1))
	decimalEquals(t, fastStochValues[2], k.Calculate(2))
	decimalEquals(t, fastStochValues[3], k.Calculate(3))
	decimalEquals(t, fastStochValues[4], k.Calculate(4))
	decimalEquals(t, fastStochValues[5], k.Calculate(5))
	decimalEquals(t, fastStochValues[6], k.Calculate(6))
	decimalEquals(t, fastStochValues[7], k.Calculate(7))
	decimalEquals(t, fastStochValues[8], k.Calculate(8))
	decimalEquals(t, fastStochValues[9], k.Calculate(9))
	decimalEquals(t, fastStochValues[10], k.Calculate(10))
}

func TestSlowStochasticIndicator(t *testing.T) {
	ts := mockTimeSeriesFl(fastStochValues...)

	window := 3

	d := NewSlowStochasticIndicator(NewClosePriceIndicator(ts), window)

	decimalEquals(t, 1, d.Calculate(0))
	decimalEquals(t, 1, d.Calculate(1))
	decimalEquals(t, (12.0/16.0+1+1)/3.0, d.Calculate(2))
	decimalEquals(t, (2.0/16.0+12.0/16.0+1)/3.0, d.Calculate(3))
	decimalEquals(t, (6.0/16.0+2.0/16.0+12.0/16.0)/3.0, d.Calculate(4))
	decimalEquals(t, (2.0/16.0+6.0/16.0+2.0/16.0)/3.0, d.Calculate(5))
	decimalEquals(t, (3.0/15.0+2.0/16.0+6.0/16.0)/3.0, d.Calculate(6))
	decimalEquals(t, (2.0/16.0+3.0/15.0+2.0/16.0)/3.0, d.Calculate(7))
	decimalEquals(t, (4.0/13.0+2.0/16.0+3.0/15.0)/3.0, d.Calculate(8))
	decimalEquals(t, (11.0/17.0+4.0/13.0+2.0/16.0)/3.0, d.Calculate(9))
	decimalEquals(t, (24.0/49.0+11.0/17.0+4.0/13.0)/3.0, d.Calculate(10))
}

func TestFastStochasticIndicatorNoPriceChange(t *testing.T) {
	price := NewClosePriceIndicator(mockTimeSeries("42.0", "42.0"))
	k := NewFastStochasticIndicator(price, price, price, 2)
	assert.Equal(t, big.NewDecimal(math.MaxFloat64).FormattedString(2), k.Calculate(1).FormattedString(2))
}
