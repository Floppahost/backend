package files

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router) {
	group := router.Group("/")

	group.Post("/upload", Upload)
	group.Get("/render/:id", Render)

}
