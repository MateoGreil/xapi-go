package xapi

import (
	"github.com/MateoGreil/xapi-go/internal/protocols/socket"
)

func (c *client) GetCommissionDef(symbol string, volume float64) (socket.Commission, error) {
	request := socket.Request{
		Command:   "getCommissionDef",
		Arguments: socket.SymbolVolumeArguments{Symbol: symbol, Volume: volume},
	}
	var result socket.Commission
	if err := c.sendSocketCommand(request, &result); err != nil {
		return socket.Commission{}, err
	}
	return result, nil
}

func (c *client) GetMarginTrade(symbol string, volume float64) (float64, error) {
	request := socket.Request{
		Command:   "getMarginTrade",
		Arguments: socket.SymbolVolumeArguments{Symbol: symbol, Volume: volume},
	}
	var result struct {
		Margin float64 `json:"margin"`
	}
	if err := c.sendSocketCommand(request, &result); err != nil {
		return 0, err
	}
	return result.Margin, nil
}

func (c *client) GetProfitCalculation(symbol string, cmd int, openPrice float64, closePrice float64, volume float64) (float64, error) {
	request := socket.Request{
		Command: "getProfitCalculation",
		Arguments: socket.ProfitCalculationArguments{
			Symbol:     symbol,
			Cmd:        cmd,
			OpenPrice:  openPrice,
			ClosePrice: closePrice,
			Volume:     volume,
		},
	}
	var result socket.Profit
	if err := c.sendSocketCommand(request, &result); err != nil {
		return 0, err
	}
	return result.Profit, nil
}
