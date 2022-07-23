package models

type CreateWork struct {
	ID          int     `json:"ID"`
	CompanyID   int     `json:"companyID"`
	Name        string  `json:"name"`
	WorkField   string  `json:"workField"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Date        string  `json:"date"`
}
