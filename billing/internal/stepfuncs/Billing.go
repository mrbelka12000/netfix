package stepfuncs

import (
	"encoding/json"
	"log"
	"math"
	"sync"
	"time"

	"github.com/mrbelka12000/netfix/billing/internal/repository"
	"github.com/mrbelka12000/netfix/billing/models"
	"github.com/mrbelka12000/netfix/billing/tools"
)

func Billing(bil <-chan []byte, exit chan struct{}, wg *sync.WaitGroup) {
	defer func() { wg.Done() }()

	for {
		var finished bool
		var started bool
		select {
		case applyJson := <-bil:
			log.Println("billing starting")
			go billingOperations(applyJson, &started, &finished)
		case <-exit:
			if !started {
				log.Println("billing down")
				break
			}
			if !finished && started {
				time.Sleep(5 * time.Second)
				if finished {
					log.Println("billing down")
					break
				} else {
					panic("AAAAAAAAAAAA")
				}
			}
		}
	}
}

func billingOperations(applyJson []byte, started, finished *bool) {

	started = tools.PtrBool(true)
	conn := repository.GetConnection()

	ap := &models.Apply{}
	err := json.Unmarshal(applyJson, &ap)
	if err != nil {
		log.Println(err.Error())
		return
	}

	tx, err := conn.Begin()
	if err != nil {
		log.Println("tx creation error: " + err.Error())
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
		return
	}

	finished = tools.PtrBool(true)
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
