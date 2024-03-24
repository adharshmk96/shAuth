package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidationErrors(c *fiber.Ctx, err error) error {
	var validErr validator.ValidationErrors
	if !errors.As(err, &validErr) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
			"msg":   err.Error(),
		})
	}

	errorList := map[string]string{}

	for _, e := range validErr {
		errorList[e.Field()] = e.Error()
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Validation failed",
		"msg":   errorList,
	})
}
