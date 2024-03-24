package handler

import (
	"errors"
	"github.com/adharshmk96/shAuth/core"
	"github.com/adharshmk96/shAuth/core/model"
	"log/slog"
	"time"

	"github.com/adharshmk96/shAuth/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email" example:"shuser@maildrop.cc"`
	Username string `json:"username" example:"user1234"`
	Password string `json:"password" validate:"required,min=8,max=20" example:"Pa$$w0rd!"`
}

// @Summary	Register an account
// @Router		/api/auth/register [post]
// @Param		RegisterRequest	body	RegisterRequest	true	"Register Request"
func (h *authHandler) Register(c *fiber.Ctx) error {
	var registerRequest RegisterRequest

	if err := c.BodyParser(&registerRequest); err != nil {
		h.logger.Error(
			"register handler: cannot parse JSON",
			slog.String("error", err.Error()),
		)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	err := h.validator.Struct(registerRequest)
	if err != nil {
		h.logger.Error(
			"register handler: validation error",
			slog.String("error", err.Error()),
		)
		return utils.ValidationErrors(c, err.(validator.ValidationErrors))
	}

	accData := &model.Account{
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}

	err = h.accountService.RegisterAccount(accData)
	if err != nil {
		message := "cannot register account"
		if errors.Is(err, core.ErrAccountExists) {
			message = "account already exists"
		}

		h.logger.Error(
			"register handler: cannot register account",
			slog.String("error", err.Error()),
		)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": message,
		})
	}

	h.logger.Info("Register Handler: account created", "email", accData.ID)

	signedJWT, err := h.accountService.GenerateJWT(accData)
	if err != nil {
		h.logger.Error(
			"register handler: cannot generate token: ",
			slog.String("error", err.Error()),
		)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot generate token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     viper.GetString("auth.cookie_name"),
		Value:    signedJWT,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		SameSite: "Lax",
	})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "account created",
		"redirect": viper.GetString("auth.redirect_url"),
	})
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"shuser@maildrop.cc"`
	Password string `json:"password" validate:"required" example:"Pa$$w0rd!"`
}

// @Summary	Login account
// @Router		/api/auth/login [post]
// @Param		loginRequest	body	LoginRequest	true	"Login Request"
func (h *authHandler) Login(c *fiber.Ctx) error {
	var loginRequest LoginRequest

	if err := c.BodyParser(&loginRequest); err != nil {
		h.logger.Error(
			"login handler: cannot parse JSON",
			slog.String("error", err.Error()),
		)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	err := h.validator.Struct(loginRequest)
	if err != nil {
		h.logger.Error(
			"login handler: validation error",
			slog.String("error", err.Error()),
		)
		return utils.ValidationErrors(c, err.(validator.ValidationErrors))
	}

	acc, err := h.accountService.Authenticate(loginRequest.Email, loginRequest.Password)
	if err != nil {
		h.logger.Error(
			"login handler: cannot authenticate account: ",
			slog.String("error", err.Error()),
		)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "cannot authenticate account",
		})
	}

	signedJWT, err := h.accountService.GenerateJWT(acc)
	if err != nil {
		h.logger.Error(
			"login handler: cannot generate token: ",
			slog.String("error", err.Error()),
		)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot generate token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     viper.GetString("auth.cookie_name"),
		Value:    signedJWT,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		SameSite: "Lax",
	})

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message":  "login success",
		"redirect": viper.GetString("auth.redirect_url"),
	})
}

// @Summary	Profile of account
// @Router		/api/auth/profile [get]
func (h *authHandler) Profile(c *fiber.Ctx) error {
	email := c.Locals("email").(string)

	acc, err := h.accountService.GetAccountByEmail(email)
	if err != nil {
		utils.ClearCookie(c, viper.GetString("auth.cookie_name"))

		h.logger.Error(
			"profile handler: cannot get account",
			slog.String("error", err.Error()),
		)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unauthorized",
			"login": viper.GetString("auth.login_url"),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"email": acc.Email,
	})
}

// @Summary	Logout account
// @Router		/api/auth/logout [post]
func (h *authHandler) Logout(c *fiber.Ctx) error {
	utils.ClearCookie(c, viper.GetString("auth.cookie_name"))

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "logout success",
	})
}

type ChangePasswordRequest struct {
	Email       string `json:"email" validate:"required,email" example:"shuser@maildrop.cc"`
	Password    string `json:"password" validate:"required,min=8,max=20" example:"Pa$$w0rd!"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=20" example:"Pa$$word1"`
}

// @Summary	Reset password
// @Router		/api/auth/reset-password [post]
// @Param		resetPasswordRequest	body	ChangePasswordRequest	true	"Reset Password Request"
func (h *authHandler) ChangePassword(c *fiber.Ctx) error {
	var changePasswordRequest ChangePasswordRequest

	if err := c.BodyParser(&changePasswordRequest); err != nil {
		h.logger.Error(
			"reset password handler: cannot parse JSON",
			slog.String("error", err.Error()),
		)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	err := h.validator.Struct(changePasswordRequest)
	if err != nil {
		h.logger.Error(
			"reset password handler: validation error",
			slog.String("error", err.Error()),
		)
		return utils.ValidationErrors(c, err.(validator.ValidationErrors))
	}

	err = h.accountService.ChangePassword(
		changePasswordRequest.Email,
		changePasswordRequest.Password,
		changePasswordRequest.NewPassword,
	)
	if err != nil {
		h.logger.Error(
			"reset password handler: cannot reset password",
			slog.String("error", err.Error()),
		)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot reset password",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "password reset success",
	})
}
