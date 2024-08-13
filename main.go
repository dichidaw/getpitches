package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func main() {
	config := viper.New()
	config.SetConfigFile(".env")
	config.AddConfigPath(".")

	errConfig := config.ReadInConfig()
	if errConfig != nil {
		panic(errConfig)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	app.Use(recover.New())

	app.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hi, there!")
	})

	app.Get("/spotify", func(ctx *fiber.Ctx) error {
		client := resty.New()
		url := "https://accounts.spotify.com/api/token"

		payload := map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     config.GetString("CLIENT_ID"),
			"client_secret": config.GetString("CLIENT_SECRET"),
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/x-www-form-urlencoded").
			SetFormData(payload).
			Post(url)

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString("Error making request")
		}

		// Send the response back
		return ctx.SendString(resp.String())
	})

	errPort := app.Listen(":3000")
	if errPort != nil {
		panic(errPort)
	}
}
