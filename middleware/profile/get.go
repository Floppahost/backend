package profile

import (
	"fmt"

	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	User 	string 	 `json:"username"`
	ID      string   `json:"ID"`
}


func GetData(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	apikey := headers["Authorization"]
	userClaims := database.VerifyUser(apikey)
	if (!userClaims.ValidUser) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": false, "message": "Unauthorized"})
	}

	profile := c.Params("profile")
	fmt.Println(profile)
	id, err := database.GetProfile(profile)

	if (err != nil) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": fmt.Sprintf("%v", err), "data": nil})
	}

	user := User{
		User: profile,
		ID: id,
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "success", "data": user})

	// FUTURE: redisgn user
}