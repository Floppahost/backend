package files

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router) {
	group := router.Group("/")

	group.Put("/", Upload)
	group.Get("/render/:id", Render)
	group.Delete("/:file", DeleteFile)
	group.Get("/render/embed", Render)
}
