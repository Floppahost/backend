package status

import (
	"os"
	"strconv"

	"github.com/floppahost/backend/lib"
	"github.com/gofiber/fiber/v2"
)

func GetInviteStatus(c *fiber.Ctx) error {
	lib.InvisibleUrl(10)
	lib.RandomEmoji(10)
	lib.AmongUs(10)
	lib.AmongUsAndEmoji(10)
	invite := os.Getenv("INVITE_ONLY")

	inviteBool, err := strconv.ParseBool(invite)

	if err != nil {
		panic("invalid invite only value; must be a bool")
	}
	return c.Status(400).JSON(fiber.Map{"error": false, "message": "success", "inviteOnly": inviteBool})
}