package techan

import (
	"math"

	"github.com/sdcoffey/big"
)

type stochasticRSIIndicator struct {
	curRSI Indicator
	minRSI Indicator
	maxRSI Indicator
	window int
}

// NewStochasticRSIIndicator returns a derivative Indicator which returns the stochastic RSI indicator for the
// given window.
// https://www.investopedia.com/terms/s/stochrsi.asp
func NewStochasticRSIIndicator(indicator Indicator, timeframe int) Indicator {
	rsiIndicator := NewRelativeStrengthIndexIndicator(indicator, timeframe)
	return stochasticRSIIndicator{
		curRSI: rsiIndicator,
		minRSI: NewMinimumValueIndicator(rsiIndicator, timeframe),
		maxRSI: NewMaximumValueIndicator(rsiIndicator, timeframe),
		window: timeframe,
	}
}

func (stochRSI stochasticRSIIndicator) Calculate(index int) big.Decimal {
	curRSI := stochRSI.curRSI.Calculate(index)
	minRSI := stochRSI.minRSI.Calculate(index)
	maxRSI := stochRSI.maxRSI.Calculate(index)

	if minRSI.EQ(maxRSI) {
		return big.NewDecimal(math.Inf(1))
	}

	return curRSI.Sub(minRSI).Div(maxRSI.Sub(minRSI)).Mul(big.NewDecimal(100))
}
