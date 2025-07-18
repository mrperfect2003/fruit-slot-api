package routes

import (
    "github.com/gofiber/fiber/v2"
    "fruit-slot-api/controller"
)

// FruitRoutes defines all routes under /play
func FruitRoutes(app *fiber.App) {
    fruit := app.Group("/play")
    fruit.Get("/", controller.PlayHandler)
    fruit.Get("/10", controller.PlayTenHandler)
}
