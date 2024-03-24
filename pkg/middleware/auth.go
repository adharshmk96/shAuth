package middleware

import (
	"github.com/adharshmk96/shAuth/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authCookie := c.Cookies(viper.GetString("auth.cookie_name"))
	if authCookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	claims, err := utils.DecodeJWT(authCookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	email, ok := claims["email"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// to access email in other handlers
	c.Locals("email", email)

	return c.Next()
}
