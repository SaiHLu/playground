package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/playground/h-reservation/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbUri          = "mongodb://root:root@localhost:27017"
	dbName         = "hotel-reservation"
	userCollection = "users"
)

func main() {
	_, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	listenAddr := flag.String("listenAddr", ":5000", "The listen port of the server")
	flag.Parse()

	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/foo", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"Message": "Hello World"})
	})

	apiV1.Get("/users", api.HandleGetUsers)
	apiV1.Get("/users/:id", api.HandleGetUser)

	app.Listen(*listenAddr)
}
