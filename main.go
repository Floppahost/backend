package main

import (
	// externos
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	// locais

	"github.com/floppahost/backend/buck"
	"github.com/floppahost/backend/configs"
	"github.com/floppahost/backend/database"
	"github.com/floppahost/backend/router"
)

func main() {

	// declaramos a var app para o Fiber, tudo via a nossa config
	app := configs.FiberApp
	app.Use(cors.New(cors.Config{AllowCredentials: true}))
	// carregamos o dotenv e verificamos se ele está funcionando
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	// conectamos à database
	database.Connect()

	// we start our bucket
	buck.Start()
	// configuramos as rotas
	router.ConnectRouter(app)

	// conectamos o app à porta 3000
	app.Listen(":4000")

}
