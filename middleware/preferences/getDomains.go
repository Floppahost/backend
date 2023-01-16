package preferences

import (
	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/handler"
	"github.com/gofiber/fiber/v2"
)

func GetDomains(c *fiber.Ctx) error {

	token := c.Cookies("token")

	domains, err := database.GetDomains(token)

	if err != nil {
		status, errMsg := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": errMsg})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success", "domains": domains})
}
