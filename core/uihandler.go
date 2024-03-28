package core

import "github.com/gofiber/fiber/v2"

type WebUIHandler interface {
	Index(c *fiber.Ctx) error
}
