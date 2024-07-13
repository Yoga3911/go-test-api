package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", sayHello)

	app.Listen(":3000")
}

func sayHello(c *fiber.Ctx) error {
	word := "Hello, World!!!"

	return c.JSON(fiber.Map{
		"message": word,
	})
}
