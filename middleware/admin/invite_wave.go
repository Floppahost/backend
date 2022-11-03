package admin

import (
	"os"
	"strconv"

	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
)

func InviteWave(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	if invite, _ := strconv.ParseBool(os.Getenv("INVITE_ONLY")); !invite {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"error": true, "message": "Invite system isn't enabled"})
	}
	err := database.InviteWave(headers["Authorization"])
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success"})
}
