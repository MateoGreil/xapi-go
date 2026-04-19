package xapi

import (
	"github.com/MateoGreil/xapi-go/internal/protocols/socket"
	"github.com/MateoGreil/xapi-go/internal/protocols/stream"
)

func (c *client) GetAllSymbols() ([]socket.Symbol, error) {
	request := struct {
		Command string `json:"command"`
	}{Command: "getAllSymbols"}
	var result []socket.Symbol
	if err := c.sendSocketCommand(request, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetSymbol(symbol string) (socket.Symbol, error) {
	request := socket.Request{
		Command:   "getSymbol",
		Arguments: socket.SymbolArguments{Symbol: symbol},
	}
	var result socket.Symbol
	if err := c.sendSocketCommand(request, &result); err != nil {
		return socket.Symbol{}, err
	}
	return result, nil
}

func (c *client) GetCalendar() ([]socket.CalendarRecord, error) {
	request := struct {
		Command string `json:"command"`
	}{Command: "getCalendar"}
	var result []socket.CalendarRecord
	if err := c.sendSocketCommand(request, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetTickPrices(symbols []string, timestamp int, level int) ([]socket.TickRecord, error) {
	request := socket.Request{
		Command: "getTickPrices",
		Arguments: socket.TickPricesArguments{
			Symbols:   symbols,
			Timestamp: timestamp,
			Level:     level,
		},
	}
	var result socket.TickPricesResponse
	if err := c.sendSocketCommand(request, &result); err != nil {
		return nil, err
	}
	return result.Quotations, nil
}

func (c *client) GetTradingHours(symbols []string) ([]socket.TradingHoursRecord, error) {
	request := socket.Request{
		Command:   "getTradingHours",
		Arguments: socket.TradingHoursArguments{Symbols: symbols},
	}
	var result []socket.TradingHoursRecord
	if err := c.sendSocketCommand(request, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetNews(start int, end int) ([]socket.NewsTopic, error) {
	request := socket.Request{
		Command:   "getNews",
		Arguments: socket.TimeRangeArguments{Start: start, End: end},
	}
	var result []socket.NewsTopic
	if err := c.sendSocketCommand(request, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetIbsHistory(start int, end int) ([]socket.IbRecord, error) {
	request := socket.Request{
		Command:   "getIbsHistory",
		Arguments: socket.TimeRangeArguments{Start: start, End: end},
	}
	var result []socket.IbRecord
	if err := c.sendSocketCommand(request, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) SubscribeTickPrices(symbol string) {
	c.streamMessageChannel <- stream.GetTickPricesRequest{
		Command:         "getTickPrices",
		StreamSessionId: c.streamSessionId,
		Symbol:          symbol,
	}
}

func (c *client) StopTickPrices(symbol string) {
	c.streamMessageChannel <- stream.StopTickPricesRequest{
		Command:         "stopTickPrices",
		StreamSessionId: c.streamSessionId,
		Symbol:          symbol,
	}
}

func (c *client) SubscribeNews() {
	c.streamMessageChannel <- stream.GetNewsRequest{
		Command:         "getNews",
		StreamSessionId: c.streamSessionId,
	}
}

func (c *client) StopNews() {
	c.streamMessageChannel <- stream.Request{
		Command:         "stopNews",
		StreamSessionId: c.streamSessionId,
	}
}
