package delivery

import (
	"context"
	"errors"
	"fmt"
	"github.com/mrbelka12000/netfix/basic/config"
	"github.com/segmentio/kafka-go"
	"time"
)

func Consumer(topic, uuid string) error {
	cfg := config.GetConf()

	conf := kafka.ReaderConfig{
		Brokers:  []string{cfg.Kafka.Brokers},
		Topic:    topic,
		MaxBytes: cfg.Kafka.MaxBytes,
	}

	reader := kafka.NewReader(conf)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			fmt.Println("Error while reading from consumer: ", err)
			return errors.New("лимит ожидания закончился")
		}

		fmt.Println(len(m.Value))
		if len(m.Value) != 0 {
			fmt.Printf("Message from %v is %v, uuid = %v\n", topic, string(m.Value), uuid)
			if string(m.Value) != uuid {
				continue
			}
			return nil
		}
		fmt.Println("No message from kafka")
		return errors.New("пришло не известно что")
	}
}
