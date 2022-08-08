package handlers

import (
	"fiber-gorm/middleware"
	"fiber-gorm/models"
	"fiber-gorm/repo"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(ctx *fiber.Ctx) error {
	loginReq := new(models.LoginRequest)
	if err := ctx.BodyParser(loginReq); err != nil {
		return err
	}

	data, err := repo.Login(loginReq)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if data.ID != 0 {
		// generate jwt
		claims := jwt.MapClaims{}
		claims["id"] = data.ID
		claims["name"] = data.Name
		claims["user_name"] = data.UserName
		claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

		generateToken, err := middleware.GenerateToken(&claims)
		if err != nil && generateToken != "-" {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "something gone wrong",
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": generateToken,
		})
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "user not found",
	})
}

func GetAllUsers(ctx *fiber.Ctx) error {
	pagination := new(models.Pagination)
	if err := ctx.QueryParser(pagination); err != nil {
		return err
	}

	if pagination.Page == 0 {
		pagination.Page = models.DefaultPage
	}
	if pagination.Limit == 0 {
		pagination.Limit = models.DefaultLimit
	}

	data, err := repo.GetAllUsers(pagination)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	return ctx.JSON(data)
}
