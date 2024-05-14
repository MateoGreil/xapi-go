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
	mutexSendMessage     sync.Mutex
}

const (
	websocketBaseURL = "wss://ws.xtb.com"
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

	var m sync.Mutex
	c := &client{
		conn:                 conn,
		streamConn:           streamConn,
		streamSessionId:      streamSessionId,
		streamMessageChannel: make(chan interface{}),
		CandlesChannel:       make(chan stream.Candle),
		mutexSendMessage:     m,
	}
	go c.pingSocket()
	go c.pingStream()
	go c.listenStream()
	go c.streamWriteJSON()

	return c, nil
}

func (c *client) SubscribeCandles(symbol string) {
	request := stream.GetCandlesRequest{
		Command:         "getCandles",
		StreamSessionId: c.streamSessionId,
		Symbol:          symbol,
	}
	c.streamMessageChannel <- request
}

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
	response := socket.Response{}
	c.mutexSendMessage.Lock()
	c.conn.WriteJSON(request)
	err := c.conn.ReadJSON(&response)
	c.mutexSendMessage.Unlock()
	if err != nil {
		return nil, err
	}
	if response.Status != true {
		return nil, fmt.Errorf("Error on sending getChartRangeRequest: %+v, response:, %+v", request, response)
	}
	return response.ReturnData.RateInfos, nil
}

func (c *client) listenStream() {
	for {
		_, message, err := c.streamConn.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
		}
		response := stream.Response{}
		err = json.Unmarshal(message, &response)
		if err != nil {
			fmt.Printf("message: %s\n", message)
			fmt.Println(err.Error())
		}
		switch response.Command {
		case "candle":
			responseCandle := stream.ResponseCandle{}
			err = json.Unmarshal(message, &responseCandle)
			if err != nil {
				fmt.Println(err.Error())
			}
			c.CandlesChannel <- responseCandle.Data
		case "keepAlive":
			fmt.Printf("keepAlive received\n")
		default:
			fmt.Printf("Unknown stream message: %s\n", message)
		}
	}
}

func (c *client) pingSocket() {
	for {
		request := struct {
			Command string `json:"command"`
		}{Command: "ping"}
		response := socket.Response{}
		c.mutexSendMessage.Lock()
		c.conn.WriteJSON(request)
		err := c.conn.ReadJSON(&response)
		c.mutexSendMessage.Unlock()
		if err != nil {
			//TODO: Handle error
			fmt.Printf("Ping socket failed: %s", err.Error())
		} else if response.Status != true {
			fmt.Errorf("Error on sending request: %+v, response:, %+v", request, response)
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
