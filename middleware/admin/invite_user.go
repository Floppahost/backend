package admin

import (
	"fmt"
	"os"
	"strconv"

	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
)

type Req struct {
	Username string `json:"username" xml:"username"`
}

func GenerateInvite(c *fiber.Ctx) error {
	parser := new(Req)
	if invite, _ := strconv.ParseBool(os.Getenv("INVITE_ONLY")); !invite {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"error": true, "message": "Invite system isn't enabled"})
	}
	if err := c.BodyParser(parser); err != nil {
		return err
	}

	headers := c.GetReqHeaders()
	err := database.GenerateInvite(headers["Authorization"], parser.Username)

	if err != nil {
		errString := fmt.Sprintf("%v", err)
		errCode := func() int {
			switch errString {
			case "you don't have permission to perform this action":
				return 401
			case "the requested user doesn't exist":
				return 404
			}
			return 500
		}
		return c.Status(errCode()).JSON(fiber.Map{"error": false, "message": errString})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "Success"})
}
