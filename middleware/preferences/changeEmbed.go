package preferences

import (
	"fmt"

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

	err := database.UpdateEmbed(token, parser.Author, parser.Description, parser.Title, parser.Name)

	if err != nil {
	if fmt.Sprintf("%v", err) == "unauthorized" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": false, "message": "unauthorized"})
	}
	
		return c.Status(500).JSON(fiber.Map{"error": true, "message": fmt.Sprintf("%s", err)})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success"})
}