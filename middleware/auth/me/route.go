package auth

import (
	// externos
	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router) {
	group := router.Group("/@me")

	group.Post("/status", Status)
	group.Post("/login", Login)

}

