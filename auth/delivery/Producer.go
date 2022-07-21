package delivery

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mrbelka12000/netfix/auth/config"

	"github.com/Shopify/sarama"
)

func (d *Delivery) Produce(cfg *config.Config) {
	// create producer
	producer, err := initProducer(cfg)
	if err != nil {
		fmt.Println("Error producer: ", err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter msg: ")
		msg, _ := reader.ReadString('\n')

		// publish without goroutene
		publish(msg, producer, cfg)

		// publish with go routene
		// go publish(msg, producer)
	}
}

func initProducer(cfg *config.Config) (sarama.SyncProducer, error) {
	// setup sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

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

func publish(message string, producer sarama.SyncProducer, cfg *config.Config) {
	// publish sync
	msg := &sarama.ProducerMessage{
		Topic: cfg.Kafka.TopicCompany,
		Value: sarama.StringEncoder(message),
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}

	// publish async
	// producer.Input() <- &sarama.ProducerMessage{

	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
}
