package models

type Login struct {
	Credential string `json:"credential"`
	Password   string `json:"password"`
	UUID       string `json:"uuid"`
}
