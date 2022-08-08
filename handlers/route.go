package handlers

import (
	"fiber-gorm/middleware"
	"github.com/gofiber/fiber/v2"
)

func InitRoute(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	r.Post("/api/login", Login)
	r.Get("/api/users", middleware.JwtRequired(), GetAllUsers)
	r.Get("/api/report-a", middleware.JwtRequired(), GetReportA)
	r.Get("/api/report-b", middleware.JwtRequired(), GetReportB)
}
