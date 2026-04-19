package socket

import "encoding/json"

type LoginResponse struct {
	Status          bool   `json:"status"`
	StreamSessionId string `json:"streamSessionId"`
	ErrorCode       string `json:"errorCode"`
	ErrorDescr      string `json:"errorDescr"`
}

type Response struct {
	Status     bool            `json:"status"`
	ErrorCode  string          `json:"errorCode"`
	ErrorDescr string          `json:"errorDescr"`
	ReturnData json.RawMessage `json:"returnData"`
}
