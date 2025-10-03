package controllers

import (
	"server/internal/pkg/models"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type WebsocketController struct {
	connections map[string]*websocket.Conn
	mutex       sync.Mutex
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
	r.mutex.Lock()
	clientIp := c.ClientIP()
	existingConn, exists := r.connections[clientIp]
	if exists {
		existingConn.Close()
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	r.connections[clientIp] = conn
	r.mutex.Unlock()
}

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
