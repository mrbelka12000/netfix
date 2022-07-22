package config

import (
	"github.com/tkanos/gonfig"
	"sync"
)

const cfgPath = "basic/config/config.json"

type config struct {
	App struct {
		Port string
	}
	Redis struct {
		Host string
		Prot string
	}
	Kafka struct {
		TopicCompany  string
		TopicCustomer string
		TopicAuth     string
		Brokers       string
		RetryMax      int
		RequiredAcks  int
		MaxBytes      int
		Successes     bool
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
