package profile

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Usuario string `json:"usuario"`
	ID      string   `json:"ID"`
}


func GetData(c *fiber.Ctx) error {
	profile := c.Params("profile")
	fmt.Println(profile)
	id, err := database.GetProfile(profile)

	if (err != nil) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": fmt.Sprintf("%v", err), "data": nil})
	}

	user := User{
		Usuario: profile,
		ID: id,
	}
	
	userJson, _ := json.Marshal(user)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "success", "data": string(userJson)})

	// FUTURE: redisgn user
}