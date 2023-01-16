package preferences

import (
	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/handler"
	"github.com/gofiber/fiber/v2"
)

type Req struct {
	Author      string `json:"author" xml:"author"`
	Title       string `json:"title" xml:"title"`
	Name        string `json:"name" xml:"name"`
	Description string `json:"description" xml:"description"`
	Color       string `json:"color" xml:"color"`
}

func ChangeEmbed(c *fiber.Ctx) error {
	parser := new(Req)
	if err := c.BodyParser(parser); err != nil {
		return err
	}

	token := c.Cookies("token")

	err := database.UpdateEmbed(token, parser.Author, parser.Description, parser.Title, parser.Name, parser.Color)

	if err != nil {
		status, errMsg := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": errMsg})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success"})
}
