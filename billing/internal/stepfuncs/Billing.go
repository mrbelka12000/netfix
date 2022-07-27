package stepfuncs

import (
	"encoding/json"
	"github.com/mrbelka12000/netfix/billing/internal/repository"
	"github.com/mrbelka12000/netfix/billing/models"
	"log"
	"math"
	"time"
)

func Billing(ch <-chan []byte, exit chan struct{}) {
	for {
		select {
		case applyJson := <-ch:
			log.Println("billing starting")
			go billingOperations(applyJson, exit)
		case <-exit:
			panic("step func billing exited")
		}
	}
}

func billingOperations(applyJson []byte, exit chan<- struct{}) {

	conn := repository.GetConnection()

	ap := &models.Apply{}
	err := json.Unmarshal(applyJson, &ap)
	if err != nil {
		log.Println(err.Error())
		exit <- struct{}{}
		return
	}

	tx, err := conn.Begin()
	if err != nil {
		log.Println("tx creation error: " + err.Error())
		exit <- struct{}{}
		return
	}
	defer tx.Commit()

	totalPrice := calculateWorkAmount(ap)
	_, err = tx.Exec(`
		INSERT INTO billing
			(applyid, amount)
		VALUES 
			($1,$2)
`, ap.ID, totalPrice)
	if err != nil {
		log.Println("create billing error: " + err.Error())
		tx.Rollback()
		exit <- struct{}{}
		return
	}

	_, err = tx.Exec(`
		UPDATE 
		    wallets
		SET 
		    amount = amount-$1
		where 
		    ownerid = $2
`, totalPrice, ap.CustomerID)
	if err != nil {
		log.Println("charge-off from customer error: " + err.Error())
		tx.Rollback()
		exit <- struct{}{}
		return
	}

	_, err = tx.Exec(`
		UPDATE 
		    wallets
		SET 
		    amount = amount+$1
		where 
		    ownerid = $2
`, totalPrice, ap.CompanyID)
	if err != nil {
		log.Println("cash deposit to company error: " + err.Error())
		tx.Rollback()
		exit <- struct{}{}
		return
	}
	log.Println("billing successfully finished")
}

func calculateWorkAmount(ap *models.Apply) int64 {
	startTime := time.Unix(ap.StartDate, 0)
	endTime := time.Unix(ap.EndDate, 0)
	diff := startTime.Sub(endTime)

	hours := diff.Hours()
	if hours < 0 {
		hours *= -1
	}
	workPrice := hours * ap.Price
	totalPrice := math.Floor(workPrice)
	return int64(totalPrice)
}
