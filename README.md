[![GitHub Activity][commits-shield]][commits]
[![License][license-shield]](LICENSE)
[![Logo](https://github.com/peterszombati/xapi-node/raw/master/docs/xtb-logo.png)](https://www.xtb.com/en)
# xStation5 API Golang Library

This project makes it possible to get data from Forex market, execute market or limit order with Golang through WebSocket connection

This module may can be used for [X-Trade Brokers](https://www.xtb.com/en) xStation5 accounts

API documentation: [http://developers.xstore.pro/documentation](http://developers.xstore.pro/documentation)

## Disclaimer

This xStation5 API Golang Library is not affiliated with, endorsed by, or in any way officially connected to the xStation5 trading platform or its parent company. The library is provided as-is and is not guaranteed to be suitable for any particular purpose. The use of this library is at your own risk, and the author(s) of this library will not be liable for any damages arising from the use or misuse of this library.
<!-- Please refer to the license file for more information. -->

**Work in progress library**. You can start using it, but breaking changes could happen, and all the endpoints are not handled by this library.

## Usage

### Authentication
```go
xapiDemoClient, err := xapi.NewClient(os.Getenv("XAPI_USER_ID"), os.Getenv("XAPI_PASSWORD"), "demo")
// or
xapiRealClient, err := xapi.NewClient(os.Getenv("XAPI_USER_ID"), os.Getenv("XAPI_PASSWORD"), "real")
```

### GetCandles

#### With subscription

```go
xapiClient.SubscribeCandles("EURUSD")
for {
	candle := <-xapiClient.CandlesChannel
	fmt.Printf("%+v\n", candle)
}
```

#### With query

```go
end := int(time.Now().Add(-24 * 1 * time.Hour).UnixMilli())
period := 1
ticks := 50
start := int(time.Now().Add(-24 * time.Hour).UnixMilli())
candles, err := xapiClient.GetCandles(end, period, start, "EURUSD", ticks)
```

| Value | Type | Description |
| ----- | ---- | ----------- |
| end | int | End of chart block (rounded down to the nearest interval and excluding) |
| period | int | Period code |
| start | int | Start of chart block (rounded down to the nearest interval and excluding) |
| symbol | string | Symbol |
| ticks | int | Number of ticks needed |

More details here : http://developers.xstore.pro/documentation/current#getChartRangeRequest

## Contributions

Contributions and feedback are welcome! If you encounter any issues, have suggestions for improvement, or would like to contribute new features, please open an issue or submit a pull request on the GitHub repository.

[commits-shield]: https://img.shields.io/github/commit-activity/y/mateogreil/xapi-go.svg?style=for-the-badge
[commits]: https://github.com/mateogreil/xapi-go/commits/master
[license-shield]: https://img.shields.io/github/license/mateogreil/xapi-go.svg?style=for-the-badge

## Alternative

- [xapi-node](https://github.com/peterszombati/xapi-node)
- [xapi-python](https://github.com/pawelkn/xapi-python)
