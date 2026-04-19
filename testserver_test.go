package xapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type mockServer struct {
	server     *httptest.Server
	streamConn *websocket.Conn
	streamMu   sync.Mutex
}

func newMockServer(t *testing.T) *mockServer {
	t.Helper()
	ms := &mockServer{}
	mux := http.NewServeMux()
	mux.HandleFunc("/demo", ms.handleMain)
	mux.HandleFunc("/demoStream", ms.handleStream)
	ms.server = httptest.NewServer(mux)
	t.Cleanup(ms.server.Close)
	return ms
}

func (ms *mockServer) mainURL() string {
	return "ws" + strings.TrimPrefix(ms.server.URL, "http") + "/demo"
}

func (ms *mockServer) streamURL() string {
	return "ws" + strings.TrimPrefix(ms.server.URL, "http") + "/demoStream"
}

func (ms *mockServer) handleMain(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			return
		}
		var req struct {
			Command   string          `json:"command"`
			Arguments json.RawMessage `json:"arguments"`
		}
		if err := json.Unmarshal(msgBytes, &req); err != nil {
			return
		}

		resp := ms.responseFor(req.Command, req.Arguments)
		if err := conn.WriteMessage(websocket.TextMessage, []byte(resp)); err != nil {
			return
		}
	}
}

func (ms *mockServer) responseFor(command string, arguments json.RawMessage) string {
	switch command {
	case "login":
		return `{"status":true,"streamSessionId":"test-stream-id"}`
	case "ping", "logout":
		return `{"status":true}`
	case "getChartRangeRequest":
		var args struct {
			Info struct {
				Ticks int `json:"ticks"`
			} `json:"info"`
		}
		json.Unmarshal(arguments, &args)
		n := args.Info.Ticks
		if n <= 0 {
			n = 1
		}
		return buildChartResponse(n)
	case "getChartLastRequest":
		return buildChartResponse(1)
	case "getAllSymbols":
		return `{"status":true,"returnData":[{"symbol":"EURUSD","ask":1.2,"bid":1.19,"description":"Euro vs US Dollar","currency":"USD","categoryName":"FX","contractSize":100000,"currencyPair":true,"currencyProfit":"USD","high":1.21,"low":1.18,"leverage":30,"lotMax":50,"lotMin":0.01,"lotStep":0.01,"percentage":100,"pipsPrecision":4,"precision":5,"quoteId":1,"spreadRaw":0.0001,"spreadTable":0.0001,"stopsLevel":0,"swapEnable":true,"swapLong":-0.5,"swapShort":-0.5,"swapType":0,"time":1000,"timeString":"2024-01-01","type":21}]}`
	case "getSymbol":
		return `{"status":true,"returnData":{"symbol":"EURUSD","ask":1.2,"bid":1.19,"description":"Euro vs US Dollar","currency":"USD","categoryName":"FX","contractSize":100000,"currencyPair":true,"currencyProfit":"USD","high":1.21,"low":1.18,"leverage":30,"lotMax":50,"lotMin":0.01,"lotStep":0.01,"percentage":100,"pipsPrecision":4,"precision":5,"quoteId":1,"spreadRaw":0.0001,"spreadTable":0.0001,"stopsLevel":0,"swapEnable":true,"swapLong":-0.5,"swapShort":-0.5,"swapType":0,"time":1000,"timeString":"2024-01-01","type":21}}`
	case "getCalendar":
		return `{"status":true,"returnData":[]}`
	case "getTickPrices":
		return `{"status":true,"returnData":{"quotations":[{"symbol":"EURUSD","ask":1.2,"bid":1.19,"high":1.21,"low":1.18,"timestamp":1000,"level":0,"spreadRaw":0.0001,"spreadTable":0.0001}]}}`
	case "getTradingHours":
		return `{"status":true,"returnData":[{"symbol":"EURUSD","quotes":[{"day":1,"fromT":0,"toT":86400000}],"trading":[{"day":1,"fromT":0,"toT":86400000}]}]}`
	case "getCurrentUserData":
		return `{"status":true,"returnData":{"companyUnit":8,"currency":"EUR","group":"demoPLN200","ibAccount":false,"leverage":100,"leverageMultiplier":0.25,"spreadType":"FLOAT","trailingStop":false}}`
	case "getMarginLevel":
		return `{"status":true,"returnData":{"balance":10000.0,"credit":0.0,"currency":"EUR","equity":10000.0,"margin":0.0,"margin_free":10000.0,"margin_level":0.0}}`
	case "getNews":
		return `{"status":true,"returnData":[]}`
	case "getIbsHistory":
		return `{"status":true,"returnData":[]}`
	case "getServerTime":
		return `{"status":true,"returnData":{"time":1704067200000,"timeString":"2024-01-01 00:00:00"}}`
	case "getStepRules":
		return `{"status":true,"returnData":[{"id":1,"name":"Forex","steps":[{"fromValue":0.1,"step":0.1}]}]}`
	case "getVersion":
		return `{"status":true,"returnData":{"version":"2.5.0"}}`
	case "getCommissionDef":
		return `{"status":true,"returnData":{"commission":0.0,"rateOfExchange":1.0}}`
	case "getMarginTrade":
		return `{"status":true,"returnData":{"margin":100.0}}`
	case "getProfitCalculation":
		return `{"status":true,"returnData":{"profit":100.0}}`
	case "getTrades":
		return `{"status":true,"returnData":[]}`
	case "getTradeRecords":
		return `{"status":true,"returnData":[]}`
	case "getTradesHistory":
		return `{"status":true,"returnData":[]}`
	case "tradeTransaction":
		return `{"status":true,"returnData":{"order":12345}}`
	case "tradeTransactionStatus":
		return `{"status":true,"returnData":{"ask":1.2,"bid":1.19,"customComment":"","message":"","order":12345,"requestStatus":3}}`
	default:
		return `{"status":true}`
	}
}

func buildChartResponse(n int) string {
	candles := make([]string, n)
	for i := range candles {
		candles[i] = `{"ctm":1000,"ctmString":"2024-01-01 00:00:00","open":1.2,"close":0.0001,"high":0.0002,"low":-0.0001,"vol":10.0}`
	}
	return fmt.Sprintf(`{"status":true,"returnData":{"digits":5,"rateInfos":[%s]}}`, strings.Join(candles, ","))
}

func (ms *mockServer) handleStream(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer func() {
		ms.streamMu.Lock()
		ms.streamConn = nil
		ms.streamMu.Unlock()
		conn.Close()
	}()
	ms.streamMu.Lock()
	ms.streamConn = conn
	ms.streamMu.Unlock()

	for {
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			return
		}
		var req struct {
			Command string `json:"command"`
		}
		if err := json.Unmarshal(msgBytes, &req); err != nil {
			return
		}
		switch req.Command {
		case "getKeepAlive":
			conn.WriteMessage(websocket.TextMessage, []byte(`{"command":"keepAlive","data":{"timestamp":0}}`))
		case "getCandles":
			conn.WriteMessage(websocket.TextMessage, []byte(`{"command":"candle","data":{"symbol":"EURUSD","open":1.2,"close":0.0001,"high":0.0002,"low":-0.0001,"vol":10.0,"ctm":1000,"ctmString":"2024-01-01 00:00:00","quoteId":1}}`))
		case "getTickPrices":
			conn.WriteMessage(websocket.TextMessage, []byte(`{"command":"tickPrices","data":{"symbol":"EURUSD","ask":1.2,"bid":1.19,"high":1.21,"low":1.18,"timestamp":1000,"level":0,"spreadRaw":0.0001,"spreadTable":0.0001,"quoteId":1}}`))
		case "getBalance":
			conn.WriteMessage(websocket.TextMessage, []byte(`{"command":"balance","data":{"balance":10000.0,"credit":0.0,"equity":10000.0,"margin":0.0,"marginFree":10000.0,"marginLevel":0.0}}`))
		case "getTrades":
			conn.WriteMessage(websocket.TextMessage, []byte(`{"command":"trade","data":{"order":12345,"order2":0,"position":12345,"symbol":"EURUSD","cmd":0,"volume":0.01,"open_price":1.2,"close_price":0.0,"sl":0.0,"tp":0.0,"profit":0.0,"storage":0.0,"margin_rate":0.0,"digits":5,"closed":false,"open_time":1000,"open_timeString":"2024-01-01","timestamp":1000}}`))
		case "getNews":
			conn.WriteMessage(websocket.TextMessage, []byte(`{"command":"news","data":{"body":"test news","bodylen":9,"key":"test-key","time":1000,"timeString":"2024-01-01","title":"Test News"}}`))
		}
		// ping, stopCandles, stopTickPrices, stopBalance, stopTrades, stopNews: no response needed
	}
}

func newTestClient(t *testing.T) (*client, *mockServer) {
	t.Helper()
	ms := newMockServer(t)
	c, err := newClientWithURLs("test-user", "test-password", ms.mainURL(), ms.streamURL())
	if err != nil {
		t.Fatalf("newTestClient: %v", err)
	}
	t.Cleanup(func() { c.Logout() })
	return c, ms
}
