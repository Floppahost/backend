package admin

import (
	"os"
	"strconv"

	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/handler"
	"github.com/gofiber/fiber/v2"
)

type req struct {
	Username string `json:"username" xml:"username"`
}

func GenerateInvite(c *fiber.Ctx) error {
	parser := new(req)
	if invite, _ := strconv.ParseBool(os.Getenv("INVITE_ONLY")); !invite {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"error": true, "message": "Invite system isn't enabled"})
	}
	if err := c.BodyParser(parser); err != nil {
		return err
	}

	token := c.Cookies("token")
	err := database.GenerateInvite(token, parser.Username)

	if err != nil {
		status, errString := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": errString})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success"})
}
