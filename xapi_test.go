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
