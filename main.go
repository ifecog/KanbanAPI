package main

import (
	"courier-api/config"
	"courier-api/models"
	"courier-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

// func main() {
// 	app := fiber.New()
// 	config.ConnectDatabase()

// 	// Migrate models
// 	config.DB.AutoMigrate(&models.Courier{}, &models.Package{})

// 	// Register routes
// 	routes.CourierRoutes(app)

// 	app.Listen(":3000")
// }

func main() {
    app := fiber.New()
    config.ConnectDatabase()
    // Migrate models
    log.Println("Migrating models...")
    err := config.DB.AutoMigrate(&models.Courier{}, &models.Package{})
    if err != nil {
        log.Fatal("Error migrating models:", err)
    }
    // Register routes
    routes.CourierRoutes(app)
    app.Listen(":3000")
}