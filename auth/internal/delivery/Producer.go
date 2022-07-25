package delivery

import (
	"fmt"
	"log"
	"time"

	"github.com/mrbelka12000/netfix/auth/config"

	"github.com/Shopify/sarama"
)

func (d *Delivery) Produce() {
	cfg := config.GetConf()

	i := 0
	for {
		msg := `
{
"email":"check",
"username":"dom",
"password":"12345",
"workField":"zaebka",
"birth":"2002-02-15"
}
`
		publish(msg, cfg.Kafka.TopicCustomer)
		time.Sleep(5 * time.Second)
		i++
	}
}

func initProducer() (sarama.SyncProducer, error) {
	// setup sarama log to stdout
	// sarama.Logger = log.New(os.Stdout, "", log.Ltime)
	cfg := config.GetConf()
	// producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = cfg.Kafka.RetryMax
	config.Producer.RequiredAcks = sarama.RequiredAcks(cfg.Kafka.RequiredAcks)
	config.Producer.Return.Successes = cfg.Kafka.Successes

	// async producer
	// prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{cfg.Kafka.Brokers}, config)

	return prd, err
}

func publish(message string, topic string) error {
	producer, err := initProducer()
	if err != nil {
		log.Println("producer error: " + err.Error())
		return err
	}
	// publish sync
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		log.Println("Error publish: ", err.Error())
		return err
	}

	// publish async
	// producer.Input() <- &sarama.ProducerMessage{

	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
	return nil
}
