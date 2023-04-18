package main

import (
	"fmt"
	"time"

	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
)

type Exchange struct {
	Timestamp                      int64
	Open, Close, High, Low, Volume big.Decimal
}

func main() {
	series := techan.NewTimeSeries()

	// fetch this from your preferred exchange
	dataset := []Exchange{
		{
			Timestamp: 1234567,
			Open:      big.NewFromInt(1),
			Close:     big.NewFromInt(1),
			High:      big.NewFromInt(1),
			Low:       big.NewFromInt(1),
			Volume:    big.NewFromInt(1),
		},
		{
			Timestamp: 1234667,
			Open:      big.NewFromInt(1),
			Close:     big.NewFromInt(2),
			High:      big.NewFromInt(3),
			Low:       big.NewFromInt(5),
			Volume:    big.NewFromInt(6),
		},
	}

	for _, data := range dataset {
		period := techan.NewTimePeriod(time.Unix(data.Timestamp, 0), time.Hour*24)
		series.AddCandle(
			&techan.Candle{
				Period:     period,
				OpenPrice:  data.Open,
				ClosePrice: data.Close,
				MaxPrice:   data.High,
				MinPrice:   data.Low,
				Volume:     data.Volume,
			},
		)
	}

	closePrices := techan.NewClosePriceIndicator(series)
	movingAverage := techan.NewEMAIndicator(closePrices, 10) // Create an exponential moving average with a window of 10

	// record trades on this object
	record := techan.NewTradingRecord()

	entryConstant := techan.NewConstantIndicator(30)
	exitConstant := techan.NewConstantIndicator(10)

	entryRule := techan.And(
		techan.NewCrossUpIndicatorRule(entryConstant, movingAverage),
		techan.PositionNewRule{}) // Is satisfied when the price ema moves above 30 and the current position is new

	exitRule := techan.And(
		techan.NewCrossDownIndicatorRule(movingAverage, exitConstant),
		techan.PositionOpenRule{}) // Is satisfied when the price ema moves below 10 and the current position is open

	strategy := techan.RuleStrategy{
		UnstablePeriod: 10,
		EntryRule:      entryRule,
		ExitRule:       exitRule,
	}

	fmt.Printf("Should I enter: %v\n", strategy.ShouldEnter(0, record))
	fmt.Printf("Should I exit: %v\n", strategy.ShouldExit(0, record))
}
