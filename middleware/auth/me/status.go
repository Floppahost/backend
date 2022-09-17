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
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Not authorized.", "data": nil, "auth": false})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"error": false, "message": "Authenticated", "data": nil, "auth": true})
}
