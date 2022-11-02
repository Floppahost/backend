package router

import (
	root "github.com/floppahost/backend/middleware"
	auth "github.com/floppahost/backend/middleware/auth"
	profile "github.com/floppahost/backend/middleware/profile"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func ConnectRouter(app *fiber.App) {
	rootRoute := app.Group("/", logger.New(logger.Config{}))
	authRoute := app.Group("/auth", logger.New(logger.Config{}))
	profileRoute := app.Group("/profile", logger.New(logger.Config{}))

	profile.Routes(profileRoute)
	auth.Routes(authRoute)
	root.Routes(rootRoute)
}
