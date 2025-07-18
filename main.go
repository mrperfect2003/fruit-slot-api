package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "log"
    "os"
    "fruit-slot-api/routes"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found. Using default port 3000")
    }

    // Get PORT from .env or fallback to default
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    // Initialize Fiber app
    app := fiber.New()

    // Register API routes
    routes.FruitRoutes(app)

    // Start the server
    log.Println("🚀 Server running at http://localhost:" + port)
    log.Fatal(app.Listen(":" + port))
}
