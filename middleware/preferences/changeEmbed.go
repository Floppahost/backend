package preferences

import (
	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
)

type Req struct {
	 Author	string `json:"author" xml:"author"`
	 Title  string  `json:"title" xml:"title"`
	 Name   string  `json:"name" xml:"name"`
	 Description string  `json:"description" xml:"description"`
}

func ChangeEmbed(c *fiber.Ctx) error {
	parser := new(Req)
	if err := c.BodyParser(parser); err != nil {
		return err
	}

	headers := c.GetReqHeaders()
	token := headers["Authorization"]
	database.UpdateEmbed(token, parser.Author, parser.Description, parser.Title, parser.Name)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success"})
}