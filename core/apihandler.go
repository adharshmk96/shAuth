package core

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	// Auth
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	Profile(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error

	// Oauth
	GoogleLogin(c *fiber.Ctx) error
	GoogleLoginCallback(c *fiber.Ctx) error
}
