package auth

import (
	// externos
	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router) {
	group := router.Group("/@me")

	// Rota de login
	group.Post("/status", Status)

}

