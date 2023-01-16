package auth

import (
	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
)

func Status(c *fiber.Ctx) error {
	token := c.Cookies("token")
	userClaims := database.VerifyUser(token)

	if !userClaims.ValidUser || userClaims.Blacklisted {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": true, "message": "not authorized", "auth": false})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"error": false, "message": "authenticated", "auth": true, "username": userClaims.Username, "uid": userClaims.Uid})
}
