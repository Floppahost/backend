package router

import (
	"github.com/floppahost/backend/middleware/admin"
	"github.com/floppahost/backend/middleware/auth"
	"github.com/floppahost/backend/middleware/files"
	"github.com/floppahost/backend/middleware/preferences"
	"github.com/floppahost/backend/middleware/profile"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func ConnectRouter(app *fiber.App) {
	filesRoute := app.Group("/files", logger.New(logger.Config{}))
	authRoute := app.Group("/auth", logger.New(logger.Config{}))
	profileRoute := app.Group("/profile", logger.New(logger.Config{}))
	adminRoute := app.Group("/admin", logger.New(logger.Config{}))
	preferencesRoute := app.Group("/preferences", logger.New(logger.Config{}))

	profile.Routes(profileRoute)
	auth.Routes(authRoute)
	files.Routes(filesRoute)
	admin.Routes(adminRoute)
	preferences.Routes(preferencesRoute)
}
