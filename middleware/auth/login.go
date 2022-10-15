package auth

import (
	// externos
	"github.com/gofiber/fiber/v2"

	// locais
	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/jwt"
	"github.com/floppahost/backend/model"
)

func Login(c *fiber.Ctx) error {
	// declaramos o tipo do parser
	parser := new(model.Users)

	// verificamos se há erros no BodyParser
	if err := c.BodyParser(parser); err != nil {
		return err
	}
	id, err := database.Login(model.Users{}, parser.User, parser.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": true, "message": "Invalid data.", "data": nil})
	}

	token, err := jwt.GenerateUserToken("system", parser.User, id)
	if err != nil {
		panic(err)
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token

	c.Cookie(cookie)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Logged in.", "data": nil})

}
