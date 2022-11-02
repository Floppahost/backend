package auth

import (
	"github.com/floppahost/backend/jwt"
	"github.com/gofiber/fiber/v2"
)

func Status(c *fiber.Ctx) error {
	header := c.GetReqHeaders()
	token := header["Authorization"]
	_, err := jwt.Validar(token)

	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": false, "message": "Not authorized", "auth": false})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"error": false, "message": "Authenticated", "auth": true})
}
