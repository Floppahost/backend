package configs

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

var FiberApp = fiber.New(fiber.Config{

	// estamos usando um encoder de JSON customizado para melhor desempenho
	// para mais informações, ler: docs.gofiber.io/guide/faster-fiber

	JSONEncoder:   json.Marshal,
	JSONDecoder:   json.Unmarshal,
	Prefork:       true, // roda o servidor em várias instâncias; leia: docs.gofiber.io/api/fiber
	StrictRouting: true, // basicamente /login/ é diferente de /login; leia: docs.gofiber.io/api/fiber
	ServerHeader:  "BR-01",
	AppName:       "JUJU Backend",
})

