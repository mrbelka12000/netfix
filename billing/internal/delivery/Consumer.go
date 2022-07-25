package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mrbelka12000/netfix/billing/config"
	"github.com/mrbelka12000/netfix/billing/models"
	"github.com/mrbelka12000/netfix/billing/tools"
	"github.com/segmentio/kafka-go"
	"log"
)

func (d *Delivery) ConsumerForWallets() {
	cfg := config.GetConf()

	conf := kafka.ReaderConfig{
		Brokers:  []string{cfg.Kafka.Brokers},
		Topic:    cfg.Kafka.TopicWallets,
		MaxBytes: cfg.Kafka.MaxBytes,
	}

	reader := kafka.NewReader(conf)

	for {

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
		fmt.Printf("%+v\n", w)
		err = d.srv.Create(w)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		gen := &models.General{UUID: w.UUID, ID: w.OwnerID}

		publish(tools.MakeJsonString(gen), cfg.Kafka.TopicCreateWallet)
	}
}

func (d *Delivery) ConsumerForBilling() {
	cfg := config.GetConf()

	conf := kafka.ReaderConfig{
		Brokers:  []string{cfg.Kafka.Brokers},
		Topic:    cfg.Kafka.TopicBilling,
		MaxBytes: cfg.Kafka.MaxBytes,
	}

	reader := kafka.NewReader(conf)

	for {

		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error while reading from consumer: ", err)
			continue
		}
		fmt.Println(string(m.Value))
	}
}
