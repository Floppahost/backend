package preferences

import (
	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/handler"
	"github.com/gofiber/fiber/v2"
)

func GetDomain(c *fiber.Ctx) error {

	headers := c.GetReqHeaders()
	token := headers["Authorization"]

	embed, err := database.GetEmbed(token)

	if err != nil {
		status, errMsg := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": errMsg})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success", "domain": embed["domain"]})
}
