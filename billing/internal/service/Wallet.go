package service

import (
	"github.com/mrbelka12000/netfix/billing/internal/repository"
	"github.com/mrbelka12000/netfix/billing/models"
)

type srvWallet struct {
	repo *repository.Repository
}

func newWallet(repo *repository.Repository) *srvWallet {
	return &srvWallet{repo}
}

func (sw *srvWallet) Create(wallet *models.Wallet) error {
	return sw.repo.Create(wallet)
}

func (sw *srvWallet) GetWalletAmount(ownerID int) (float64, error) {
	return sw.repo.GetWalletAmount(ownerID)
}
