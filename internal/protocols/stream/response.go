package stream

type Response struct {
	Command string `json:"command"`
}

type KeepAliveResponse struct {
	Command string        `json:"command"`
	Data    KeepAliveData `json:"data"`
}

type KeepAliveData struct {
	Timestamp int `json:"timestamp"`
}

type ResponseCandle struct {
	Command string `json:"response"`
	Data    Candle `json:"data"`
}

// TODO: Move it to a common package (stream and socket use it)
type Candle struct {
	Close     float64 `json:"close"`
	Ctm       int64   `json:"ctm"`
	CtmString string  `json:"ctmString"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Open      float64 `json:"open"`
	QuoteId   int     `json:"quoteId"`
	Symbol    string  `json:"symbol"`
	Vol       float64 `json:"vol"`
}
