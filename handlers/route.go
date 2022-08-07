package handlers

import "github.com/gofiber/fiber/v2"

func InitRoute(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Warudo")
	})
}
