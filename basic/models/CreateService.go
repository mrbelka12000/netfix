package models

type CreateService struct {
	Name        string  `json:"name"`
	WorkField   string  `json:"workField"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Date        string  `json:"date"`
}
