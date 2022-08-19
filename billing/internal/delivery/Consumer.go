package delivery

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/mrbelka12000/netfix/billing/config"
	"github.com/mrbelka12000/netfix/billing/models"
	"github.com/mrbelka12000/netfix/billing/tools"
	"github.com/segmentio/kafka-go"
)

func (d *Delivery) ConsumerForWallets(exit chan struct{}, wg *sync.WaitGroup) {
	cfg := config.GetConf()

	conf := kafka.ReaderConfig{
		Brokers:  []string{cfg.Kafka.Brokers},
		Topic:    cfg.Kafka.TopicWallets,
		MaxBytes: cfg.Kafka.MaxBytes,
	}

	reader := kafka.NewReader(conf)

	var finished bool
	go closeReader(reader, exit, wg, &finished)

	for !finished {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error while reading from consumer: ", err)
			continue
		}

		w := &models.Wallet{}
		err = json.Unmarshal(m.Value, &w)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		err = d.srv.Create(w)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		gen := &models.General{UUID: w.UUID, ID: w.OwnerID}

		err = publish(tools.MakeJsonString(gen), cfg.Kafka.TopicCreateWallet)
		if err != nil {
			log.Println("publish error: " + err.Error())
			continue
		}
		log.Println("wallet successfully created")
	}
}

func (d *Delivery) ConsumerForBilling(bil chan<- []byte, exit chan struct{}, wg *sync.WaitGroup) {
	cfg := config.GetConf()

	conf := kafka.ReaderConfig{
		Brokers:  []string{cfg.Kafka.Brokers},
		Topic:    cfg.Kafka.TopicBilling,
		MaxBytes: cfg.Kafka.MaxBytes,
	}

	reader := kafka.NewReader(conf)

	var finished bool
	go closeReader(reader, exit, wg, &finished)

	for !finished {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error while reading from consumer: ", err)
			continue
		}
		ap := &models.Apply{}

		//minimal json validation
		err = json.Unmarshal(m.Value, &ap)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		err = ap.Validate()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		bil <- m.Value
	}
}

func (d *Delivery) ConsumerForGetWallet(exit chan struct{}, wg *sync.WaitGroup) {
	cfg := config.GetConf()

	conf := kafka.ReaderConfig{
		Brokers:  []string{cfg.Kafka.Brokers},
		Topic:    cfg.Kafka.TopicGetWallet,
		MaxBytes: cfg.Kafka.MaxBytes,
	}

	reader := kafka.NewReader(conf)

	var finished bool
	go closeReader(reader, exit, wg, &finished)

	for !finished {

		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error while reading from consumer: ", err)
			continue
		}
		g := &models.General{}

		err = json.Unmarshal(m.Value, &g)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		amount, err := d.srv.GetWalletAmount(g.ID)
		if err != nil {
			log.Println("get wallet by ownerID error: " + err.Error())
			continue
		}
		g.Amount = amount

		err = publish(tools.MakeJsonString(g), cfg.Kafka.TopicGetWalletResp)
		if err != nil {
			log.Println("publish error: " + err.Error())
			continue
		}
		log.Println("response successfully created")
	}
}

func closeReader(reader *kafka.Reader, exit chan struct{}, wg *sync.WaitGroup, finished *bool) {
	defer func() {
		exit <- struct{}{}
		finished = tools.PtrBool(true)
		wg.Done()
	}()

	<-exit

	err := reader.Close()
	if err != nil {
		log.Printf("kafka reader close error: %v \n", err)
	} else {
		log.Println("kafka reader closed")
	}
}
