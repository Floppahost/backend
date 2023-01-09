package preferences

import (
	"errors"
	"fmt"
	"strings"

	"github.com/floppahost/backend/database"
	"github.com/gofiber/fiber/v2"
)

type Reques struct {
	Path	string `json:"path" xml:"path"`
	Mode 	string `json:"mode" xml:"mode"`
	Amount	int `json:"amount" xml:"amount"`
}

func ChangePath(c *fiber.Ctx) error {
	parser := new(Reques)
	if err := c.BodyParser(parser); err != nil {
		return err
	}

	headers := c.GetReqHeaders()
	token := headers["Authorization"]
	err := validate(parser)

	if err != nil {
		return c.Status(406).JSON(fiber.Map{"error": true, "message": fmt.Sprintf("%v", err)})
	}

	if parser.Mode == "custom" {
		err := database.ChangePathMode(token, parser.Mode, parser.Amount)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": true, "message": fmt.Sprintf("%v", err)})
		}

		err = database.ChangePathValue(token, parser.Path)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": true, "message": fmt.Sprintf("%v", err)})
		}
	}

	err = database.ChangePathMode(token, parser.Mode, parser.Amount)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": true, "message": fmt.Sprintf("%v", err)})
	}
	
	return c.Status(200).JSON(fiber.Map{"error": false, "message": "success"})
}

func validate (parser *Reques) error {
	mode := parser.Mode
	err := validateMode(mode, parser.Path, parser.Amount)
	
	if err != nil {
		return err
	}

	if mode == "custom" && parser.Path != "" {
		path := parser.Path
		err := validatePath(path)

		if err != nil {
			return err
		}
	}

	return nil
}

func validateMode(mode string, path string, amount int) error {
	if mode != "random" && mode != "custom" && mode != "invisible" && mode != "emoji" && mode != "amongus" && mode != "amongus_emoji" {
		return errors.New("invalid path mode")
	}

	if (mode == "amongus" || mode == "amongus_emoji" || mode == "emoji") && amount <= 0 {
		return errors.New("the amount can not be less than one")
	}

	if (mode == "amongus" || mode == "amongus_emoji" || mode == "emoji") && amount > 5 {
		return errors.New("the amount can not be greater than five")
	}
	if mode == "custom" && path == "" {
		return errors.New("path can not be blank")
	}

	return nil
}

func validatePath(path string) error {
	if strings.Contains(path, "/") ||  strings.Contains(path, "\\"){
		return errors.New("the path can not contain any kind of slashes")
	}

	if len(path) > 10 {
		return errors.New("the path can not contain more than 10 characters")
	}

	return nil
}