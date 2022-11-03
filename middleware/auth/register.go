package auth

import (
	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
)

type Request struct {
	email    string
	password string
	invite   string
	username string
}

func Register(c *fiber.Ctx) error {
	// we declare the type of the parser
	parser := new(Request)

	// we verify if there are any errors in the request
	if err := c.BodyParser(parser); err != nil {
		return err
	}

	database.Register(parser.username, parser.password, parser.email, parser.invite)
	return c.SendString("a")
}
