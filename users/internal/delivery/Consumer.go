package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mrbelka12000/netfix/users/config"
	"github.com/mrbelka12000/netfix/users/models"
	"github.com/mrbelka12000/netfix/users/redis"
	"github.com/mrbelka12000/netfix/users/tools"
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

		role := models.Role{ID: id, UserType: models.Cmp}
		err = redis.SetValue(gen.UUID, tools.MakeJsonString(role))
		if err != nil {
			log.Println("may be panic? :" + err.Error())
			continue
		}

		gen.ID = id
		err = publish(tools.MakeJsonString(gen), cfg.Kafka.TopicAuth)
		if err != nil {
			log.Println(err.Error())
			continue
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

		log.Println("successfully created")

		role := models.Role{ID: id, UserType: models.Cust}
		err = redis.SetValue(gen.UUID, tools.MakeJsonString(role))
		if err != nil {
			log.Println("may be panic? :" + err.Error())
			continue
		}

		gen.ID = id
		err = publish(tools.MakeJsonString(gen), cfg.Kafka.TopicAuth)
		if err != nil {
			log.Println(err.Error())
			continue
		}
	}
}

func (d *Delivery) ConsumerForGetCompany() {
	cfg := config.GetConf()

	conf := kafka.ReaderConfig{
		Brokers:  []string{cfg.Kafka.Brokers},
		Topic:    cfg.Kafka.TopicGetCompany,
		MaxBytes: cfg.Kafka.MaxBytes,
	}

	reader := kafka.NewReader(conf)

	for {

		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error while reading from consumer: ", err)
			continue
		}

		gen := &models.General{}
		err = json.Unmarshal(m.Value, &gen)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		uuid := gen.UUID
		gen, err = d.srv.GetByID(gen.ID)
		if err != nil {
			log.Println("get company by id error: " + err.Error())
			continue
		}
		gen.UUID = uuid

		err = publish(tools.MakeJsonString(gen), cfg.Kafka.TopicUserGetResp)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
