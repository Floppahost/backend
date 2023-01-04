package preferences

import (
	"fmt"

	"github.com/floppahost/backend/database"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func GetEmbed(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	token := headers["Authorization"]

	embed, err := database.GetEmbed(token)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": true, "message": fmt.Sprintf("%s", err)})
	}

	jsonEmbed, _ := json.Marshal(embed)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success", "data": string(jsonEmbed)})
}