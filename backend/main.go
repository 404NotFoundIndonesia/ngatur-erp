package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(config())
	app.Use(cors.New())

	api := app.Group("api/v1")

	api.Get("/", rootV1Handler)

	log.Fatal(app.Listen(":8000"))
}
