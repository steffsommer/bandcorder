package controllers

import (
	"maps"
	"server/internal/pkg/interfaces"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type WebsocketEvent struct {
	name interfaces.EventID
	data any
}

type WebsocketController struct {
	connections map[string]*websocket.Conn
}

func NewWebsocketController() *WebsocketController {
	return &WebsocketController{
		connections: make(map[string]*websocket.Conn),
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// HandleWebsocketUpgrade handles the connection upgrade from HTTP to a websocket connection
// If a client has an existing connection, the old connection is closed and a new
// one created.
func (r *WebsocketController) HandleWebsocketUpgrade(c *gin.Context) {
	clientIp := c.ClientIP()
	existingConn, exists := r.connections[clientIp]
	if exists {
		existingConn.Close()
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	defer conn.Close()
	for {
		conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
		time.Sleep(time.Second)
	}
}

func (r *WebsocketController) Send(event interfaces.EventID, data any) {
	ev := WebsocketEvent{
		name: event,
		data: data,
	}
	for conn := range maps.Values(r.connections) {
		err := conn.WriteJSON(ev)
		if err != nil {
			logrus.Errorf("Failed to send message %+v via websocket: %s", data, err.Error())
		}
	}
}
