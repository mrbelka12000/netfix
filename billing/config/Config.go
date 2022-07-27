package config

import (
	"sync"

	"github.com/tkanos/gonfig"
)

const cfgPath = "billing/config/config.json"

type config struct {
	App struct {
		SchemaUp   string
		SchemaDown string
	}
	Postgres struct {
		POSTGRES_DB       string
		POSTGRES_HOST     string
		POSTGRES_PASSWORD string
		POSTGRES_PORT     string
		POSTGRES_USER     string
	}
	Redis struct {
		Host string
		Prot string
	}
	Kafka struct {
		TopicBilling       string
		TopicCreateWallet  string
		TopicWallets       string
		TopicGetWallet     string
		TopicGetWalletResp string
		Brokers            string
		RetryMax           int
		RequiredAcks       int
		MaxBytes           int
		Successes          bool
	}
}

var (
	cfg  *config
	once sync.Once
)

// GetConf singleton implementation.
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
