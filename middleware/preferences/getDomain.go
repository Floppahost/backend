package preferences

import (
	"fmt"

	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
)

func GetDomain(c *fiber.Ctx) error {

	headers := c.GetReqHeaders()
	token := headers["Authorization"]

	embed, err := database.GetEmbed(token)

	if err != nil {
		if fmt.Sprintf("%v", err) == "unauthorized" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": false, "message": "unauthorized"})
		}
		return c.Status(500).JSON(fiber.Map{"error": true, "message": fmt.Sprintf("%s", err)})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success", "domain": embed["domain"]})
}