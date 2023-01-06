package preferences

import (
	"fmt"

	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
)

type Request struct {
	 Domain string  `json:"domain" xml:"domain"`
}

func ChangeDomain(c *fiber.Ctx) error {
	parser := new(Request)
	if err := c.BodyParser(parser); err != nil {
		return err
	}

	headers := c.GetReqHeaders()
	token := headers["Authorization"]
	err := database.UpdateDomain(token, parser.Domain)

	if err != nil {
		if fmt.Sprintf("%v", err) == "unauthorized" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": false, "message": "unauthorized"})
		}
		return c.Status(404).JSON(fiber.Map{"error": false, "message": "this domain doesn't exist"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "success"})
}