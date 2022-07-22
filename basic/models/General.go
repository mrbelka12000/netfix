package models

type General struct {
	ID        int     `json:"ID"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Email     string  `json:"email"`
	WorkField *string `json:"workField,omitempty"`
	Birth     *string `json:"birth,omitempty"`
	UUID      string  `json:"uuid"`
}
