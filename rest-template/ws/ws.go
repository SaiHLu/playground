package ws

import (
	"github.com/gofiber/contrib/websocket"
)

func serveWs(c *websocket.Conn, roomId string) {
	go h.Run()

	conn := &connection{ws: c, send: make(chan []byte)}
	sub := subscription{conn, roomId}

	h.register <- sub

	go sub.writePump()
	sub.readPump()
}
