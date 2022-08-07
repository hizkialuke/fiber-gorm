package main

import (
	"fiber-gorm/database"
	"fiber-gorm/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// init db
	database.InitDB()

	// init route
	app := fiber.New()
	handlers.InitRoute(app)

	app.Listen(":8000")
}
