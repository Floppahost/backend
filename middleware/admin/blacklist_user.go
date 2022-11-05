package admin

import (
	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/handler"
	"github.com/gofiber/fiber/v2"
)

type Req struct {
	Username string `json:"username" xml:"username"`
	Reason string `json:"reason" xml:"reason"`
}

func BlacklistUser(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	parser := new(Req)
	if err := c.BodyParser(parser); err != nil {
		return err
	}
	err := database.BlacklistUser(headers["Authorization"],parser.Username, parser.Reason)

	if err != nil {
		status, errString := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": errString})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success"})
}