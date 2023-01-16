package files

import (
	"context"

	"github.com/floppahost/backend/buck"
	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)

func DeleteFile(c *fiber.Ctx) error {
	token := c.Cookies("token")
	file := c.Params("file")

	_, err := database.ValidateUpload(token, file)

	if err != nil {
		status, err := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": err})
	}
	bucket := buck.Bucket

	err = bucket.RemoveObject(context.Background(), "files", file, minio.RemoveObjectOptions{})

	if err != nil {
		status, _ := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": "something unexpected happened; contact an admin"})
	}

	_, err = database.DeleteUpload(token, file)

	if err != nil {
		status, err := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": err})
	}
	return c.SendString(file)
}
