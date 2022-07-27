package config

import (
	"sync"

	"github.com/tkanos/gonfig"
)

const cfgPath = "basic/config/config.json"

type config struct {
	App struct {
		Port       string
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
		Port string
	}
	Kafka struct {
		TopicCompany       string
		TopicCustomer      string
		TopicAuth          string
		TopicGetCompany    string
		TopicGetCustomer   string
		TopicUserGetResp   string
		TopicGetWallet     string
		TopicGetWalletResp string
		TopicBilling       string
		TopicWallets       string
		TopicCreateWallet  string
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
