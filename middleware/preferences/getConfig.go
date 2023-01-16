package preferences

import (
	"fmt"
	"os"

	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetConfig(c *fiber.Ctx) error {

	token := c.Cookies("token")

	userClaims := database.VerifyUser(token)
	if !userClaims.ValidUser || userClaims.Blacklisted {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": true, "message": "you don't have permission to perform this action"})
	}

	config := fmt.Sprintf(`{
		"Version": "14.0.0",
		"Name": "floppa.host",
		"DestinationType": "ImageUploader, FileUploader",
		"RequestMethod": "POST",
		"RequestURL": "https://api.floppa.host/files/upload",
		"Headers": {
		  "Authorization": "%s",
		},
		"Body": "MultipartFormData",
		"FileFormName": "file",
		"URL": "{json:url}",
		"ThumbnailURL": "{json:file_url}",
		"ErrorMessage": "{json:error}"
	  }`, userClaims.ApiKey)

	random_uuid := uuid.NewString()
	os.WriteFile(os.Getenv("FILE_PATH")+random_uuid, []byte(config), 0644)
	c.Download(os.Getenv("FILE_PATH")+random_uuid, "floppa-host-config.sxcu")
	err := os.Remove(os.Getenv("FILE_PATH") + random_uuid)
	if err != nil {
		panic("PANIC! couldn't remove temporary config")
	}
	return c.Status(400).JSON(fiber.Map{"error": false, "message": "Success"})
}
