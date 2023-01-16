package admin

import (
	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/handler"
	"github.com/gofiber/fiber/v2"
)

func UnblacklistUser(c *fiber.Ctx) error {
	type req struct {
		Username string `json:"username" xml:"username"`
		Reason   string `json:"reason" xml:"reason"`
	}

	token := c.Cookies("token")
	parser := new(req)
	if err := c.BodyParser(parser); err != nil {
		return err
	}
	err := database.UnblacklistUser(token, parser.Username, parser.Reason)

	if err != nil {
		status, errString := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": errString})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success"})
}
