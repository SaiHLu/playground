package ws

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func WsRoute(app *fiber.App) fiber.Router {
	// go h.Run()

	api := app.Group("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	api.Get("/:roomId", websocket.New(func(c *websocket.Conn) {
		log.Println("room: ", c.Params("roomId"))

		app.Static("/", "./public")

		serveWs(c, c.Params("roomId"))
	}))

	return api

}
