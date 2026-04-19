package xapi

import (
	"github.com/MateoGreil/xapi-go/internal/protocols/socket"
)

func (c *client) GetServerTime() (socket.ServerTime, error) {
	request := struct {
		Command string `json:"command"`
	}{Command: "getServerTime"}
	var result socket.ServerTime
	if err := c.sendSocketCommand(request, &result); err != nil {
		return socket.ServerTime{}, err
	}
	return result, nil
}

func (c *client) GetStepRules() ([]socket.StepRule, error) {
	request := struct {
		Command string `json:"command"`
	}{Command: "getStepRules"}
	var result []socket.StepRule
	if err := c.sendSocketCommand(request, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetVersion() (string, error) {
	request := struct {
		Command string `json:"command"`
	}{Command: "getVersion"}
	var result socket.Version
	if err := c.sendSocketCommand(request, &result); err != nil {
		return "", err
	}
	return result.Version, nil
}
