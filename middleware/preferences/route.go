package preferences

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router) {
	change := router.Group("/change")
	get := router.Group("/get")

	change.Post("/domain", ChangeDomain)
	get.Get("/domain", GetDomain)

	change.Post("/embed", ChangeEmbed)
	get.Get("/embed", GetEmbed)

	get.Get("/domains", GetDomains)

	change.Post("/password", ChangePassword)

	get.Get("/uploads", GetUploads)

	get.Get("/config", GetConfig)

	get.Get("/invites", GetInvites)

	change.Post("/path", ChangePath)
}
