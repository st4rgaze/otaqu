package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/st4rgaze/otaqu/app/controllers"
)

// routes list API
func Setup(app *fiber.App) {
	api := app.Group("/api")

	users := api.Group("/hotel")
	users.Get("/", controllers.GetAllHotels)
}
