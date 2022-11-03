package admin

import (
	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
)

func InviteWave(c *fiber.Ctx) error {
	database.InviteWave("")
	return c.SendString("a")
}
