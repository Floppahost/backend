package files

import (
	"context"
	"fmt"
	"os"

	"github.com/floppahost/backend/buck"
	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)


func Upload(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	apikey := headers["Apikey"]
	userClaims := database.VerifyUserApiKey(apikey)
	if (!userClaims.ValidUser) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": false, "message": "Unauthorized"})
	}

	ctx := context.Background()
	file, err := c.FormFile("file")

	if err != nil {
		return err
	}
	Bucket := buck.Bucket
	path := fmt.Sprintf("%s%s", os.Getenv("FILE_PATH"), file.Filename)
	c.SaveFile(file, fmt.Sprintf("./%s", path))
	_ = Bucket.MakeBucket(context.Background(), userClaims.Username, minio.MakeBucketOptions{})
	bucketName := userClaims.Username
	fileName := file.Filename
	fileHeader := fmt.Sprintf("%s", file.Header)


	info, err := Bucket.FPutObject(ctx, bucketName, fileName, path, minio.PutObjectOptions{ContentType: fileHeader})
	if err != nil {
		return err
	}
	os.Remove(path)

	return c.SendString(info.Bucket)
}
