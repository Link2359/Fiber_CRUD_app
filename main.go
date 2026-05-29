package main

import (
	"fiber-app/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()


	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Home Page")
	})

	app.Listen(":3000")
}