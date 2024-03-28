package server

import (
	"embed"
	"github.com/adharshmk96/shAuth/internals/apihandler"
	"github.com/adharshmk96/shAuth/internals/service"
	"github.com/adharshmk96/shAuth/internals/storage/memory"
	"github.com/adharshmk96/shAuth/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/swagger"
	"net/http"
)

//go:embed public_dist/*
var webAssets embed.FS

//go:embed public_dist/index.html
var indexHTML embed.FS

func setupRoutes(server fiber.Router) {

	accStorage := memory.New()
	accService := service.New(accStorage)
	accHandler := apihandler.New(accService)

	// API routes
	apiRoutes := server.Group("/api")

	// Auth routes
	apiRoutes.Post("/auth/register", accHandler.Register)
	apiRoutes.Post("/auth/login", accHandler.Login)
	apiRoutes.Post("/auth/logout", accHandler.Logout)
	apiRoutes.Get("/auth/profile", middleware.AuthMiddleware, accHandler.Profile)
	apiRoutes.Post("/auth/change-password", middleware.AuthMiddleware, accHandler.ChangePassword)

	// Oauth routes
	apiRoutes.Get("/oauth/google", accHandler.GoogleLogin)
	apiRoutes.Get("/oauth/google/callback", accHandler.GoogleLoginCallback)

	// Common routes
	server.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	server.Get("/swagger/*", swagger.HandlerDefault) // default

	// UI routes
	server.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(webAssets),
		PathPrefix: "public_dist/static",
	}))
	server.Use("/icons", filesystem.New(filesystem.Config{
		Root:       http.FS(webAssets),
		PathPrefix: "public_dist/icons",
	}))
	server.Get("/ui/*", func(c *fiber.Ctx) error {
		return c.Render("public_dist/index", fiber.Map{})
	})
	server.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/ui")
	})
}
