package stream

type Response struct {
	Command string `json:"command"`
}

type KeepAliveResponse struct {
	Command string        `json:"command"`
	Data    KeepAliveData `json:"data"`
}

type ResponseCandle struct {
	Command string `json:"command"`
	Data    Candle `json:"data"`
}

type ResponseTick struct {
	Command string `json:"command"`
	Data    Tick   `json:"data"`
}

type ResponseBalance struct {
	Command string  `json:"command"`
	Data    Balance `json:"data"`
}

type ResponseTrade struct {
	Command string `json:"command"`
	Data    Trade  `json:"data"`
}

type ResponseNews struct {
	Command string `json:"command"`
	Data    News   `json:"data"`
}
