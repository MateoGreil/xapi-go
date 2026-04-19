package socket

type Request struct {
	Command   string      `json:"command"`
	Arguments interface{} `json:"arguments"`
}

type LoginRequestArguments struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

type InfoArguments struct {
	Info interface{} `json:"info"`
}

type GetCandlesInfo struct {
	End    int    `json:"end"`
	Period int    `json:"period"`
	Start  int    `json:"start"`
	Symbol string `json:"symbol"`
	Ticks  int    `json:"ticks"`
}

type GetChartLastInfo struct {
	Period int    `json:"period"`
	Start  int    `json:"start"`
	Symbol string `json:"symbol"`
}

type SymbolArguments struct {
	Symbol string `json:"symbol"`
}

type TickPricesArguments struct {
	Symbols   []string `json:"symbols"`
	Timestamp int      `json:"timestamp"`
	Level     int      `json:"level"`
}

type TradingHoursArguments struct {
	Symbols []string `json:"symbols"`
}

type TimeRangeArguments struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type GetTradesArguments struct {
	OpenedOnly bool `json:"openedOnly"`
}

type GetTradeRecordsArguments struct {
	Orders []int `json:"orders"`
}

type SymbolVolumeArguments struct {
	Symbol string  `json:"symbol"`
	Volume float64 `json:"volume"`
}

type ProfitCalculationArguments struct {
	Symbol     string  `json:"symbol"`
	Cmd        int     `json:"cmd"`
	OpenPrice  float64 `json:"openPrice"`
	ClosePrice float64 `json:"closePrice"`
	Volume     float64 `json:"volume"`
}
