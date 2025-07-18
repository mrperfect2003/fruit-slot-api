package model

// FruitResult represents a single spin result
type FruitResult struct {
    Fruits  []string `json:"fruits"`
    Message string   `json:"message"`
}

// MultiPlayResult holds multiple spin results and total win count
type MultiPlayResult struct {
    Results  []FruitResult `json:"results"`
    WinCount int           `json:"win_count"`
}
