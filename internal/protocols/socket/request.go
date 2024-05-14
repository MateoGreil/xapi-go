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
