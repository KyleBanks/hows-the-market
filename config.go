package main

import (
	"github.com/fatih/color"
)

const (
	feelingsPerfect = ":chart_with_upwards_trend: :money_mouth_face:"
	feelingsGood    = ":chart_with_upwards_trend: :joy:"
	feelingsBad     = ":chart_with_downwards_trend: :cry:"
	feelingsAwful   = ":chart_with_downwards_trend: :cold_sweat:"

	marketsClosedMessage = ":thinking: the markets aren't open today, but here's the last trading day:\n"
)

var (
	colorPositive = color.New(color.FgGreen)
	colorNegative = color.New(color.FgRed)
)

type index struct {
	display string
	symbol  string
}

var defaultIndices = []index{
	{"S&P 500", "^GSPC"},
	{"DOW", "^DJI"},
	{"NASDAQ", "^IXIC"},
}
