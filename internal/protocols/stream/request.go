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
