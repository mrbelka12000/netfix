package models

type General struct {
	ID        int    `json:"ID"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	WorkField string `json:"workField"`
	Birth     string `json:"birth"`
	UUID      string `json:"uuid"`
}
