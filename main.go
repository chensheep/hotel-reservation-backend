package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/chensheep/hotel-reservation-backend/api"
	"github.com/chensheep/hotel-reservation-backend/db"
)

const (
	uri = "mongodb://localhost:27017"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(fiber.Map{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// initialize the handlers
	userStore := db.NewMongoUserStore(client)
	userHandler := api.NewUserHandler(userStore)

	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Post("/user", userHandler.HandlePostUser)

	app.Listen(*listenAddr)
}
