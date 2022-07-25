package repository

import (
	"github.com/mrbelka12000/netfix/billing/models"
	"log"
)

type repoWallet struct {
}

func newWallet() *repoWallet {
	return &repoWallet{}
}

const unlimAmount = 1423821348218124

func (rw *repoWallet) Create(wallet *models.Wallet) error {

	conn := GetConnection()

	err := conn.QueryRow(`
	INSERT INTO wallets
		(ownerId, amount)
	VALUES 
		($1,$2)
	RETURNING 
		ID
`, wallet.OwnerID, unlimAmount).Scan(&wallet.ID)
	if err != nil {
		log.Println("create wallet error: " + err.Error())
		return err
	}

	log.Println("wallet successfully created")
	return nil
}
