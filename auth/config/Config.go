package config

import (
	"github.com/tkanos/gonfig"
)

const cfgPath = "auth/config/config.json"

type Config struct {
	Postgres struct {
		POSTGRES_DB       string
		POSTGRES_HOST     string
		POSTGRES_PASSWORD string
		POSTGRES_PORT     string
		POSTGRES_USER     string
	}
	Kafka struct {
		TopicCompany  string
		TopicCustomer string
		TopicLogin    string
		Brokers       string
		RetryMax      int
		RequiredAcks  int
		MaxBytes      int
		Successes     bool
	}
}

var (
	cfg    *Config
	exists bool
)

func GetConf() *Config {
	if exists {
		return cfg
	}

	cfg = parseConf()
	if !exists {
		panic("bad config")
	}
	return cfg
}

func parseConf() *Config {
	cfg = &Config{}

	err := gonfig.GetConf(cfgPath, cfg)
	if err != nil {
		return nil
	}

	exists = true
	return cfg
}
