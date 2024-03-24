package handler

import (
	"github.com/adharshmk96/shAuth/core"
	"github.com/adharshmk96/shAuth/server/infra"
	"github.com/go-playground/validator/v10"
	"log/slog"
)

type authHandler struct {
	accountService core.AccountService
	validator      *validator.Validate
	logger         *slog.Logger
}

func New(service core.AccountService) core.AuthHandler {
	validate := validator.New(validator.WithRequiredStructEnabled())
	logger := infra.GetLogger()

	return &authHandler{
		logger:         logger,
		accountService: service,
		validator:      validate,
	}
}
