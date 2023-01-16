package preferences

import (
	"fmt"

	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
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

	return c.Status(200).JSON(fiber.Map{"error": false, "message": "success", "config": config})
}
