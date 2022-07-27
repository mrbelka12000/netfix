package delivery

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mrbelka12000/netfix/basic/config"
	"github.com/mrbelka12000/netfix/basic/models"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func Consumer(topic, uuid string) (*models.General, error) {
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
			return nil, errors.New("лимит ожидания закончился")
		}

		if len(m.Value) != 0 {
			gen := &models.General{}
			fmt.Printf("Message from %v is %v uuid = %v\n", topic, string(m.Value), uuid)
			err = json.Unmarshal(m.Value, &gen)
			if err != nil {
				log.Println(err.Error())
				return nil, err
			}
			if gen.UUID != uuid {
				log.Println("uuid does not match")
				continue
			}
			gen.UUID = ""
			return gen, nil
		}
		fmt.Println("No message from kafka")
		return nil, errors.New("пришло не известно что")
	}
}
