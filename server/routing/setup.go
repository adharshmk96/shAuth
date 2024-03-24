package routing

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "github.com/adharshmk96/shAuth/docs"
)

func SetupRoutes(server *fiber.App) {
	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("SH Auth Server is running!")
	})

	server.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	server.Get("/swagger/*", swagger.HandlerDefault) // default

	apiRoutes := server.Group("/api")

	setupAccountRoutes(apiRoutes)
}
