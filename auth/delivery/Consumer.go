package delivery

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"

	"github.com/mrbelka12000/netfix/auth/config"
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
