package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/st4rgaze/otaqu/app/routes"
	"github.com/st4rgaze/otaqu/config"
	"github.com/st4rgaze/otaqu/database/migrations"
	"github.com/st4rgaze/otaqu/scheduler"
)

func main() {
	app := fiber.New()

	// Load environment variables
	err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the database
	config.ConnectDB()

	// Run migrations
	err = migrations.CreateHotelsTable()
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Set up routes
	routes.Setup(app)

	// Run the scheduler in a separate goroutine
	go scheduler.Run()

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
