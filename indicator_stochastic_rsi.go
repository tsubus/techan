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

// NewStochasticRSIIndicator returns a derivative Indicator which returns the stochastic RSI indicator for the given
// window.
// https://www.investopedia.com/terms/s/stochrsi.asp
func NewStochasticRSIIndicator(indicator Indicator, timeframe, rsiTimeframe int) Indicator {
	rsiIndicator := NewRelativeStrengthIndexIndicator(indicator, rsiTimeframe)
	return stochasticRSIIndicator{
		curRSI: rsiIndicator,
		minRSI: NewMinimumValueIndicator(rsiIndicator, rsiTimeframe),
		maxRSI: NewMaximumValueIndicator(rsiIndicator, rsiTimeframe),
		window: timeframe,
	}
}

func (k stochasticRSIIndicator) Calculate(index int) big.Decimal {
	curRSI := k.curRSI.Calculate(index)
	minRSI := k.minRSI.Calculate(index)
	maxRSI := k.maxRSI.Calculate(index)

	if minRSI.EQ(maxRSI) {
		return big.NewDecimal(math.Inf(1))
	}

	return curRSI.Sub(minRSI).Div(maxRSI.Sub(minRSI)).Mul(big.NewDecimal(100))
}

type rsiKIndicator struct {
	stochasticRSI Indicator
	window        int
}

// NewFastStochasticRSIIndicator returns a derivative Indicator which returns the fast stochastic RSI indicator (%K)
// for the given window.
func NewFastStochasticRSIIndicator(stochasticRSI Indicator, timeframe int) Indicator {
	return rsiKIndicator{stochasticRSI, timeframe}
}

func (srsi rsiKIndicator) Calculate(index int) big.Decimal {
	return NewSimpleMovingAverage(srsi.stochasticRSI, srsi.window).Calculate(index)
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
