package files

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/AvraamMavridis/randomcolor"
	"github.com/floppahost/backend/buck"
	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	fileName := file.Filename
	fileHeader := fmt.Sprintf("%s", file.Header)

	fileNameSplit := strings.Split(fileName, ".")
	fileExtension := ""

	if len(fileNameSplit) >= 1 {
	for i := 1; len(fileNameSplit) > i; i++ {
			fileExtension = fileExtension + "." + fileNameSplit[i]
	}
}

	generated_uuid := uuid.NewString()
	objectName := generated_uuid + fileExtension

	info, err := Bucket.FPutObject(ctx, "files", objectName, path, minio.PutObjectOptions{ContentType: fileHeader})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": true, "message": "something weird happened. Please, try again; if the error persists, contact the support"})
	}
	os.Remove(path)

	embed, erro := database.GetEmbed(apikey)

	author := func() string {b := fmt.Sprintf("%v", embed["author"]); if embed["author"] == nil {b = ""}; return b}
	description := func() string {b := fmt.Sprintf("%v", embed["description"]); if embed["description"] == nil {b = ""}; return b}
	name := func() string {b := fmt.Sprintf("%v", embed["name"]); if embed["name"] == nil {b = ""}; return b}
	title := func() string {b := fmt.Sprintf("%v", embed["title"]); if embed["title"] == nil {b = ""}; return b}
	embedColor := func() string {b := fmt.Sprintf("%v", embed["title"]); if embed["title"] == nil {b = ""}; return b}
	enabled, _ := strconv.ParseBool(fmt.Sprintf("%v", embed["enabled"]))

	color := embedColor()
	if color == "random" {
		color = randomcolor.GetRandomColorInHex()
	}

	if erro != nil {
		return erro
	}
	

	database.Upload(author(), name(), description(), title(), enabled, userClaims.Uid, info.Key, color, generated_uuid, fileName)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success", "uploadId": generated_uuid})
}
