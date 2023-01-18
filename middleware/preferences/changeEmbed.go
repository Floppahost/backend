package preferences

import (
	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/handler"
	"github.com/gofiber/fiber/v2"
)

type Req struct {
	SiteName    string `json:"site_name" xml:"site_name"`
	SiteNameUrl string `json:"site_name_url" xml:"site_name_url"`
	Title       string `json:"title" xml:"title"`
	Description string `json:"description" xml:"description"`
	Author      string `json:"author" xml:"author"`
	AuthorUrl   string `json:"author_url" xml:"author_url"`
	Color       string `json:"color" xml:"color"`
}

func ChangeEmbed(c *fiber.Ctx) error {
	parser := new(Req)
	if err := c.BodyParser(parser); err != nil {
		return err
	}

	token := c.Cookies("token")

	err := database.UpdateEmbed(token, parser.SiteName, parser.SiteName, parser.Title, parser.Description, parser.Author, parser.Author, parser.Color)

	if err != nil {
		status, errMsg := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": errMsg})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success"})
}
