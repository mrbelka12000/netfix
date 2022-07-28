package models

type Login struct {
	ID         int    `json:"ID"`
	Credential string `json:"credential"`
	Password   string `json:"password"`
	UserType   string
	UUID       string `json:"uuid"`
}
