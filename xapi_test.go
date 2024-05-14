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
