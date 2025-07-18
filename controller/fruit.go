package controller

import (
    "github.com/gofiber/fiber/v2"
    "fruit-slot-api/repo"
    "fruit-slot-api/model"
)

// PlayHandler handles /play route (1 spin)
func PlayHandler(c *fiber.Ctx) error {
    fruits := repo.GetRandomFruits()
    message := "Try again!"
    if repo.CheckWin(fruits) {
        message = "You win!"
    }

    response := model.FruitResult{
        Fruits:  fruits,
        Message: message,
    }

    return c.Status(fiber.StatusOK).JSON(response)
}

// PlayTenHandler handles /play/10 route (10 spins)
func PlayTenHandler(c *fiber.Ctx) error {
    var allResults []model.FruitResult
    winCount := 0

    for i := 0; i < 10; i++ {
        fruits := repo.GetRandomFruits()
        message := "Try again!"
        if repo.CheckWin(fruits) {
            message = "You win!"
            winCount++
        }

        allResults = append(allResults, model.FruitResult{
            Fruits:  fruits,
            Message: message,
        })
    }

    return c.Status(fiber.StatusOK).JSON(model.MultiPlayResult{
        Results:  allResults,
        WinCount: winCount,
    })
}
