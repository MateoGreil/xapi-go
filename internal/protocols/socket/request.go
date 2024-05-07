package socket

type Request struct {
	Command   string           `json:"command"`
	Arguments RequestArguments `json:"arguments"`
}

type RequestArguments interface{}

type LoginRequestArguments struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}
