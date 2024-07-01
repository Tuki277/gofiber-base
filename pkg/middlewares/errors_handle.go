package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"main.go/pkg/constanst"
	"main.go/pkg/response"
)

func ErrorHandlingMiddleware(c *fiber.Ctx) error {
	defer func() {
		if err := recover(); err != nil {
			errData := fiber.NewError(fiber.StatusInternalServerError, constanst.INTERNAL_SERVER_ERROR)
			c.Status(fiber.StatusInternalServerError).JSON(response.ResponseData(nil, nil, errData, fiber.StatusInternalServerError))
		}
	}()
	return c.Next()
}
