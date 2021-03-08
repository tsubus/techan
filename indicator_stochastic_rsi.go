package techan

import (
	"math"

	"github.com/sdcoffey/big"
)

type rsiKIndicator struct {
	curRSI Indicator
	minRSI Indicator
	maxRSI Indicator
	window int
}

// NewFastStochasticRSIIndicator returns a derivative Indicator which returns the fast stochastic RSI indicator (%K)
// for the given window.
// https://www.investopedia.com/terms/s/stochrsi.asp
func NewFastStochasticRSIIndicator(indicator Indicator, timeframe int) Indicator {
	rsiIndicator := NewRelativeStrengthIndexIndicator(indicator, timeframe)
	return rsiKIndicator{
		curRSI: rsiIndicator,
		minRSI: NewMinimumValueIndicator(rsiIndicator, timeframe),
		maxRSI: NewMaximumValueIndicator(rsiIndicator, timeframe),
		window: timeframe,
	}
}

func (k rsiKIndicator) Calculate(index int) big.Decimal {
	curRSI := k.curRSI.Calculate(index)
	minRSI := k.minRSI.Calculate(index)
	maxRSI := k.maxRSI.Calculate(index)

	if minRSI.EQ(maxRSI) {
		return big.NewDecimal(math.Inf(1))
	}

	return curRSI.Sub(minRSI).Div(maxRSI.Sub(minRSI)).Mul(big.NewDecimal(100))
}

type rsiDIndicator struct {
	k      Indicator
	window int
}

// NewSlowStochasticRSIIndicator returns a derivative Indicator which returns the slow stochastic RSI indicator (%D)
// for the given window.
func NewSlowStochasticRSIIndicator(k Indicator, timeframe int) Indicator {
	return rsiDIndicator{k, timeframe}
}

func (d rsiDIndicator) Calculate(index int) big.Decimal {
	return NewSimpleMovingAverage(d.k, d.window).Calculate(index)
}
