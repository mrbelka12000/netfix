package delivery

import (
	"log"

	"github.com/Shopify/sarama"
	"github.com/mrbelka12000/netfix/basic/config"
)

func initProducer() (sarama.SyncProducer, error) {
	// setup sarama log to stdout
	//sarama.Logger = log.New(os.Stdout, "", log.Ltime)
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

func Publish(message, topic string) error {
	producer, err := initProducer()
	if err != nil {
		log.Println("error producer: " + err.Error())
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	p, o, err := producer.SendMessage(msg)
	if err != nil {
		log.Println("send message error: " + err.Error())
		return err
	}

	log.Print("Message: ", message)
	log.Println("Partition: ", p)
	log.Println("Offset: ", o)
	return nil
}
