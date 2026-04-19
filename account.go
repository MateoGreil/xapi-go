package xapi

import (
	"github.com/MateoGreil/xapi-go/internal/protocols/socket"
	"github.com/MateoGreil/xapi-go/internal/protocols/stream"
)

func (c *client) GetCurrentUserData() (socket.UserData, error) {
	request := struct {
		Command string `json:"command"`
	}{Command: "getCurrentUserData"}
	var result socket.UserData
	if err := c.sendSocketCommand(request, &result); err != nil {
		return socket.UserData{}, err
	}
	return result, nil
}

func (c *client) GetMarginLevel() (socket.MarginLevel, error) {
	request := struct {
		Command string `json:"command"`
	}{Command: "getMarginLevel"}
	var result socket.MarginLevel
	if err := c.sendSocketCommand(request, &result); err != nil {
		return socket.MarginLevel{}, err
	}
	return result, nil
}

func (c *client) SubscribeBalance() {
	c.streamMessageChannel <- stream.GetBalanceRequest{
		Command:         "getBalance",
		StreamSessionId: c.streamSessionId,
	}
}

func (c *client) StopBalance() {
	c.streamMessageChannel <- stream.Request{
		Command:         "stopBalance",
		StreamSessionId: c.streamSessionId,
	}
}
