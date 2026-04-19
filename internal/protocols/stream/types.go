package stream

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

type Tick struct {
	Ask         float64 `json:"ask"`
	AskVolume   int     `json:"askVolume"`
	Bid         float64 `json:"bid"`
	BidVolume   int     `json:"bidVolume"`
	High        float64 `json:"high"`
	Level       int     `json:"level"`
	Low         float64 `json:"low"`
	QuoteId     int     `json:"quoteId"`
	SpreadRaw   float64 `json:"spreadRaw"`
	SpreadTable float64 `json:"spreadTable"`
	Symbol      string  `json:"symbol"`
	Timestamp   int64   `json:"timestamp"`
}

type Balance struct {
	Balance     float64 `json:"balance"`
	Credit      float64 `json:"credit"`
	Equity      float64 `json:"equity"`
	Margin      float64 `json:"margin"`
	MarginFree  float64 `json:"marginFree"`
	MarginLevel float64 `json:"marginLevel"`
}

type Trade struct {
	ClosePrice      float64  `json:"close_price"`
	CloseTime       *int64   `json:"close_time"`
	CloseTimeString *string  `json:"close_timeString"`
	Closed          bool     `json:"closed"`
	Cmd             int      `json:"cmd"`
	Comment         string   `json:"comment"`
	Commission      *float64 `json:"commission"`
	CustomComment   string   `json:"customComment"`
	Digits          int      `json:"digits"`
	Expiration      *int64   `json:"expiration"`
	MarginRate      float64  `json:"margin_rate"`
	Offset          int      `json:"offset"`
	OpenPrice       float64  `json:"open_price"`
	OpenTime        int64    `json:"open_time"`
	OpenTimeString  string   `json:"open_timeString"`
	Order           int      `json:"order"`
	Order2          int      `json:"order2"`
	Position        int      `json:"position"`
	Profit          float64  `json:"profit"`
	Sl              float64  `json:"sl"`
	Storage         float64  `json:"storage"`
	Symbol          string   `json:"symbol"`
	Timestamp       int64    `json:"timestamp"`
	Tp              float64  `json:"tp"`
	Type            int      `json:"type"`
	Volume          float64  `json:"volume"`
}

type News struct {
	Body       string `json:"body"`
	BodyLen    int    `json:"bodylen"`
	Key        string `json:"key"`
	Time       int64  `json:"time"`
	TimeString string `json:"timeString"`
	Title      string `json:"title"`
}

type KeepAliveData struct {
	Timestamp int64 `json:"timestamp"`
}
