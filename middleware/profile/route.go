package profile

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router) {
	group := router.Group("/")

	group.Post("get/:profile", GetData)

}

