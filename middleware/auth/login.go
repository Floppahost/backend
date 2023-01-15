package auth

import (
	// externos

	"github.com/gofiber/fiber/v2"

	// locais
	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/handler"
)

type Req struct {
	Username string `json:"username" xml:"username"`
	Password string `json:"password" xml:"password"`
}

func Login(c *fiber.Ctx) error {
	// declaramos o tipo do parser
	parser := new(Req)
	// verificamos se há erros no BodyParser
	if err := c.BodyParser(parser); err != nil {
		return err
	}
	token, err := database.Login(parser.Username, parser.Password)

	if err != nil {
		status, errString := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": errString})
	}

	c.Set("access-control-allow-origin", "floppa.host")
	c.Set("access-control-allow-credentials", "true")
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Secure:   true,
		HTTPOnly: true,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "logged in."})

}
