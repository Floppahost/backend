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
	parser := new(model.BaseUser)

	// verificamos se há erros no BodyParser
	if err := c.BodyParser(parser); err != nil {
		return err
	}
	id, err := database.Login(model.BaseUser{}, parser.Usuario, parser.Senha)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": true, "message": "Dados incorretos.", "data": nil})
	}
	token, err := jwt.GenerateUserToken("system", parser.Usuario, id, true)
	if err != nil {
		panic(err)
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token

	c.Cookie(cookie)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Logado com sucesso.", "data": nil})

}
