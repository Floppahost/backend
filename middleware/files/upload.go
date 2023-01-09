package files

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/AvraamMavridis/randomcolor"
	"github.com/floppahost/backend/buck"
	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/lib"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func Upload(c *fiber.Ctx) error {
	startTime := time.Now()
	headers := c.GetReqHeaders()
	apikey := headers["Authorization"]
	userClaims := database.VerifyUser(apikey)
	if (!userClaims.ValidUser) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": false, "message": "unauthorized"})
	}

	ctx := context.Background()
	file, err := c.FormFile("file")

	if err != nil {
		return err
	}

	Bucket := buck.Bucket
	fsPath := fmt.Sprintf("%s%s", os.Getenv("FILE_PATH"), file.Filename)
	filePath := fmt.Sprintf("./%s", fsPath)
	c.SaveFile(file, filePath)
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

	info, err := Bucket.FPutObject(ctx, "files", objectName, fsPath, minio.PutObjectOptions{ContentType: fileHeader})
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{"error": true, "message": "something weird happened. Please, try again; if the error persists, contact the support"})
	}

	fileSize := file.Size
	os.Remove(fsPath)

	embed, erro := database.GetEmbed(apikey)

	author := func() string {b := fmt.Sprintf("%v", embed["author"]); if embed["author"] == nil {b = ""}; return b}
	description := func() string {b := fmt.Sprintf("%v", embed["description"]); if embed["description"] == nil {b = ""}; return b}
	name := func() string {b := fmt.Sprintf("%v", embed["name"]); if embed["name"] == nil {b = ""}; return b}
	title := func() string {b := fmt.Sprintf("%v", embed["title"]); if embed["title"] == nil {b = ""}; return b}
	embedColor := func() string {b := fmt.Sprintf("%v", embed["title"]); if embed["title"] == nil {b = ""}; return b}
	domain := fmt.Sprintf("%v", embed["domain"])
	path_mode := fmt.Sprintf("%v", embed["path_mode"])
	path_amount, _ := strconv.ParseInt(fmt.Sprintf("%v", embed["path_amount"]), 10, 64)
	path_value := func() string {b := fmt.Sprintf("%v", embed["path"]); if embed["path"] == nil {b = "hello"}; return b}
	enabled, _ := strconv.ParseBool(fmt.Sprintf("%v", embed["enabled"]))

	color := embedColor()
	
	if color == "random" {
		color = randomcolor.GetRandomColorInHex()
	}

	if erro != nil {
		return erro
	}
	
	var path string
	var upload_url string
	switch path_mode {
		case "amongus":
			path = lib.AmongUs(int(path_amount))
			upload_url = path
		case "amongus_emoji":
			path = lib.AmongUsAndEmoji(int(path_amount))
			upload_url = path
		case "emoji":
			path = lib.RandomEmoji(int(path_amount))
			upload_url = path
		case "invisible":
			path = lib.InvisibleUrl(15)
			upload_url = path
		case "custom":
			path = path_value() + lib.InvisibleUrl(15)
			upload_url = path
		default: 
			path = generated_uuid
			upload_url = path
	}
	file_url := "cdn.floppa.host/files/" +  objectName
	upload_url = domain + "/i/" + upload_url

	endTime := time.Since(startTime)
	embedFields := lib.EmbedPlaceholders(title(), description(), name(), author(), apikey, fileSize, fileName, endTime)
	err = database.Upload(embedFields.Author, embedFields.Name, embedFields.Description, embedFields.Title, enabled, path, userClaims.Uid, info.Key, color, generated_uuid, fileName, file_url, upload_url, apikey)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": true, "message": "something wrong happened"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success", "url": upload_url, "file_url": file_url})
}
