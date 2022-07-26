package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/mrbelka12000/netfix/users/config"
	"github.com/mrbelka12000/netfix/users/internal/repository"
	"github.com/mrbelka12000/netfix/users/models"
	"github.com/mrbelka12000/netfix/users/redis"
	"github.com/mrbelka12000/netfix/users/tools"
	"github.com/segmentio/kafka-go"
)

func (d *Delivery) ConsumerForCompany(exit chan struct{}, wg *sync.WaitGroup) {
	cfg := config.GetConf()

	conf := kafka.ReaderConfig{
		Brokers:  []string{cfg.Kafka.Brokers},
		Topic:    cfg.Kafka.TopicCompany,
		MaxBytes: cfg.Kafka.MaxBytes,
	}

	reader := kafka.NewReader(conf)
	var finished bool
	go closeReader(reader, exit, wg, &finished)

	for !finished {

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

		conn := repository.GetConnection()
		tx, err := conn.Begin()
		if err != nil {
			log.Println("tx creation error: " + err.Error())
			continue
		}

		id, err := d.srv.Register(gen, tx)
		if err != nil {
			tx.Commit()
			log.Println("registration error: " + err.Error())
			continue
		}

		err = d.srv.RegisterCompany(&models.Company{ID: id, WorkField: gen.WorkField}, tx)
		if err != nil {
			tx.Commit()
			log.Println("registration error: " + err.Error())
			continue
		}

		tx.Commit()
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

func (d *Delivery) ConsumerForCustomer(exit chan struct{}, wg *sync.WaitGroup) {
	cfg := config.GetConf()
	conf := kafka.ReaderConfig{
		Brokers:  []string{cfg.Kafka.Brokers},
		Topic:    cfg.Kafka.TopicCustomer,
		MaxBytes: cfg.Kafka.MaxBytes,
	}

	reader := kafka.NewReader(conf)
	var finished bool
	go closeReader(reader, exit, wg, &finished)

	for !finished {
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

		conn := repository.GetConnection()
		tx, err := conn.Begin()
		if err != nil {
			log.Println("tx creation error: " + err.Error())
			continue
		}

		id, err := d.srv.Register(gen, tx)
		if err != nil {
			tx.Commit()
			log.Println("registration error: " + err.Error())
			continue
		}

		err = d.srv.RegisterCustomer(&models.Customer{ID: id, Birth: gen.Birth}, tx)
		if err != nil {
			tx.Commit()
			log.Println("registration error: " + err.Error())
			continue
		}

		tx.Commit()
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

func (d *Delivery) ConsumerForGetCompany(exit chan struct{}, wg *sync.WaitGroup) {
	cfg := config.GetConf()

	conf := kafka.ReaderConfig{
		Brokers:  []string{cfg.Kafka.Brokers},
		Topic:    cfg.Kafka.TopicGetCompany,
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

		gen := &models.General{}
		err = json.Unmarshal(m.Value, &gen)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		uuid := gen.UUID
		gen, err = d.srv.Company.GetByID(gen.ID)
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

func (d *Delivery) ConsumerForGetCustomer(exit chan struct{}, wg *sync.WaitGroup) {
	cfg := config.GetConf()

	conf := kafka.ReaderConfig{
		Brokers:  []string{cfg.Kafka.Brokers},
		Topic:    cfg.Kafka.TopicGetCustomer,
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

		gen := &models.General{}
		err = json.Unmarshal(m.Value, &gen)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		uuid := gen.UUID
		gen, err = d.srv.Customer.GetByID(gen.ID)
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

func (d *Delivery) ConsumerForLogin(exit chan struct{}, wg *sync.WaitGroup) {
	cfg := config.GetConf()

	conf := kafka.ReaderConfig{
		Brokers:  []string{cfg.Kafka.Brokers},
		Topic:    cfg.Kafka.TopicLogin,
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
		l := &models.Login{}

		err = json.Unmarshal(m.Value, &l)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		err = d.srv.Login(l)
		if err != nil {
			log.Println("no such user: " + err.Error())
			continue
		}

		role := &models.Role{ID: l.ID, UserType: l.UserType}

		err = redis.SetValue(l.UUID, tools.MakeJsonString(role))
		if err != nil {
			log.Println("session create error: " + err.Error())
			continue
		}
		err = publish(tools.MakeJsonString(l), cfg.Kafka.TopicLoginResp)
		if err != nil {
			log.Println("publish error: " + err.Error())
			continue
		}
		log.Println("login successfully")
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
