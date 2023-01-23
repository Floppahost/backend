package configs

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

var FiberApp = fiber.New(fiber.Config{

	// we are using a diferent json encoder to boost performace; read: docs.gofiber.io/guide/faster-fiber
	BodyLimit: 	   10 * 1024 * 1024,
	JSONEncoder:   json.Marshal,
	JSONDecoder:   json.Unmarshal,
	Prefork:       true, // read: docs.gofiber.io/api/fiber
	ServerHeader:  "REGION-01",
	AppName:       "Floppa.Host",
})

