package main

import (
	"flag"
	"strings"
)

func parseIndices() []index {
	var symbols string
	flag.StringVar(&symbols, "symbols", "", "[Optional] A comma separated list of symbols/indices to use as indicators.")
	flag.Parse()

	return stringsToIndices(symbols)
}

func stringsToIndices(symbols string) []index {
	if len(symbols) == 0 {
		return defaultIndices
	}

	var indices []index
	for _, s := range strings.Split(symbols, ",") {
		s = strings.TrimSpace(s)
		if len(s) == 0 {
			continue
		}
		indices = append(indices, index{s, s})
	}

	return indices
}
