package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	app.Use(recover.New())

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hi, there!")
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
