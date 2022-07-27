package models

type General struct {
	ID        int     `json:"ID"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Email     string  `json:"email"`
	WorkField *string `json:"workField,omitempty"`
	Birth     *string `json:"birth,omitempty"`
	Amount    float64 `json:"amount"`
	UUID      string  `json:"uuid"`
}
