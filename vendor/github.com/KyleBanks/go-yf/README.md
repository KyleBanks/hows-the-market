# go-yf

`go-yf` is a small wrapper around the Yahoo Finance v8 API written in Go.

The Yahoo Finance API is known to be unstable with breaking changes introduced without notice, and is not officially supported for external use, so use at your own risk.

## Installation

```
$ go get github.com/KyleBanks/go-yf
```

## Usage

The library exposes a single `GetStock` function that accepts a stock symbol, range, and interval:

```go
import (
	"fmt"
	"yf"
)

stock, err := yf.GetStock("GOOG", yf.RangeOneYear, yf.IntervalOneDay)
if err != nil {
	panic(err)
}

fmt.Printf("%v: %v", stock.Meta.Symbol, stock.Indicators.Quote[0].Close)
```

## License

```
MIT License

Copyright (c) 2017 Kyle Banks

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```