package main

import (
	"fiberexample/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	bootstrap()
	
	app := fiber.New()
	routes.Setup(app)

	app.Listen(":3000")
}

func bootstrap() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}