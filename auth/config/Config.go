package config

import (
	"github.com/tkanos/gonfig"
	"sync"
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
	App struct {
		SchemaUp string
	}
	Redis struct {
		Host string
		Port string
	}
}

var (
	cfg  *Config
	once sync.Once
)

//GetConf singleton implementation.
func GetConf() *Config {
	once.Do(func() {
		cfg = parseConf()
		if cfg == nil {
			panic("bad config")
		}
	})

	return cfg
}

func parseConf() *Config {
	cfg = &Config{}

	err := gonfig.GetConf(cfgPath, cfg)
	if err != nil {
		return nil
	}

	return cfg
}
