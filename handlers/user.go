package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"fiber-gorm/database"
	"fiber-gorm/middleware"
	"fiber-gorm/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func Login(ctx *fiber.Ctx) error {
	loginReq := new(models.LoginRequest)
	if err := ctx.BodyParser(loginReq); err != nil {
		return err
	}

	var model models.User
	hash := md5.Sum([]byte(loginReq.Password))
	encode := hex.EncodeToString(hash[:])
	err := database.DBConn.
		Where("user_name = ?", loginReq.UserName).
		Where("password = ?", encode).
		First(&model).Error

	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if model.ID != 0 {
		// generate jwt
		claims := jwt.MapClaims{}
		claims["id"] = model.ID
		claims["name"] = model.Name
		claims["user_name"] = model.UserName
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
	var model []*models.User
	err := database.DBConn.Find(&model).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println(err)
	}

	return ctx.JSON(model)
}
