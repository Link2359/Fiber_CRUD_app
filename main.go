package main

import (
	"fiber-app/database"
	"fiber-app/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database.Connect()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Home Page")
	})

	students := app.Group("/students")

	students.Get("/read/", handlers.GetAllStudents)
	students.Get("/read/:id", handlers.GetStudent)
	students.Post("/create/", handlers.CreateStudent)
	students.Delete("/delete/:id", handlers.DeleteStudent)
	students.Patch("/update/:id", handlers.UpdateStudent)

	app.Listen(":3000")
}
