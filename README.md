[![Logo](https://github.com/peterszombati/xapi-node/raw/master/docs/xtb-logo.png)](https://www.xtb.com/en)
# xStation5 API Golang Library

This project makes it possible to get data from Forex market, execute market or limit order with Golang through WebSocket connection

This module may can be used for [X-Trade Brokers](https://www.xtb.com/en) xStation5 accounts

API documentation: [http://developers.xstore.pro/documentation](http://developers.xstore.pro/documentation)

## Disclaimer

This xStation5 API Golang Library is not affiliated with, endorsed by, or in any way officially connected to the xStation5 trading platform or its parent company. The library is provided as-is and is not guaranteed to be suitable for any particular purpose. The use of this library is at your own risk, and the author(s) of this library will not be liable for any damages arising from the use or misuse of this library.
<!-- Please refer to the license file for more information. -->

## Usage

### Authentication
```go
xapiDemoClient, err := xapi.NewClient(os.Getenv("XAPI_USER_ID"), os.Getenv("XAPI_PASSWORD"), "demo")
// or
xapiRealClient, err := xapi.NewClient(os.Getenv("XAPI_USER_ID"), os.Getenv("XAPI_PASSWORD"), "real")
```

### GetCandles

```go
xapiClient.SubscribeCandles("EURUSD")
for {
	candle := <-xapiClient.CandlesChannel
	fmt.Printf("%+v\n", candle)
}
```
