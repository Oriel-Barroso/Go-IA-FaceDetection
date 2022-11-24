package main

import (
	"github.com/Oriel-Barroso/golangBackend/Imagen"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	Imagen.SetupRoutes(app)
	app.Listen(":3000")
}
