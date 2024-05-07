package socket

type LoginResponse struct {
	Status          bool   `json:"status"`
	StreamSessionId string `json:"streamSessionId"`
	ErrorCode       string `json:"errorCode"`
	ErrorDescr      string `json:"errorDescr"`
}
