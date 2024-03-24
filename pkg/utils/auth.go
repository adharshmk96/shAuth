package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func IsAuthenticated(c *fiber.Ctx) bool {
	authCookie := c.Cookies(viper.GetString("auth.cookie_name"))
	if authCookie == "" {
		return false
	}

	claims, err := DecodeJWT(authCookie)
	if err != nil {
		return false
	}

	_, ok := claims["email"].(string)
	if !ok {
		return false
	}

	return true
}
