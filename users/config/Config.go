package config

import (
	"github.com/tkanos/gonfig"
	"sync"
)

const cfgPath = "users/config/config.json"

type config struct {
	Postgres struct {
		POSTGRES_DB       string
		POSTGRES_HOST     string
		POSTGRES_PASSWORD string
		POSTGRES_PORT     string
		POSTGRES_USER     string
	}
	Kafka struct {
		TopicCompany     string
		TopicCustomer    string
		TopicAuth        string
		TopicGetCompany  string
		TopicGetCustomer string
		TopicUserGetResp string
		Brokers          string
		RetryMax         int
		RequiredAcks     int
		MaxBytes         int
		Successes        bool
	}
	App struct {
		SchemaUp string
	}
	Redis struct {
		Host string
		Port string
	}
}

var (
	cfg  *config
	once sync.Once
)

//GetConf singleton implementation.
func GetConf() *config {
	once.Do(func() {
		cfg = parseConf()
		if cfg == nil {
			panic("bad config")
		}
	})

	return cfg
}

func parseConf() *config {
	cfg = &config{}

	err := gonfig.GetConf(cfgPath, cfg)
	if err != nil {
		return nil
	}

	return cfg
}
