package models

import "errors"

type Work struct {
	ID          int     `json:"ID"`
	CompanyID   int     `json:"companyID"`
	Name        string  `json:"name"`
	WorkField   string  `json:"workField"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Date        string  `json:"date"`
	CompanyName string  `json:"companyName"`
}

func (cw *Work) Validate() error {
	if cw.Name == "" {
		return errors.New("missing name")
	}
	if cw.WorkField == "" {
		return errors.New("missing work field")
	}
	if cw.Description == "" {
		return errors.New("missing description")
	}
	if cw.Price < 0 {
		return errors.New("invalid price")
	}

	return nil
}
