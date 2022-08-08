package middleware

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

var SecretKey = "HIZKIA_LUKE"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "-", err
	}

	return webToken, nil
}

func VerifyToken(stringToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected login method")
		}

		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func JwtRequired() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(SecretKey),
		ErrorHandler: jwtError,
	})
}

func jwtError(ctx *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing or malformed JWT",
		})
	} else {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token invalid",
		})
	}
}

type LogedUser struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Exp      int64  `json:"exp"`
}

func GetLogedUser(ctx context.Context) *LogedUser {
	user := ctx.Value("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userID := int64(claims["id"].(float64))
	exp := int64(claims["exp"].(float64))
	return &LogedUser{
		ID:       userID,
		Name:     claims["name"].(string),
		UserName: claims["user_name"].(string),
		Exp:      exp,
	}
}
