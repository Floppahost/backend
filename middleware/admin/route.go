package admin

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router) {
	group := router.Group("/")

	group.Post("/wave", InviteWave)
	group.Post("/invites/purge", PurgeInvites)
	group.Post("/invite/user", GenerateInvite)
	group.Post("/blacklist", BlacklistUser)
	group.Post("/unblacklist", UnblacklistUser)
}
