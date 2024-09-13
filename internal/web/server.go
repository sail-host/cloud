package web

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sail-host/cloud/config"
)

func NewServer() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(fmt.Sprintf(":%s", config.GetConfig().Port))
}
