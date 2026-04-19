package xapi

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/MateoGreil/xapi-go/internal/protocols/socket"
)

// --- Integration guard ---

func skipIfNoCredentials(t *testing.T) {
	t.Helper()
	if os.Getenv("XAPI_USER_ID") == "" || os.Getenv("XAPI_PASSWORD") == "" {
		t.Skip("integration test: set XAPI_USER_ID and XAPI_PASSWORD to run")
	}
}

func integrationClient(t *testing.T) *client {
	t.Helper()
	skipIfNoCredentials(t)
	c, err := NewClient(os.Getenv("XAPI_USER_ID"), os.Getenv("XAPI_PASSWORD"), "demo")
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	return c
}

// --- Unit tests (mock server, always run) ---

func TestNewClient(t *testing.T) {
	c, _ := newTestClient(t)
	if c == nil {
		t.Error("expected non-nil client")
	}
}

func TestSubscribeCandles(t *testing.T) {
	c, _ := newTestClient(t)
	c.SubscribeCandles("EURUSD")
	select {
	case candle := <-c.CandlesChannel:
		if candle.Symbol != "EURUSD" {
			t.Errorf("expected EURUSD candle, got symbol %q", candle.Symbol)
		}
	case <-time.After(2 * time.Second):
		t.Error("timed out waiting for candle")
	}
}

func TestSubscribeTickPrices(t *testing.T) {
	c, _ := newTestClient(t)
	c.SubscribeTickPrices("EURUSD")
	select {
	case tick := <-c.TickPricesChannel:
		if tick.Symbol != "EURUSD" {
			t.Errorf("expected EURUSD tick, got symbol %q", tick.Symbol)
		}
	case <-time.After(2 * time.Second):
		t.Error("timed out waiting for tick")
	}
}

func TestSubscribeBalance(t *testing.T) {
	c, _ := newTestClient(t)
	c.SubscribeBalance()
	select {
	case balance := <-c.BalanceChannel:
		if balance.Balance <= 0 {
			t.Errorf("expected positive balance, got %f", balance.Balance)
		}
	case <-time.After(2 * time.Second):
		t.Error("timed out waiting for balance")
	}
}

func TestSubscribeTrades(t *testing.T) {
	c, _ := newTestClient(t)
	c.SubscribeTrades()
	select {
	case trade := <-c.TradesChannel:
		if trade.Order == 0 {
			t.Error("expected non-zero order")
		}
	case <-time.After(2 * time.Second):
		t.Error("timed out waiting for trade")
	}
}

func TestSubscribeNews(t *testing.T) {
	c, _ := newTestClient(t)
	c.SubscribeNews()
	select {
	case news := <-c.NewsChannel:
		if news.Title == "" {
			t.Error("expected non-empty news title")
		}
	case <-time.After(2 * time.Second):
		t.Error("timed out waiting for news")
	}
}

func TestGetChartLast(t *testing.T) {
	c, _ := newTestClient(t)
	start := int(time.Now().Add(-24 * time.Hour).UnixMilli())
	candles, err := c.GetChartLast(1, start, "EURUSD")
	if err != nil {
		t.Error(err)
	}
	if len(candles) == 0 {
		t.Error("expected at least one candle")
	}
}

func TestGetCandles(t *testing.T) {
	c, _ := newTestClient(t)
	start := int(time.Now().Add(-24 * time.Hour).UnixMilli())
	end := int(time.Now().UnixMilli())

	candles, err := c.GetCandles(end, 1, start, "EURUSD", 1)
	if err != nil {
		t.Error(err)
	}
	if len(candles) != 1 {
		t.Errorf("expected 1 candle, got %d", len(candles))
	}

	candles, err = c.GetCandles(end, 1, start, "EURUSD", 50)
	if err != nil {
		t.Error(err)
	}
	if len(candles) != 50 {
		t.Errorf("expected 50 candles, got %d", len(candles))
	}
}

func TestGetAllSymbols(t *testing.T) {
	c, _ := newTestClient(t)
	symbols, err := c.GetAllSymbols()
	if err != nil {
		t.Error(err)
	}
	if len(symbols) == 0 {
		t.Error("expected at least one symbol")
	}
}

func TestGetSymbol(t *testing.T) {
	c, _ := newTestClient(t)
	symbol, err := c.GetSymbol("EURUSD")
	if err != nil {
		t.Error(err)
	}
	if symbol.Symbol != "EURUSD" {
		t.Errorf("expected EURUSD, got %s", symbol.Symbol)
	}
}

func TestGetCalendar(t *testing.T) {
	c, _ := newTestClient(t)
	_, err := c.GetCalendar()
	if err != nil {
		t.Error(err)
	}
}

func TestGetTickPrices(t *testing.T) {
	c, _ := newTestClient(t)
	ticks, err := c.GetTickPrices([]string{"EURUSD"}, 0, 0)
	if err != nil {
		t.Error(err)
	}
	if len(ticks) == 0 {
		t.Error("expected at least one tick")
	}
}

func TestGetTradingHours(t *testing.T) {
	c, _ := newTestClient(t)
	hours, err := c.GetTradingHours([]string{"EURUSD"})
	if err != nil {
		t.Error(err)
	}
	if len(hours) == 0 {
		t.Error("expected at least one trading hours record")
	}
}

func TestGetNews(t *testing.T) {
	c, _ := newTestClient(t)
	start := int(time.Now().Add(-7 * 24 * time.Hour).UnixMilli())
	_, err := c.GetNews(start, 0)
	if err != nil {
		t.Error(err)
	}
}

func TestGetIbsHistory(t *testing.T) {
	c, _ := newTestClient(t)
	start := int(time.Now().Add(-30 * 24 * time.Hour).UnixMilli())
	_, err := c.GetIbsHistory(start, 0)
	if err != nil {
		t.Error(err)
	}
}

func TestGetCurrentUserData(t *testing.T) {
	c, _ := newTestClient(t)
	data, err := c.GetCurrentUserData()
	if err != nil {
		t.Error(err)
	}
	if data.Currency == "" {
		t.Error("expected non-empty currency")
	}
}

func TestGetMarginLevel(t *testing.T) {
	c, _ := newTestClient(t)
	level, err := c.GetMarginLevel()
	if err != nil {
		t.Error(err)
	}
	if level.Currency == "" {
		t.Error("expected non-empty currency")
	}
}

func TestGetServerTime(t *testing.T) {
	c, _ := newTestClient(t)
	st, err := c.GetServerTime()
	if err != nil {
		t.Error(err)
	}
	if st.Time == 0 {
		t.Error("expected non-zero server time")
	}
}

func TestGetStepRules(t *testing.T) {
	c, _ := newTestClient(t)
	rules, err := c.GetStepRules()
	if err != nil {
		t.Error(err)
	}
	if len(rules) == 0 {
		t.Error("expected at least one step rule")
	}
}

func TestGetVersion(t *testing.T) {
	c, _ := newTestClient(t)
	v, err := c.GetVersion()
	if err != nil {
		t.Error(err)
	}
	if v == "" {
		t.Error("expected non-empty version")
	}
}

func TestGetCommissionDef(t *testing.T) {
	c, _ := newTestClient(t)
	_, err := c.GetCommissionDef("EURUSD", 1.0)
	if err != nil {
		t.Error(err)
	}
}

func TestGetMarginTrade(t *testing.T) {
	c, _ := newTestClient(t)
	margin, err := c.GetMarginTrade("EURUSD", 1.0)
	if err != nil {
		t.Error(err)
	}
	if margin <= 0 {
		t.Errorf("expected positive margin, got %f", margin)
	}
}

func TestGetProfitCalculation(t *testing.T) {
	c, _ := newTestClient(t)
	_, err := c.GetProfitCalculation("EURUSD", 0, 1.2000, 1.2100, 1.0)
	if err != nil {
		t.Error(err)
	}
}

func TestGetTrades(t *testing.T) {
	c, _ := newTestClient(t)
	_, err := c.GetTrades(true)
	if err != nil {
		t.Error(err)
	}
}

func TestGetTradesHistory(t *testing.T) {
	c, _ := newTestClient(t)
	start := int(time.Now().Add(-30 * 24 * time.Hour).UnixMilli())
	_, err := c.GetTradesHistory(start, 0)
	if err != nil {
		t.Error(err)
	}
}

func TestTradeTransaction(t *testing.T) {
	c, _ := newTestClient(t)
	orderNum, err := c.TradeTransaction(socket.TradeTransInfo{
		Cmd:    0,
		Symbol: "EURUSD",
		Volume: 0.01,
		Price:  1.2,
		Type:   0,
	})
	if err != nil {
		t.Error(err)
		return
	}
	if orderNum == 0 {
		t.Error("expected non-zero order number")
	}
	status, err := c.TradeTransactionStatus(orderNum)
	if err != nil {
		t.Error(err)
	}
	if status.RequestStatus == 0 {
		t.Error("expected non-zero request status")
	}
}

func TestLogout(t *testing.T) {
	c, _ := newTestClient(t)
	if err := c.Logout(); err != nil {
		t.Error(err)
	}
}

// --- Integration tests (require XAPI_USER_ID + XAPI_PASSWORD) ---

func TestIntegrationNewClientBadCredentials(t *testing.T) {
	skipIfNoCredentials(t)

	_, err := NewClient(os.Getenv("XAPI_USER_ID"), "wrong-password", "demo")
	if err == nil || err.Error() != "userPasswordCheck: Invalid login or password" {
		t.Errorf("unexpected error: %v", err)
	}

	_, err = NewClient("wrong-user-id", os.Getenv("XAPI_PASSWORD"), "demo")
	if err == nil || err.Error() != "Invalid parameters" {
		fmt.Println(err)
		t.Errorf("unexpected error: %v", err)
	}
}

func TestIntegrationGetCandles(t *testing.T) {
	c := integrationClient(t)
	start := int(time.Now().Add(-24 * time.Hour).UnixMilli())
	end := int(time.Now().UnixMilli())

	candles, err := c.GetCandles(end, 1, start, "EURUSD", 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(candles) != 1 {
		t.Errorf("expected 1 candle, got %d", len(candles))
	}

	candles, err = c.GetCandles(end, 1, start, "EURUSD", 50)
	if err != nil {
		t.Fatal(err)
	}
	if len(candles) != 50 {
		t.Errorf("expected 50 candles, got %d", len(candles))
	}
}

func TestIntegrationSubscribeCandles(t *testing.T) {
	c := integrationClient(t)
	c.SubscribeCandles("EURUSD")
	select {
	case candle := <-c.CandlesChannel:
		fmt.Printf("%+v\n", candle)
	case <-time.After(2 * time.Minute):
		t.Error("timed out waiting for candle")
	}
}
