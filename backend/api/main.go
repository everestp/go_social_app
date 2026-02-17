package main

import (
	"log"
	"server/database"
	_ "server/docs" // Import the generated Swagger docs
	"server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

// @title Fiber Golang RestApi
// @version 1.0
// @description This is the Swagger docs for REST API Golang Fiber
// @host localhost:5000
// @BasePath /
// @schema http
// @securityDefinitions.apiKey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by space and the token
func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	// Connect to the database
	database.Connect()

	// Initialize Fiber app
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOriginsFunc: func(origin string) bool {
			return true
		},
	}))

	// Root route
	// @Summary Welcome message
	// @Description Get welcome message for the social app
	// @Tags root
	// @Accept json
	// @Produce json
	// @Success 200 {string} string "Welcome to social app"
	// @Router / [get]
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to social app")
	})
 // setup routes
   routes.SetupRoutes(app)
	// Swagger docs route
	app.Get("/swagger/*", swagger.HandlerDefault) // now it works with docs package

	// Start server
	if err := app.Listen(":5000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
