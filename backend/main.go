package main

import (
	"4th_Assignment/db"
	"4th_Assignment/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db.ConnectDb()

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}
		return c.Next()
	})

	app.Post("/submit-contact", handlers.ContactSubmit)

	app.Listen(":3000")

}
