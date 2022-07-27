package repository

import "github.com/mrbelka12000/netfix/billing/models"

type Wallet interface {
	Create(wallet *models.Wallet) error
	GetWalletAmount(ownerID int) (float64, error)
}

type Billing interface {
}

type Repository struct {
	Wallet
	Billing
}

func NewRepo() *Repository {
	return &Repository{
		Wallet: newWallet(),
	}
}
