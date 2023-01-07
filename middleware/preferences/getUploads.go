package preferences

import (
	"fmt"

	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
)

type QueryParser struct {
	Page     int     `query:"page"`
}
func GetUploads(c *fiber.Ctx) error {
	parser := new(QueryParser)

	if err := c.QueryParser(parser); err != nil {
		return err
	}

	headers := c.GetReqHeaders()
	token := headers["Authorization"]

	uploads, maxPages, err := database.GetUploads(token, parser.Page)

	if err != nil {
		if fmt.Sprintf("%v", err) == "unauthorized" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": false, "message": "unauthorized"})
		}
		return c.Status(500).JSON(fiber.Map{"error": true, "message": fmt.Sprintf("%s", err)})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success", "data": uploads, "max_pages": maxPages})
}