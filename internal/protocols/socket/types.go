package socket

type ChartResponse struct {
	Digits    int      `json:"digits"`
	RateInfos []Candle `json:"rateInfos"`
}

type Candle struct {
	Close     float64 `json:"close"`
	Ctm       int64   `json:"ctm"`
	CtmString string  `json:"ctmString"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Open      float64 `json:"open"`
	Vol       float64 `json:"vol"`
}

type Symbol struct {
	Ask                float64  `json:"ask"`
	Bid                float64  `json:"bid"`
	CategoryName       string   `json:"categoryName"`
	ContractSize       int      `json:"contractSize"`
	Currency           string   `json:"currency"`
	CurrencyPair       bool     `json:"currencyPair"`
	CurrencyProfit     string   `json:"currencyProfit"`
	Description        string   `json:"description"`
	Expiration         *int64   `json:"expiration"`
	GroupName          string   `json:"groupName"`
	High               float64  `json:"high"`
	InitialMargin      int      `json:"initialMargin"`
	InstantMaxVolume   int      `json:"instantMaxVolume"`
	Leverage           float64  `json:"leverage"`
	LongOnly           bool     `json:"longOnly"`
	LotMax             float64  `json:"lotMax"`
	LotMin             float64  `json:"lotMin"`
	LotStep            float64  `json:"lotStep"`
	Low                float64  `json:"low"`
	MarginHedged       int      `json:"marginHedged"`
	MarginHedgedStrong bool     `json:"marginHedgedStrong"`
	MarginMaintenance  *int     `json:"marginMaintenance"`
	MarginMode         int      `json:"marginMode"`
	Percentage         float64  `json:"percentage"`
	PipsPrecision      int      `json:"pipsPrecision"`
	Precision          int      `json:"precision"`
	ProfitMode         int      `json:"profitMode"`
	QuoteId            int      `json:"quoteId"`
	ShortSelling       bool     `json:"shortSelling"`
	SpreadRaw          float64  `json:"spreadRaw"`
	SpreadTable        float64  `json:"spreadTable"`
	Starting           *int64   `json:"starting"`
	StepRuleId         int      `json:"stepRuleId"`
	StopsLevel         int      `json:"stopsLevel"`
	SwapRollover3days  int      `json:"swap_rollover3days"`
	SwapEnable         bool     `json:"swapEnable"`
	SwapLong           float64  `json:"swapLong"`
	SwapShort          float64  `json:"swapShort"`
	SwapType           int      `json:"swapType"`
	Symbol             string   `json:"symbol"`
	TickSize           *float64 `json:"tickSize"`
	TickValue          *float64 `json:"tickValue"`
	Time               int64    `json:"time"`
	TimeString         string   `json:"timeString"`
	TrailingEnabled    bool     `json:"trailingEnabled"`
	Type               int      `json:"type"`
}

type CalendarRecord struct {
	Country  string `json:"country"`
	Current  string `json:"current"`
	Forecast string `json:"forecast"`
	Impact   string `json:"impact"`
	Period   string `json:"period"`
	Previous string `json:"previous"`
	Time     int64  `json:"time"`
	Title    string `json:"title"`
}

type TickRecord struct {
	Ask         float64 `json:"ask"`
	AskVolume   *int    `json:"askVolume"`
	Bid         float64 `json:"bid"`
	BidVolume   *int    `json:"bidVolume"`
	High        float64 `json:"high"`
	Level       int     `json:"level"`
	Low         float64 `json:"low"`
	SpreadRaw   float64 `json:"spreadRaw"`
	SpreadTable float64 `json:"spreadTable"`
	Symbol      string  `json:"symbol"`
	Timestamp   int64   `json:"timestamp"`
}

type TickPricesResponse struct {
	Quotations []TickRecord `json:"quotations"`
}

type QuotesRecord struct {
	Day   int   `json:"day"`
	FromT int64 `json:"fromT"`
	ToT   int64 `json:"toT"`
}

type TradingRecord struct {
	Day   int   `json:"day"`
	FromT int64 `json:"fromT"`
	ToT   int64 `json:"toT"`
}

type TradingHoursRecord struct {
	Quotes  []QuotesRecord  `json:"quotes"`
	Symbol  string          `json:"symbol"`
	Trading []TradingRecord `json:"trading"`
}

type UserData struct {
	CompanyUnit        int     `json:"companyUnit"`
	Currency           string  `json:"currency"`
	Group              string  `json:"group"`
	IbAccount          bool    `json:"ibAccount"`
	Leverage           int     `json:"leverage"`
	LeverageMultiplier float64 `json:"leverageMultiplier"`
	SpreadType         string  `json:"spreadType"`
	TrailingStop       bool    `json:"trailingStop"`
}

type MarginLevel struct {
	Balance     float64 `json:"balance"`
	Credit      float64 `json:"credit"`
	Currency    string  `json:"currency"`
	Equity      float64 `json:"equity"`
	Margin      float64 `json:"margin"`
	MarginFree  float64 `json:"margin_free"`
	MarginLevel float64 `json:"margin_level"`
}

type NewsTopic struct {
	Body       string `json:"body"`
	BodyLen    int    `json:"bodylen"`
	Key        string `json:"key"`
	Time       int64  `json:"time"`
	TimeString string `json:"timeString"`
	Title      string `json:"title"`
}

type IbRecord struct {
	ClosePrice *float64 `json:"closePrice"`
	Login      *string  `json:"login"`
	Nominal    *float64 `json:"nominal"`
	OpenPrice  *float64 `json:"openPrice"`
	Side       *int     `json:"side"`
	Surname    *string  `json:"surname"`
	Symbol     *string  `json:"symbol"`
	Timestamp  *int64   `json:"timestamp"`
	Volume     *float64 `json:"volume"`
}

type ServerTime struct {
	Time       int64  `json:"time"`
	TimeString string `json:"timeString"`
}

type StepRecord struct {
	FromValue float64 `json:"fromValue"`
	Step      float64 `json:"step"`
}

type StepRule struct {
	Id    int          `json:"id"`
	Name  string       `json:"name"`
	Steps []StepRecord `json:"steps"`
}

type Commission struct {
	Commission     float64 `json:"commission"`
	RateOfExchange float64 `json:"rateOfExchange"`
}

type Profit struct {
	Profit float64 `json:"profit"`
}

type Version struct {
	Version string `json:"version"`
}

type TradeRecord struct {
	ClosePrice       float64  `json:"close_price"`
	CloseTime        *int64   `json:"close_time"`
	CloseTimeString  *string  `json:"close_timeString"`
	Closed           bool     `json:"closed"`
	Cmd              int      `json:"cmd"`
	Comment          string   `json:"comment"`
	Commission       *float64 `json:"commission"`
	CustomComment    string   `json:"customComment"`
	Digits           int      `json:"digits"`
	Expiration       *int64   `json:"expiration"`
	ExpirationString *string  `json:"expirationString"`
	MarginRate       float64  `json:"margin_rate"`
	Offset           int      `json:"offset"`
	OpenPrice        float64  `json:"open_price"`
	OpenTime         int64    `json:"open_time"`
	OpenTimeString   string   `json:"open_timeString"`
	Order            int      `json:"order"`
	Order2           int      `json:"order2"`
	Position         int      `json:"position"`
	Profit           float64  `json:"profit"`
	Sl               float64  `json:"sl"`
	Storage          float64  `json:"storage"`
	Symbol           *string  `json:"symbol"`
	Timestamp        int64    `json:"timestamp"`
	Tp               float64  `json:"tp"`
	Volume           float64  `json:"volume"`
}

type TradeTransInfo struct {
	Cmd           int     `json:"cmd"`
	CustomComment string  `json:"customComment"`
	Expiration    int64   `json:"expiration"`
	Offset        int     `json:"offset"`
	Order         int     `json:"order"`
	Price         float64 `json:"price"`
	Sl            float64 `json:"sl"`
	Symbol        string  `json:"symbol"`
	Tp            float64 `json:"tp"`
	Type          int     `json:"type"`
	Volume        float64 `json:"volume"`
}

type TradeTransactionResponse struct {
	Order int `json:"order"`
}

type TradeTransactionStatusResponse struct {
	Ask           float64 `json:"ask"`
	Bid           float64 `json:"bid"`
	CustomComment string  `json:"customComment"`
	Message       string  `json:"message"`
	Order         int     `json:"order"`
	RequestStatus int     `json:"requestStatus"`
}
