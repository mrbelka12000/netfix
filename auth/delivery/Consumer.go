package delivery

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"

	"github.com/mrbelka12000/netfix/auth/config"
	"github.com/mrbelka12000/netfix/auth/models"
)

func (d *Delivery) ConsumerForCompany(cfg *config.Config) {
	conf := kafka.ReaderConfig{
		Brokers:  []string{cfg.Kafka.Brokers},
		Topic:    cfg.Kafka.TopicCompany,
		MaxBytes: cfg.Kafka.MaxBytes,
	}

	reader := kafka.NewReader(conf)

	for {

		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error while reading from consumer: ", err)
			continue
		}

		fmt.Print("Message for company is : ", string(m.Value))
		gen := &models.General{}

		if err := json.Unmarshal(m.Value, &gen); err == nil {
			_, err = d.srv.Register(gen, models.Cmp)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Println(err.Error())
		}
	}
}

func (d *Delivery) ConsumerForCustomer(cfg *config.Config) {
	conf := kafka.ReaderConfig{
		Brokers:  []string{cfg.Kafka.Brokers},
		Topic:    cfg.Kafka.TopicCustomer,
		MaxBytes: cfg.Kafka.MaxBytes,
	}

	reader := kafka.NewReader(conf)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error while reading from consumer: ", err)
			continue
		}
		fmt.Println("Message is : ", string(m.Value))
	}
}
