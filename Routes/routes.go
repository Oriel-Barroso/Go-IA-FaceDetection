package Routes

import (
	imagenController "github.com/Oriel-Barroso/golangBackend/Imagen"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/postImage", imagenController.PostImage)
}
