package model

type MenuItem struct {
    PK          string  `json:"PK"`
    SK          string  `json:"SK"`
    Description string  `json:"description"`
    GSI1        string  `json:"GSI1"`
    Image       string  `json:"image"`
    Name        string  `json:"name"`
    Price       float64 `json:"price"`
}
