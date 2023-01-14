package files

import (
	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
)

func Render(c *fiber.Ctx) error {

	uploadId := c.Params("id")
	file, err := database.GetUpload(uploadId)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": false, "message": "the upload doesn't exist"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success", "data": file})
}
