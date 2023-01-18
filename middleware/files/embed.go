package files

import "github.com/gofiber/fiber/v2"

type ParseQuery struct {
	AuthorName   string `query:"author_name"`
	AuthorUrl    string `query:"author_url"`
	ProviderName string `query:"provider_name"`
	ProviderUrl  string `query:"provider_url"`
}

func RenderEmbed(c *fiber.Ctx) error {
	parser := new(ParseQuery)

	if err := c.QueryParser(parser); err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"type": "link", "version": "1.0", "author_name": parser.AuthorName, "author_url": parser.AuthorUrl, "provider_url": parser.ProviderUrl, "provider_name": parser.ProviderName})

}
