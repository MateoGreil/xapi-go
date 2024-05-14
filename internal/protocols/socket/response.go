package socket

type LoginResponse struct {
	Status          bool   `json:"status"`
	StreamSessionId string `json:"streamSessionId"`
	ErrorCode       string `json:"errorCode"`
	ErrorDescr      string `json:"errorDescr"`
}

type Response struct {
	Status          bool   `json:"status"`
	StreamSessionId string `json:"streamSessionId"`
	ErrorCode       string `json:"errorCode"`
	ErrorDescr      string `json:"errorDescr"`
	ReturnData      struct {
		Digits    int      `json:"digits"`
		RateInfos []Candle `json:"rateInfos"`
	} `json:"returnData"`
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
	Vol       float64 `json:"vol"`
}
