package stream

type Request struct {
	Command         string `json:"command"`
	StreamSessionId string `json:"streamSessionId"`
}
