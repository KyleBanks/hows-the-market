package main

import (
	"flag"
	"strings"
)

type args struct {
	indices        []index
	showTimes      bool
	showPercentage bool
}

func (a *args) setIndices(symbols string) {
	if len(symbols) == 0 {
		a.indices = defaultIndices
		return
	}

	for _, s := range strings.Split(symbols, ",") {
		s = strings.TrimSpace(s)
		if len(s) == 0 {
			continue
		}

		a.indices = append(a.indices, index{s, s})
	}
}

func (a args) validate() bool {
	if len(a.indices) == 0 {
		return false
	}

	return true
}

func newArgs() args {
	var a args
	var symbols string

	flag.StringVar(&symbols, "symbols", "", "A comma separated list of symbols/indices to use as indicators.")
	flag.BoolVar(&a.showPercentage, "percent", false, "If set, displays the percentage increase/decrease of prices during the current trading period.")
	flag.BoolVar(&a.showTimes, "time", false, "If set, displays the market time when prices were loaded.")
	flag.Parse()

	a.setIndices(symbols)
	return a
}
