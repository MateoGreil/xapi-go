package xapi

import (
	"github.com/MateoGreil/xapi-go/internal/protocols/socket"
	"github.com/MateoGreil/xapi-go/internal/protocols/stream"
)

func (c *client) GetTrades(openedOnly bool) ([]socket.TradeRecord, error) {
	request := socket.Request{
		Command:   "getTrades",
		Arguments: socket.GetTradesArguments{OpenedOnly: openedOnly},
	}
	var result []socket.TradeRecord
	if err := c.sendSocketCommand(request, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetTradeRecords(orders []int) ([]socket.TradeRecord, error) {
	request := socket.Request{
		Command:   "getTradeRecords",
		Arguments: socket.GetTradeRecordsArguments{Orders: orders},
	}
	var result []socket.TradeRecord
	if err := c.sendSocketCommand(request, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetTradesHistory(start int, end int) ([]socket.TradeRecord, error) {
	request := socket.Request{
		Command:   "getTradesHistory",
		Arguments: socket.TimeRangeArguments{Start: start, End: end},
	}
	var result []socket.TradeRecord
	if err := c.sendSocketCommand(request, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) SubscribeTrades() {
	c.streamMessageChannel <- stream.GetTradesRequest{
		Command:         "getTrades",
		StreamSessionId: c.streamSessionId,
	}
}

func (c *client) StopTrades() {
	c.streamMessageChannel <- stream.Request{
		Command:         "stopTrades",
		StreamSessionId: c.streamSessionId,
	}
}

func (c *client) TradeTransaction(info socket.TradeTransInfo) (int, error) {
	request := socket.Request{
		Command:   "tradeTransaction",
		Arguments: socket.TradeTransactionArguments{TradeTransInfo: info},
	}
	var result socket.TradeTransactionResponse
	if err := c.sendSocketCommand(request, &result); err != nil {
		return 0, err
	}
	return result.Order, nil
}

func (c *client) TradeTransactionStatus(order int) (socket.TradeTransactionStatusResponse, error) {
	request := socket.Request{
		Command:   "tradeTransactionStatus",
		Arguments: socket.TradeTransactionStatusArguments{Order: order},
	}
	var result socket.TradeTransactionStatusResponse
	if err := c.sendSocketCommand(request, &result); err != nil {
		return socket.TradeTransactionStatusResponse{}, err
	}
	return result, nil
}
