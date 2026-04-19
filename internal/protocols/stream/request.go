package stream

type Request struct {
	Command         string `json:"command"`
	StreamSessionId string `json:"streamSessionId"`
}

type GetCandlesRequest struct {
	Command         string `json:"command"`
	StreamSessionId string `json:"streamSessionId"`
	Symbol          string `json:"symbol"`
}

type GetTickPricesRequest struct {
	Command         string `json:"command"`
	StreamSessionId string `json:"streamSessionId"`
	Symbol          string `json:"symbol"`
	MinArrivalTime  int    `json:"minArrivalTime,omitempty"`
	MaxLevel        int    `json:"maxLevel,omitempty"`
}

type GetBalanceRequest struct {
	Command         string `json:"command"`
	StreamSessionId string `json:"streamSessionId"`
}

type GetTradesRequest struct {
	Command         string `json:"command"`
	StreamSessionId string `json:"streamSessionId"`
}

type GetNewsRequest struct {
	Command         string `json:"command"`
	StreamSessionId string `json:"streamSessionId"`
}
