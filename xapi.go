package xapi

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/MateoGreil/xapi-go/internal/protocols/socket"
	"github.com/MateoGreil/xapi-go/internal/protocols/stream"
	"github.com/gorilla/websocket"
)

type client struct {
	conn                 *websocket.Conn
	streamConn           *websocket.Conn
	streamSessionId      string
	streamMessageChannel chan interface{}
	CandlesChannel       chan stream.Candle
	TickPricesChannel    chan stream.Tick
	BalanceChannel       chan stream.Balance
	TradesChannel        chan stream.Trade
	NewsChannel          chan stream.News
	mutexSendMessage     sync.Mutex
}

const (
	websocketBaseURL = "wss://ws.xapi.pro"
	pingInterval     = 5 * time.Minute
)

func NewClient(userId string, password string, connectionType string) (*client, error) {
	var websocketURL string
	var websocketStreamURL string

	switch connectionType {
	case "demo":
		websocketURL = websocketBaseURL + "/demo"
		websocketStreamURL = websocketBaseURL + "/demoStream"
	case "real":
		websocketURL = websocketBaseURL + "/real"
		websocketStreamURL = websocketBaseURL + "/realStream"
	}

	conn, _, err := websocket.DefaultDialer.Dial(websocketURL, nil)
	if err != nil {
		return nil, err
	}

	streamSessionId, err := login(conn, userId, password)
	if err != nil {
		return nil, err
	}

	streamConn, _, err := websocket.DefaultDialer.Dial(websocketStreamURL, nil)
	if err != nil {
		return nil, err
	}
	getKeepAlive(streamConn, streamSessionId)

	c := &client{
		conn:                 conn,
		streamConn:           streamConn,
		streamSessionId:      streamSessionId,
		streamMessageChannel: make(chan interface{}),
		CandlesChannel:       make(chan stream.Candle),
		TickPricesChannel:    make(chan stream.Tick),
		BalanceChannel:       make(chan stream.Balance),
		TradesChannel:        make(chan stream.Trade),
		NewsChannel:          make(chan stream.News),
		mutexSendMessage:     sync.Mutex{},
	}
	go c.pingSocket()
	go c.pingStream()
	go c.listenStream()
	go c.streamWriteJSON()

	return c, nil
}

func (c *client) listenStream() {
	for {
		_, message, err := c.streamConn.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		response := stream.Response{}
		if err = json.Unmarshal(message, &response); err != nil {
			fmt.Printf("stream unmarshal error: %s\n", err.Error())
			continue
		}
		switch response.Command {
		case "candle":
			var r stream.ResponseCandle
			if err = json.Unmarshal(message, &r); err == nil {
				c.CandlesChannel <- r.Data
			}
		case "tickPrices":
			var r stream.ResponseTick
			if err = json.Unmarshal(message, &r); err == nil {
				c.TickPricesChannel <- r.Data
			}
		case "balance":
			var r stream.ResponseBalance
			if err = json.Unmarshal(message, &r); err == nil {
				c.BalanceChannel <- r.Data
			}
		case "trade":
			var r stream.ResponseTrade
			if err = json.Unmarshal(message, &r); err == nil {
				c.TradesChannel <- r.Data
			}
		case "news":
			var r stream.ResponseNews
			if err = json.Unmarshal(message, &r); err == nil {
				c.NewsChannel <- r.Data
			}
		case "keepAlive":
			// heartbeat, no action needed
		default:
			fmt.Printf("unknown stream message: %s\n", message)
		}
	}
}

func (c *client) pingSocket() {
	for {
		request := struct {
			Command string `json:"command"`
		}{Command: "ping"}
		if err := c.sendSocketCommand(request, nil); err != nil {
			fmt.Printf("ping socket failed: %s\n", err.Error())
		}
		time.Sleep(pingInterval)
	}
}

func (c *client) pingStream() {
	for {
		request := stream.Request{
			Command:         "ping",
			StreamSessionId: c.streamSessionId,
		}
		c.streamMessageChannel <- request
		time.Sleep(pingInterval)
	}
}

func (c *client) streamWriteJSON() {
	for {
		message := <-c.streamMessageChannel
		c.streamConn.WriteJSON(message)
		fmt.Printf("messageStream: %+v\n", message)
	}
}

func (c *client) sendSocketCommand(request interface{}, result interface{}) error {
	c.mutexSendMessage.Lock()
	defer c.mutexSendMessage.Unlock()
	if err := c.conn.WriteJSON(request); err != nil {
		return err
	}
	response := socket.Response{}
	if err := c.conn.ReadJSON(&response); err != nil {
		return err
	}
	if !response.Status {
		return fmt.Errorf("%s: %s", response.ErrorCode, response.ErrorDescr)
	}
	if result != nil {
		return json.Unmarshal(response.ReturnData, result)
	}
	return nil
}

func (c *client) Logout() error {
	request := struct {
		Command string `json:"command"`
	}{Command: "logout"}
	return c.sendSocketCommand(request, nil)
}

func login(conn *websocket.Conn, userId string, password string) (string, error) {
	request := socket.Request{
		Command: "login",
		Arguments: socket.LoginRequestArguments{
			UserId:   userId,
			Password: password,
		},
	}

	conn.WriteJSON(request)
	response := socket.LoginResponse{}
	err := conn.ReadJSON(&response)
	if err != nil {
		return "", err
	}
	if response.Status == false {
		return "", fmt.Errorf("%+v", response.ErrorDescr)
	}
	return response.StreamSessionId, nil
}

func getKeepAlive(conn *websocket.Conn, streamSessionId string) {
	keepAliveReq := stream.Request{
		Command:         "getKeepAlive",
		StreamSessionId: streamSessionId,
	}
	conn.WriteJSON(keepAliveReq)
	_, message, err := conn.ReadMessage()
	if err != nil {
		// TODO: Handle errors
		fmt.Println(err.Error())
	}
	response := stream.KeepAliveResponse{}
	err = json.Unmarshal(message, &response)
	if err != nil {
		// TODO: Handle errors
		fmt.Println(err.Error())
	}
	if response.Command != "keepAlive" {
		// TODO: Handle errors
		fmt.Println(err.Error())
	}
}
