package models

type Wallet struct {
	ID      int    `json:"ID"`
	OwnerID int    `json:"ownerID"`
	Amount  int    `json:"amount"`
	UUID    string `json:"uuid"`
}
