package main

import (
	"os"

	_ "main.go/docs"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // load .env file automatic
	"main.go/pkg/configs"
	"main.go/pkg/middlewares"
	"main.go/pkg/routes"
	"main.go/pkg/utils"
	"main.go/platform/database"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @BasePath /api
// @in header
// @name Authorization
func main() {
	// Define Fiber config.
	configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New()

	// Database
	database.OpenDBConnection()

	// middlewares
	app.Use(middlewares.ErrorHandlingMiddleware)

	// Routes.
	routes.SwaggerRoute(app)  // Register route for swagger.
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
