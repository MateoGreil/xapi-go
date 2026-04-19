package xapi

import (
	"github.com/MateoGreil/xapi-go/internal/protocols/socket"
	"github.com/MateoGreil/xapi-go/internal/protocols/stream"
)

func (c *client) GetCandles(end int, period int, start int, symbol string, ticks int) ([]socket.Candle, error) {
	request := socket.Request{
		Command: "getChartRangeRequest",
		Arguments: socket.InfoArguments{
			Info: socket.GetCandlesInfo{
				End:    end,
				Period: period,
				Start:  start,
				Symbol: symbol,
				Ticks:  ticks,
			},
		},
	}
	var result socket.ChartResponse
	if err := c.sendSocketCommand(request, &result); err != nil {
		return nil, err
	}
	return result.RateInfos, nil
}

func (c *client) GetChartLast(period int, start int, symbol string) ([]socket.Candle, error) {
	request := socket.Request{
		Command: "getChartLastRequest",
		Arguments: socket.InfoArguments{
			Info: socket.GetChartLastInfo{
				Period: period,
				Start:  start,
				Symbol: symbol,
			},
		},
	}
	var result socket.ChartResponse
	if err := c.sendSocketCommand(request, &result); err != nil {
		return nil, err
	}
	return result.RateInfos, nil
}

func (c *client) SubscribeCandles(symbol string) {
	c.streamMessageChannel <- stream.GetCandlesRequest{
		Command:         "getCandles",
		StreamSessionId: c.streamSessionId,
		Symbol:          symbol,
	}
}

func (c *client) StopCandles(symbol string) {
	c.streamMessageChannel <- stream.StopCandlesRequest{
		Command:         "stopCandles",
		StreamSessionId: c.streamSessionId,
		Symbol:          symbol,
	}
}
