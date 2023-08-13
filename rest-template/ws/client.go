package ws

import (
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type connection struct {
	ws   *websocket.Conn
	send chan []byte
}

type subscription struct {
	conn *connection
	room string
}

func (c *connection) write(messageType int, data []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(messageType, data)
}

func (s subscription) readPump() {
	connection := s.conn

	defer func() {
		h.unregister <- s
		connection.ws.Close()
	}()

	connection.ws.SetReadLimit(maxMessageSize)
	connection.ws.SetReadDeadline(time.Now().Add(pongWait))
	connection.ws.SetPongHandler(func(string) error { connection.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, msg, err := connection.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("read messages error: %+v\n", err)
			}
			break
		}
		m := message{msg, s.room}
		h.broadcast <- m
	}
}

func (s *subscription) writePump() {
	connection := s.conn
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		close(connection.send)
	}()

	for {
		select {
		case message, ok := <-connection.send:
			if !ok {
				connection.write(websocket.CloseMessage, []byte{})
				return
			}

			if err := connection.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := connection.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}

	}
}
