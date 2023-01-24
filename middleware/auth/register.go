package auth

import (
	"fmt"
	"regexp"

	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/handler"
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

	re := regexp.MustCompile(`[[:^ascii:]\s]`)
	test_username := re.FindAllString(parser.Username, -1)
	test_email := re.FindAllString(parser.Email, -1)
	test_password := re.FindAllString(parser.Password, -1)
	if len(test_username) > 0 || len(test_email) > 0 || len(test_password) > 0 {
		fmt.Println("invalid characters on username, email or password")
	}
	token, err := database.Register(parser.Username, parser.Password, parser.Email, parser.Invite)

	if err != nil {
		status, errString := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": errString})
	}

	c.Cookie(&fiber.Cookie{
		Domain:   ".floppa.host",
		Path:     "/",
		Name:     "token",
		Value:    token,
		Secure:   true,
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "user registered"})
}
