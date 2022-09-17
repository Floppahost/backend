package router

import (
	auth "github.com/floppahost/backend/middleware/auth/me"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func ConnectRouter(app *fiber.App) {
	authRoute := app.Group("/auth", logger.New(logger.Config{}))

	// auth_Clientes.Routes(auth)
	auth.Routes(authRoute)
}

