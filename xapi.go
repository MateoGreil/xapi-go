package xapi

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MateoGreil/xapi-go/internal/protocols/socket"
	"github.com/MateoGreil/xapi-go/internal/protocols/stream"
	"github.com/gorilla/websocket"
)

type client struct {
	conn                 *websocket.Conn
	streamConn           *websocket.Conn
	streamSessionId      string
	socketMessageChannel chan interface{}
	streamMessageChannel chan interface{}
	CandlesChannel       chan stream.Candle
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

	c := &client{
		conn:                 conn,
		streamConn:           streamConn,
		streamSessionId:      streamSessionId,
		socketMessageChannel: make(chan interface{}),
		streamMessageChannel: make(chan interface{}),
		CandlesChannel:       make(chan stream.Candle),
	}
	go c.pingSocket()
	go c.pingStream()
	go c.listenStream()
	go c.socketWriteJSON()
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
		request := socket.Request{
			Command:   "ping",
			Arguments: nil,
		}
		c.socketMessageChannel <- request
		// response := socket.Response{}
		// err := c.conn.ReadJSON(&response)
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
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

func (c *client) socketWriteJSON() {
	for {
		message := <-c.socketMessageChannel
		c.conn.WriteJSON(message)
		fmt.Printf("messageSocket: %+v\n", message)
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
