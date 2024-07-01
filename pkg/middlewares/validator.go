package middlewares

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"main.go/pkg/constanst"
)

func ParseAndValidate(c *fiber.Ctx, data interface{}) error {
	if err := c.BodyParser(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, constanst.IS_VALID_JSON_REQUEST)
	}

	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}
