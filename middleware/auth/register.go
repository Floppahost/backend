package auth

import (
	"github.com/floppahost/backend/model"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	// we declare the type of the parser
	parser := new(model.Users)

	// we verify if there are any errors in the request
	if err := c.BodyParser(parser); err != nil {
		return err
	}

	return c.SendString("a")
}
