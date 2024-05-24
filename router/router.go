package router

import (
	"game_planner_backend/database"
	"game_planner_backend/handler"
	"game_planner_backend/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Docs
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
	}))

	// Middleware
	api := app.Group("/api", logger.New())

	// User
	userHandler := handler.NewUserHandler(repository.NewUserRepo(database.DB))
	user := api.Group("/user")
	user.Post("/", userHandler.CreateUser)
	user.Get("/", userHandler.GetAllUsers)
	user.Get("/:id", userHandler.GetOneUser)
	user.Put("/:id", userHandler.UpdateUser)
	user.Delete("/:id", userHandler.DeleteUser)
}
