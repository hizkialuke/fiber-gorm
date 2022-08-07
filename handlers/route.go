package handlers

import (
	"fiber-gorm/middleware"
	"github.com/gofiber/fiber/v2"
)

func InitRoute(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	r.Post("/login", Login)
	r.Get("/users", middleware.JwtRequired(), GetAllUsers)
}
