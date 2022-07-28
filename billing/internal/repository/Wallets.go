package repository

import (
	"log"

	"github.com/mrbelka12000/netfix/billing/models"
)

type repoWallet struct {
}

func newWallet() *repoWallet {
	return &repoWallet{}
}

const unlimAmount = 10000

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

func (rw *repoWallet) GetWalletAmount(ownerId int) (float64, error) {
	conn := GetConnection()
	var amount float64
	err := conn.QueryRow(`
		SELECT 
		    amount
		FROM
		    wallets
		WHERE 
		    OwnerID=$1
`, ownerId).Scan(&amount)
	if err != nil {
		log.Println("get wallet error: " + err.Error())
		return 0, err
	}

	return amount, nil
}
