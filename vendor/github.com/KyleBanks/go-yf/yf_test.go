package yf

import (
	"testing"
)

func init() {
	DebugLogging = true
}

func TestGetStock(t *testing.T) {
	symbol := "GOOG"

	stock, err := GetStock(symbol, RangeOneYear, IntervalOneDay)
	if err != nil {
		t.Fatal(err)
	}

	if stock.Meta.Symbol != symbol {
		t.Fatalf("Unexpected Symbol: expected=%v, got=%v", symbol, stock.Meta.Symbol)
	}
}
