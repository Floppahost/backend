package preferences

import (
	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/handler"
	"github.com/gofiber/fiber/v2"
)

type Request struct {
	Domain string `json:"domain" xml:"domain"`
}

func ChangeDomain(c *fiber.Ctx) error {
	parser := new(Request)
	if err := c.BodyParser(parser); err != nil {
		return err
	}

	token := c.Cookies("token")
	err := database.UpdateDomain(token, parser.Domain)

	if err != nil {
		status, errMsg := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": errMsg})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "success"})
}
