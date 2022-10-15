package profile

import "github.com/gofiber/fiber/v2"

func GetData(c *fiber.Ctx) error {
	return c.SendString(c.Params("profile"))
}