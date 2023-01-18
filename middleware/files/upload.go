package files

import (
	"context"
	"fmt"
	"net/url"
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
	if !userClaims.ValidUser {
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
	fileHeader := fmt.Sprintf("%v", file.Header)

	fileNameSplit := strings.Split(fileName, ".")
	fileExtension := ""

	if len(fileNameSplit) >= 1 {
		for i := 1; len(fileNameSplit) > i; i++ {
			fileExtension = fileExtension + "." + fileNameSplit[i]
		}
	}

	generated_uuid := uuid.NewString()
	objectName := generated_uuid + fileExtension

	_, err = Bucket.FPutObject(ctx, "files", objectName, fsPath, minio.PutObjectOptions{ContentType: fileHeader})
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{"error": true, "message": "something weird happened. Please, try again; if the error persists, contact the support"})
	}

	fileSize := file.Size
	os.Remove(fsPath)

	embed, err := database.GetEmbed(apikey)

	author := func() string {
		b := fmt.Sprintf("%v", embed["author"])
		if embed["author"] == nil {
			b = ""
		}
		return b
	}

	author_url := func() string {
		b := fmt.Sprintf("%v", embed["author_url"])
		if embed["author_url"] == nil {
			b = ""
		}
		return b
	}

	description := func() string {
		b := fmt.Sprintf("%v", embed["description"])
		if embed["description"] == nil {
			b = ""
		}
		return b
	}
	site_name := func() string {
		b := fmt.Sprintf("%v", embed["site_name"])
		if embed["site_name"] == nil {
			b = ""
		}
		return b
	}

	site_name_url := func() string {
		b := fmt.Sprintf("%v", embed["site_name_url"])
		if embed["site_name_url"] == nil {
			b = ""
		}
		return b
	}

	title := func() string {
		b := fmt.Sprintf("%v", embed["title"])
		if embed["title"] == nil {
			b = ""
		}
		return b
	}
	embedColor := func() string {
		b := fmt.Sprintf("%v", embed["color"])
		if embed["color"] == nil {
			b = ""
		}
		return b
	}
	domain := fmt.Sprintf("%v", embed["domain"])
	path_mode := fmt.Sprintf("%v", embed["path_mode"])
	path_amount, _ := strconv.ParseInt(fmt.Sprintf("%v", embed["path_amount"]), 10, 64)
	path_value := func() string {
		b := fmt.Sprintf("%v", embed["path"])
		if embed["path"] == nil {
			b = "hello"
		}
		return b
	}
	color := embedColor()

	if color == "random" {
		color = randomcolor.GetRandomColorInHex()
	}

	if err != nil {
		return err
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
	file_url := fmt.Sprintf("%s/files/%s", os.Getenv("BUCKET_ENDPOINT"), objectName)
	upload_url = fmt.Sprintf("%s/i/%s", domain, path)

	endTime := time.Since(startTime)
	embedFields := lib.EmbedPlaceholders(site_name(), site_name_url(), title(), description(), author(), author_url(), apikey, fileSize, fileName, endTime)

	urlEscape := url.QueryEscape(path)
	err = database.Upload(embedFields.SiteName, embedFields.SiteNameUrl, embedFields.Title, embedFields.Description, embedFields.Author, embedFields.AuthorUrl, color, userClaims.Uid, fileName, file_url, upload_url, urlEscape, objectName, apikey, fileHeader)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": true, "message": "something wrong happened"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success", "url": upload_url, "file_url": file_url})
}
