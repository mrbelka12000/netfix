package models

type Session struct {
	ID     int    `json:"id"`
	Cookie string `json:"cookie"`
}
