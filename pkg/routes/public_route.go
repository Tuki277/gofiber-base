package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/app/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	//heath check
	a.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	// Create routes group.
	route := a.Group("/api/v1")

	// base
	base := route.Group("/base")
	base.Get("/", controllers.BaseList)
	base.Get("/:id", controllers.BaseGetById)
	base.Post("/", controllers.BaseCreate)
	base.Put("/:id", controllers.BaseUpdate)
	base.Delete("/:id", controllers.BaseDelete)
}
