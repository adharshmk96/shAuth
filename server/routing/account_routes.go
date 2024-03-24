package routing

import (
	"github.com/adharshmk96/shAuth/internals/handler"
	"github.com/adharshmk96/shAuth/internals/service"
	"github.com/adharshmk96/shAuth/internals/storage/memory"
	"github.com/adharshmk96/shAuth/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupAccountRoutes(router fiber.Router) {
	accStorage := memory.New()
	accService := service.New(accStorage)
	accHandler := handler.New(accService)

	// Auth routes
	router.Post("/auth/register", accHandler.Register)
	router.Post("/auth/login", accHandler.Login)
	router.Post("/auth/logout", accHandler.Logout)
	router.Get("/auth/profile", middleware.AuthMiddleware, accHandler.Profile)
	router.Post("/auth/change-password", middleware.AuthMiddleware, accHandler.ChangePassword)

	// Oauth routes
	router.Get("/oauth/google", accHandler.GoogleLogin)
	router.Get("/oauth/google/callback", accHandler.GoogleLoginCallback)
}
