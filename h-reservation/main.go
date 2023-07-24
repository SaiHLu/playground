package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/playground/h-reservation/api"
	"github.com/playground/h-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbUri          = "mongodb://root:root@localhost:27017"
	dbName         = "hotel-reservation"
	userCollection = "users"
)

var fiberConfig = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	listenAddr := flag.String("listenAddr", ":5000", "The listen port of the server")
	flag.Parse()

	app := fiber.New(fiberConfig)
	apiV1 := app.Group("/api/v1")

	userHandler := api.NewUserHandler(db.NewMongoDBStore(client))

	apiV1.Get("/foo", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"Message": "Hello World"})
	})

	apiV1.Get("/users", userHandler.HandleGetUsers)
	apiV1.Get("/users/:id", userHandler.HandleGetUser)

	app.Listen(*listenAddr)
}
