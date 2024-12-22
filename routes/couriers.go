package routes

import (
	"courier-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func CourierRoutes(app *fiber.App) {
	api := app.Group("/api/couriers")
	api.Post("/register", controllers.RegisterCourier)
}