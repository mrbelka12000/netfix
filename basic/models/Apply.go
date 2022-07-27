package models

type WorkActions struct {
	ID         int `json:"ID"`
	CustomerID int `json:"customerID"`
	WorkA
	StartDate int64 `json:"startDate"`
	EndDate   int64 `json:"endDate"`
}
type WorkA struct {
	WorkID    int     `json:"workID"`
	Price     float64 `json:"price"`
	CompanyID int     `json:"companyID"`
}
