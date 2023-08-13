package ws

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func WsRoute(app *fiber.App) fiber.Router {
	api := app.Group("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		log.Println("Upgrade weboskcet error")
		return fiber.ErrUpgradeRequired
	})

	api.Get("/:roomId", websocket.New(func(c *websocket.Conn) {
		log.Println("allowed: ", c.Locals("allowed"))

		serveWs(c, c.Params("roomId"))
	}))

	return api

}
