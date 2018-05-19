package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/KyleBanks/go-yf"
	"github.com/fatih/color"
	"github.com/kyokomi/emoji"
)

func main() {
	a := newArgs()
	if !a.validate() {
		flag.Usage()
		return
	}

	if !validateDay(time.Now().Weekday()) {
		emoji.Println(marketsClosedMessage)
	}

	positiveCount, err := processIndices(a)
	if err != nil {
		panic(err)
	}

	printFeelings(positiveCount, len(a.indices))
}

func validateDay(d time.Weekday) bool {
	return d != time.Saturday && d != time.Sunday
}

func processIndices(a args) (int, error) {
	var positiveCount int
	for _, i := range a.indices {
		positive, err := process(a, i)
		if err != nil {
			return 0, err
		}

		if positive {
			positiveCount++
		}
	}
	return positiveCount, nil
}

func process(a args, i index) (bool, error) {
	// Use a range of one week in case today is a weekend or holiday, this
	// at least provides the most recent day.
	index, err := yf.GetStock(i.symbol, "1w", yf.IntervalOneDay)
	if err != nil || len(index.Indicators.Quote) == 0 || len(index.Indicators.Quote[0].Open) == 0 {
		colorNegative.Printf("%v:\n  Data Unavailable\n", i.symbol)
		return false, nil
	}

	open := index.Indicators.Quote[0].Open[0]
	close := index.Indicators.Quote[0].Close[0]

	var positive bool
	var c *color.Color
	if close >= open {
		positive = true
		c = colorPositive
	} else {
		positive = false
		c = colorNegative
	}

	var pctStr string
	if a.showPercentage {
		pct := (close - open) / open * 100
		pctStr = fmt.Sprintf("(%.02f%%)", pct)
	}

	var timeStr string
	if a.showTimes {
		t := time.Unix(index.Timestamp[0], 0)
		loc, err := time.LoadLocation(index.Meta.ExchangeTimezoneName)
		if err == nil && loc != nil {
			t = t.In(loc)
		}
		timeStr = fmt.Sprintf("  %v\n", t.Format(dateTimeFormat))
	}

	c.Printf("%v:\n  $%.02f -> $%.02f %v\n%v", i.display, open, close, pctStr, timeStr)
	return positive, nil
}

func printFeelings(positiveCount, numIndices int) {
	var s string
	if positiveCount == numIndices {
		s = feelingsPerfect
	} else if positiveCount > numIndices/2 && numIndices > 1 {
		s = feelingsGood
	} else if positiveCount >= 1 {
		s = feelingsBad
	} else {
		s = feelingsAwful
	}
	emoji.Println("\n " + s)
}
