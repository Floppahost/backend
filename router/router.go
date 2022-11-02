package router

import (
	"github.com/floppahost/backend/middleware/auth"
	"github.com/floppahost/backend/middleware/files"
	"github.com/floppahost/backend/middleware/profile"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func ConnectRouter(app *fiber.App) {
	filesRoute := app.Group("/files", logger.New(logger.Config{}))
	authRoute := app.Group("/auth", logger.New(logger.Config{}))
	profileRoute := app.Group("/profile", logger.New(logger.Config{}))

	profile.Routes(profileRoute)
	auth.Routes(authRoute)
	files.Routes(filesRoute)
}
