package admin

import (
	"os"
	"strconv"

	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
)

type Req struct {
	Username string `json:"username" xml:"username"`
}

func GenerateInvite(c *fiber.Ctx) error {
	parser := new(Req)
	if invite, _ := strconv.ParseBool(os.Getenv("INVITE_ONLY")); !invite {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"error": true, "message": "Invite system isn't enabled"})
	}
	if err := c.BodyParser(parser); err != nil {
		return err
	}
	headers := c.GetReqHeaders()
	database.GenerateInvite(headers["Authorization"], parser.Username)
	return c.SendString("a")
}
