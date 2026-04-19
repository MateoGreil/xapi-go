package xapi

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient(os.Getenv("XAPI_USER_ID"), "wrong-password", "demo")
	if err.Error() != "userPasswordCheck: Invalid login or password" {
		t.Error(err)
	}

	_, err = NewClient("wrong-user-id", os.Getenv("XAPI_PASSWORD"), "demo")
	if err.Error() != "Invalid parameters" {
		fmt.Println(err)
		t.Error(err)
	}

	_, err = NewClient(os.Getenv("XAPI_USER_ID"), os.Getenv("XAPI_PASSWORD"), "demo")
	if err != nil {
		t.Error(err)
	}
}

func TestSuscribeCandles(t *testing.T) {
	xapiClient, err := NewClient(os.Getenv("XAPI_USER_ID"), os.Getenv("XAPI_PASSWORD"), "demo")
	if err != nil {
		t.Error(err)
	}

	xapiClient.SubscribeCandles("EURUSD")
	select {
	case candle := <-xapiClient.CandlesChannel:
		fmt.Printf("%+v\n", candle)
	case <-time.After(2 * time.Minute):
		t.Error("Did not receive candles")
	}
}

func TestGetChartLast(t *testing.T) {
	xapiClient, err := NewClient(os.Getenv("XAPI_USER_ID"), os.Getenv("XAPI_PASSWORD"), "demo")
	if err != nil {
		t.Fatal(err)
	}
	start := int(time.Now().Add(-24 * time.Hour).UnixMilli())
	candles, err := xapiClient.GetChartLast(1, start, "EURUSD")
	if err != nil {
		t.Error(err)
	}
	if len(candles) == 0 {
		t.Error("expected at least one candle")
	}
}

func TestGetCandles(t *testing.T) {
	xapiClient, err := NewClient(os.Getenv("XAPI_USER_ID"), os.Getenv("XAPI_PASSWORD"), "demo")
	if err != nil {
		t.Error(err)
	}

	start := int(time.Now().Add(-24 * 1 * time.Hour).UnixMilli())
	period := 1
	ticks := 1
	end := int(time.Now().Add(-24 * 1 * time.Hour).UnixMilli())
	candles, err := xapiClient.GetCandles(start, period, end, "EURUSD", ticks)
	if err != nil {
		t.Error(err)
	} else {
		length := len(candles)
		if length != 1 {
			t.Errorf("Should contain 1 candle, but contains %d", length)
		}
	}

	ticks = 50
	candles, err = xapiClient.GetCandles(start, period, end, "EURUSD", ticks)
	if err != nil {
		t.Error(err)
	} else {
		length := len(candles)
		if length != 50 {
			t.Errorf("Should contain 50 candle, but contains %d", length)
		}
	}

}

func TestGetAllSymbols(t *testing.T) {
	xapiClient, err := NewClient(os.Getenv("XAPI_USER_ID"), os.Getenv("XAPI_PASSWORD"), "demo")
	if err != nil {
		t.Fatal(err)
	}
	symbols, err := xapiClient.GetAllSymbols()
	if err != nil {
		t.Error(err)
	}
	if len(symbols) == 0 {
		t.Error("expected at least one symbol")
	}
}

func TestGetSymbol(t *testing.T) {
	xapiClient, err := NewClient(os.Getenv("XAPI_USER_ID"), os.Getenv("XAPI_PASSWORD"), "demo")
	if err != nil {
		t.Fatal(err)
	}
	symbol, err := xapiClient.GetSymbol("EURUSD")
	if err != nil {
		t.Error(err)
	}
	if symbol.Symbol != "EURUSD" {
		t.Errorf("expected EURUSD, got %s", symbol.Symbol)
	}
}

func TestGetCalendar(t *testing.T) {
	xapiClient, err := NewClient(os.Getenv("XAPI_USER_ID"), os.Getenv("XAPI_PASSWORD"), "demo")
	if err != nil {
		t.Fatal(err)
	}
	_, err = xapiClient.GetCalendar()
	if err != nil {
		t.Error(err)
	}
}

func TestGetTickPrices(t *testing.T) {
	xapiClient, err := NewClient(os.Getenv("XAPI_USER_ID"), os.Getenv("XAPI_PASSWORD"), "demo")
	if err != nil {
		t.Fatal(err)
	}
	ticks, err := xapiClient.GetTickPrices([]string{"EURUSD"}, 0, 0)
	if err != nil {
		t.Error(err)
	}
	if len(ticks) == 0 {
		t.Error("expected at least one tick")
	}
}

func TestGetTradingHours(t *testing.T) {
	xapiClient, err := NewClient(os.Getenv("XAPI_USER_ID"), os.Getenv("XAPI_PASSWORD"), "demo")
	if err != nil {
		t.Fatal(err)
	}
	hours, err := xapiClient.GetTradingHours([]string{"EURUSD"})
	if err != nil {
		t.Error(err)
	}
	if len(hours) == 0 {
		t.Error("expected at least one trading hours record")
	}
}

func TestGetNews(t *testing.T) {
	xapiClient, err := NewClient(os.Getenv("XAPI_USER_ID"), os.Getenv("XAPI_PASSWORD"), "demo")
	if err != nil {
		t.Fatal(err)
	}
	start := int(time.Now().Add(-7 * 24 * time.Hour).UnixMilli())
	_, err = xapiClient.GetNews(start, 0)
	if err != nil {
		t.Error(err)
	}
}

func TestGetIbsHistory(t *testing.T) {
	xapiClient, err := NewClient(os.Getenv("XAPI_USER_ID"), os.Getenv("XAPI_PASSWORD"), "demo")
	if err != nil {
		t.Fatal(err)
	}
	start := int(time.Now().Add(-30 * 24 * time.Hour).UnixMilli())
	_, err = xapiClient.GetIbsHistory(start, 0)
	if err != nil {
		t.Error(err)
	}
}
