package controllers

import (
	"server/internal/pkg/models"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type WebsocketController struct {
	connections map[string]*websocket.Conn
	mutex       sync.Mutex
}

// NewRecordingController creates a new WebsocketController
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
// Connections are deemed healthy for as long as writes to it are successful. As soon as one
// write fails, the connection is closed.
// The server will respond with PONG messages as an answer to PINGs. Clients should ping
// frequently to determine their connection health.
func (r *WebsocketController) HandleWebsocketUpgrade(c *gin.Context) {
	r.mutex.Lock()
	clientIp := c.ClientIP()
	existingConn, exists := r.connections[clientIp]
	if exists {
		logrus.Errorf("Client with IP %s already has a connection, closing the old one", clientIp)
		err := existingConn.Close()
		if err != nil {
			logrus.Errorf("Failed to close connection of client with IP %s", clientIp)
		}
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.Errorf("Failed to upgrade connection for client with IP: %s: %s", clientIp, err.Error())
		return
	}
	conn.SetPingHandler(func(appData string) error {
		deadline := time.Now().Add(time.Second)
		return conn.WriteControl(websocket.PongMessage, []byte(appData), deadline)
	})
	r.connections[clientIp] = conn
	r.mutex.Unlock()

	// Read incoming messages (required for ping handler to get invoked)
	go func() {
		defer func() {
			r.mutex.Lock()
			conn.Close()
			delete(r.connections, clientIp)
			r.mutex.Unlock()
		}()

		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				break
			}
		}
	}()
}

// Dispatch sends an event to all public Websocket clients (app clients)
// If sending to a client fails, a connection error is assumed and the
// connection gets closed
func (r *WebsocketController) Dispatch(event models.EventLike) {
	r.mutex.Lock()
	for ip, conn := range r.connections {
		err := conn.WriteJSON(event)
		if err != nil {
			logrus.Errorf("Failed to send to %s: %s", ip, err.Error())
			conn.Close()
			delete(r.connections, ip)
		}
	}
	r.mutex.Unlock()
}
