package service

import (
	"github.com/mrbelka12000/netfix/billing/internal/repository"
	"github.com/mrbelka12000/netfix/billing/models"
)

type Wallet interface {
	Create(wallet *models.Wallet) error
	GetWalletAmount(ownerID int) (float64, error)
}

type Billing interface {
}

type Service struct {
	Wallet
	Billing
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Wallet: newWallet(repo),
	}
}
