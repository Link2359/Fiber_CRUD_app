package main

import (
	"fiber-app/database"
	"fiber-app/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()


	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Home Page")
	})

	students := app.Group("/students")
	
	students.Get("/read/", handlers.GetAllStudents)
	students.Get("/read/:id", handlers.GetStudent)
	students.Post("/create/", handlers.CreateStudent)

	app.Listen(":3000")
}