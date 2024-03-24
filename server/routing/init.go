package routing

import (
	_ "github.com/adharshmk96/shAuth/docs"
	"github.com/adharshmk96/shAuth/pkg/middleware"
	"github.com/adharshmk96/shAuth/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
)

func SetupRoutes(server *fiber.App) {
	// Use the "filesystem" middleware for serving static files
	server.Get("/ui/login", func(ctx *fiber.Ctx) error {

		if utils.IsAuthenticated(ctx) {
			return ctx.Redirect(viper.GetString("auth.redirect_url"))
		}

		return ctx.Render("login", fiber.Map{
			"Title":    "Login",
			"LoginApi": "/api/auth/login",
		}, "base")
	})

	server.Get("/ui/register", func(ctx *fiber.Ctx) error {

		if utils.IsAuthenticated(ctx) {
			return ctx.Redirect(viper.GetString("auth.redirect_url"))
		}

		return ctx.Render("register", fiber.Map{
			"Title":       "Register",
			"RegisterApi": "/api/auth/register",
		}, "base")
	})

	server.Get("/ui/profile", middleware.AuthRedirect, func(ctx *fiber.Ctx) error {
		email := ctx.Locals("email").(string)

		return ctx.Render("profile", fiber.Map{
			"Title":     "Profile",
			"Email":     email,
			"LogoutApi": "/api/auth/logout",
			"LoginUI":   viper.GetString("auth.login_url"),
		}, "base")
	})

	apiRoutes := server.Group("/api")

	setupAccountRoutes(apiRoutes)

	server.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	server.Get("/swagger/*", swagger.HandlerDefault) // default

	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("SH Auth Server is running!")
	})

}
