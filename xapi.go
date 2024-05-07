package xapi

import (
	"fmt"
	"time"

	socket "github.com/MateoGreil/xapi-go/internal/protocols/socket"
	stream "github.com/MateoGreil/xapi-go/internal/protocols/stream"
	"github.com/gorilla/websocket"
)

type client struct {
	conn            *websocket.Conn
	streamConn      *websocket.Conn
	streamSessionId string
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
	getKeepAlive(conn, streamSessionId)

	c := &client{
		conn:            conn,
		streamConn:      streamConn,
		streamSessionId: streamSessionId,
	}
	go c.pingSocket()
	go c.pingStream()

	return c, nil
}

func (c *client) pingSocket() {
	for {
		request := socket.Request{
			Command:   "ping",
			Arguments: nil,
		}
		c.conn.WriteJSON(request)
		response := socket.Response{}
		err := c.conn.ReadJSON(&response)
		if err != nil {
			fmt.Println(err.Error())
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
		c.streamConn.WriteJSON(request)
		time.Sleep(pingInterval)
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
}
