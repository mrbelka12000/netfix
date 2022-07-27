package models

import "errors"

type Apply struct {
	ID         int `json:"ID"`
	CustomerID int `json:"customerID"`
	Work
	StartDate int64 `json:"startDate"`
	EndDate   int64 `json:"endDate"`
}

type Work struct {
	WorkID    int     `json:"workID"`
	Price     float64 `json:"price"`
	CompanyID int     `json:"companyID"`
}

func (wa *Apply) Validate() error {

	if wa.ID == 0 {
		return errors.New("invalid apply id")
	}
	if wa.CustomerID == 0 {
		return errors.New("invalid customer id")
	}
	if wa.WorkID == 0 {
		return errors.New("invalid work id")
	}
	if wa.StartDate == 0 {
		return errors.New("invalid start date")
	}
	if wa.EndDate == 0 {
		return errors.New("invalid end date")
	}
	return nil
}
