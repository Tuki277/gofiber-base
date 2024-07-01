package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/pkg/constanst"
	"main.go/pkg/response"
)

// NotFoundRoute func for describe 404 Error route.
func NotFoundRoute(a *fiber.App) {
	// Register new special route.
	a.Use(
		// Anonymous function.
		func(c *fiber.Ctx) error {
			err := fiber.NewError(fiber.StatusNotFound, constanst.ENDPOINT_NOT_FOUND)
			// Return HTTP 404 status and JSON response.
			return c.Status(fiber.StatusNotFound).JSON(response.ResponseData(nil, nil, err, fiber.StatusNotFound))
		},
	)
}
