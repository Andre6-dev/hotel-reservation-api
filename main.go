package main

import (
	"flag"

	"github.com/Andre6-dev/hotel-reservation-api/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddress := flag.String("ListenAddress", ":3000", "Address to listen on")
	flag.Parse()
	// Create new Fiber instance
	app := fiber.New()

	// Create route on root path, "/"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandlerListUser)
	apiv1.Get("/user/:id", api.HandlerGetUser)

	// Start server on port 3000
	app.Listen(*listenAddress)
}
