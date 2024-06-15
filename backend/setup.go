package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func config() fiber.Config {
	return fiber.Config{
		AppName:      "GoHadits",
		ErrorHandler: errorHandler,
	}
}

// GET /api/v1/
func rootV1Handler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"maintainer": "M. Iqbal Effendi <iqbaleff214@gmail.com>",
		"license":    "MIT",
		"source":     "https://github.com/404NotFoundIndonesia/ngatur-erp",
		"version":    "1",
	})
}

// Error handler response
func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	err = c.Status(code).JSON(map[string]any{
		"code":    code,
		"message": e.Message,
		"status":  "error",
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]any{
			"code":    fiber.StatusInternalServerError,
			"message": "Internal Server Error",
			"status":  "error",
		})
	}

	return nil
}
