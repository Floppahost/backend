package preferences

import (
	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/handler"
	"github.com/gofiber/fiber/v2"
)

type QueryParser struct {
	Page int `query:"page"`
}

func GetUploads(c *fiber.Ctx) error {
	parser := new(QueryParser)

	if err := c.QueryParser(parser); err != nil {
		return err
	}

	token := c.Cookies("token")

	uploads, maxPages, err := database.GetUploads(token, parser.Page)

	if err != nil {
		status, errMsg := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": errMsg})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success", "data": uploads, "max_pages": maxPages})
}
