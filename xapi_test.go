package xapi

import (
	"fmt"
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	xapiClient, err := NewClient(os.Getenv("XAPI_USER_ID"), os.Getenv("XAPI_PASSWORD"), "demo")
	if err != nil {
		t.Error(err)
	}

	_, err = NewClient(os.Getenv("XAPI_USER_ID"), "wrong-password", "demo")
	if err.Error() != "userPasswordCheck: Invalid login or password" {
		t.Error(err)
	}

	_, err = NewClient("wrong-user-id", os.Getenv("XAPI_PASSWORD"), "demo")
	if err.Error() != "Invalid parameters" {
		fmt.Println(err)
		t.Error(err)
	}
}
