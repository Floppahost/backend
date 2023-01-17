package admin

import (
	"os"
	"strconv"

	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/handler"
	"github.com/gofiber/fiber/v2"
)

func PurgeInvites(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if invite, _ := strconv.ParseBool(os.Getenv("INVITE_ONLY")); !invite {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"error": true, "message": "Invite system isn't enabled"})
	}
	err := database.PurgeInvites(token)

	if err != nil {
		status, errString := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": errString})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success"})
}
