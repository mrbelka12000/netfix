package models

import "errors"

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

func (g *General) Validate() error {
	if g.Username == "" {
		return errors.New("missing username")
	}
	if g.Password == "" {
		return errors.New("missing password")
	}
	if g.Email == "" {
		return errors.New("missing email")
	}
	if g.WorkField != nil && *g.WorkField == "" {
		return errors.New("missing work field")
	}
	if g.Birth != nil && *g.Birth == "" {
		return errors.New("missing date of birth")
	}
	return nil
}
