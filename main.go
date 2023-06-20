package main

import (
	"context"
	"flag"
	"log"

	"github.com/Andre6-dev/hotel-reservation-api/api"
	"github.com/Andre6-dev/hotel-reservation-api/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"

// Create a new fiber instance with custom config
var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddress := flag.String("ListenAddress", ":5000", "Address to listen on")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	// Create new Fiber instance
	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	// Define routes
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandlerListUsers)
	apiv1.Get("/user/:id", userHandler.HandlerGetUser)
	apiv1.Delete("/user/:id", userHandler.HandlerDeleteUser)
	apiv1.Put("/user/:id", userHandler.HandlerPutUser)
	// Start server on port 3000
	app.Listen(*listenAddress)
}
