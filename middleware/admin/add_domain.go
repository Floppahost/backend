package admin

import (
	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/handler"
	"github.com/gofiber/fiber/v2"
)

type Reque struct {
	Domain   string `json:"domain" xml:"domain"`
	Wildcard bool   `json:"wildcard" xml:"wildcard"`
	Username string `json:"username" xml:"username"`
}

func AddDomain(c *fiber.Ctx) error {
	token := c.Cookies("token")
	parser := new(Reque)
	if err := c.BodyParser(parser); err != nil {
		return err
	}
	err := database.AddDomain(token, parser.Domain, parser.Wildcard, parser.Username)
	if err != nil {
		status, errString := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": errString})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success"})
}
