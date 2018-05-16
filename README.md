# hows-the-market

A totally professional, stylish, sophisticated and generally badass tool for checking in on the market from your command line.

## Install

Installation from source requires a working Go environment:

``` 
$ go get -u github.com/KyleBanks/hows-the-market
```

## Usage

There's not a whole lot to it, just run `hows-the-market` to see your fate:

```
$ hows-the-market
```

![Example output](./_docs/usage1.png)

By default, `hows-the-market` uses the S&P 500, DOW and NASDAQ to measure your fortunes. If you prefer to use your own symbols ([and you probably should](https://www.investopedia.com/articles/investing/010917/opinion-dow-stupid.asp)), feel free to do so using the `--symbols` argument:

```
$ hows-the-market --symbols BRK-B,V,^VIX,WM,XAW.TO  
```

![Custom symbols output](./_docs/usage2.png)

**Note**: `hows-the-market` uses [Yahoo! Finance](https://finance.yahoo.com) under-the-hood, so you'll need to make sure your symbols are consistent with theirs. For most US stocks the symbol is exactly what you'd expect, however for indices and OTC stocks you'll want to grab the symbol from the Yahoo! URL or beside the name like so:

![How to grab special symbols from Yahoo! Finance](./_docs/special_symbols.png)

## Disclaimer

`hows-the-market` uses the Yahoo! Finance API, which is known to be unstable with breaking changes introduced without notice, and is not officially supported for external use, so use at your own risk. Oh, it's likely also against their Terms of Service, so consider yourself warned.

## License

`hows-the-market` is made available under the [MIT License](./LICENSE).
