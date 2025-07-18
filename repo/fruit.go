package repo

import (
    "crypto/rand"
    "math/big"
    "fruit-slot-api/config"
)

// getRandomIndex returns a secure random index
func getRandomIndex(max int64) int {
    nBig, _ := rand.Int(rand.Reader, big.NewInt(max))
    return int(nBig.Int64())
}

// GetRandomFruits returns 3 randomly selected fruits
func GetRandomFruits() []string {
    fruits := config.FruitList
    result := make([]string, 3)

    for i := 0; i < 3; i++ {
        result[i] = fruits[getRandomIndex(int64(len(fruits)))]
    }

    return result
}

// CheckWin determines if at least 2 fruits are the same
func CheckWin(fruits []string) bool {
    counts := make(map[string]int)
    for _, fruit := range fruits {
        counts[fruit]++
        if counts[fruit] >= 2 {
            return true
        }
    }
    return false
}
