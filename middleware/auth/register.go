package auth

import (
	"fmt"

	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
)

type Request struct {
	Email    string `json:"email" xml:"email"`
	Password string `json:"password" xml:"password"`
	Invite   string `json:"invite" xml:"invite"`
	Username string `json:"username" xml:"username"`
}

func Register(c *fiber.Ctx) error {
	// we declare the type of the parser
	parser := new(Request)

	// we verify if there are any errors in the request
	if err := c.BodyParser(parser); err != nil {
		return err
	}
	err := database.Register(parser.Username, parser.Password, parser.Email, parser.Invite)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": fmt.Sprintf("%s", err)})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "User registered"})
}
