package models

const (
	Cmp  = "Company"
	Cust = "Customer"
)

type Role struct {
	ID       int    `json:"ID"`
	UserType string `json:"userType"`
}
