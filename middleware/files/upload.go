package files

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/floppahost/backend/buck"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)

func Upload(c *fiber.Ctx) error {
	ctx := context.Background()
	file, err := c.FormFile("file")

	if err != nil {
		return err
	}
	path := fmt.Sprintf("./%s", os.Getenv("FILE_PATH"))
	c.SaveFile(file, fmt.Sprintf("./%s", path))
	bucketName := os.Getenv("MINION_BUCKET_NAME")
	fileName := file.Filename
	fileHeader := fmt.Sprintf("%s", file.Header)

	Bucket := buck.Bucket

	info, err := Bucket.FPutObject(ctx, bucketName, fileName, path, minio.PutObjectOptions{ContentType: fileHeader})
	if err != nil {
		log.Fatalln(err)
	}

	return c.SendString(info.Bucket)
}
