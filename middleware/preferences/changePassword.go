package preferences

import (
	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/handler"
	"github.com/gofiber/fiber/v2"
)

type Reque struct {
	OldPasword  string `json:"old_password" xml:"old_password"`
	NewPassword string `json:"new_password" xml:"new_password"`
}

func ChangePassword(c *fiber.Ctx) error {
	parser := new(Reque)
	if err := c.BodyParser(parser); err != nil {
		return err
	}

	headers := c.GetReqHeaders()
	token := headers["Authorization"]

	err := database.ChangePassword(token, parser.OldPasword, parser.NewPassword)

	if err != nil {
		status, errMsg := handler.Errors(err)
		return c.Status(status).JSON(fiber.Map{"error": true, "message": errMsg})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success"})
}
