package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mrbelka12000/netfix/auth/config"
	"github.com/mrbelka12000/netfix/auth/models"
	"github.com/segmentio/kafka-go"
	"log"
)

func (d *Delivery) ConsumerForCompany() {
	cfg := config.GetConf()

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

		fmt.Printf("Message for company is : %v \n", string(m.Value))
		gen := &models.General{}
		if err = json.Unmarshal(m.Value, &gen); err != nil {
			log.Println("unmarshall error: " + err.Error())
			continue
		}
		id, err := d.srv.Register(gen)
		if err != nil {
			log.Println("registration error: " + err.Error())
			continue
		}
		err = d.srv.RegisterCompany(&models.Company{ID: id, WorkField: gen.WorkField})
		if err != nil {
			log.Println("registration error: " + err.Error())
			continue
		}

		log.Println("successfully created")
		err = publish(gen.UUID, cfg.Kafka.TopicAuth)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func (d *Delivery) ConsumerForCustomer() {
	cfg := config.GetConf()
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

		fmt.Printf("Message for customer is : %v \n", string(m.Value))

		gen := &models.General{}
		if err = json.Unmarshal(m.Value, &gen); err != nil {
			log.Println("unmarshall error: " + err.Error())
			continue
		}
		id, err := d.srv.Register(gen)
		if err != nil {
			log.Println("registration error: " + err.Error())
			continue
		}

		err = d.srv.RegisterCustomer(&models.Customer{ID: id, Birth: gen.Birth})
		if err != nil {
			log.Println("registration error: " + err.Error())
			continue
		}

	}
}
