package redis

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/mrbelka12000/netfix/users/config"
)

func SetValue(key, value string) error {
	c := getClient()
	if c == nil {
		log.Println("не удалось получить клиента для редис")
		return errors.New("не удалось получить клиента для редис")
	}

	err := c.Set(key, value, 3600*time.Second).Err()
	if err != nil {
		log.Println("сan not set value: " + err.Error())
		return err
	}

	return nil
}

func GetValue(key string) (string, error) {
	c := getClient()
	if c == nil {
		log.Println("не удалось получить клиента для редис")
		return "", errors.New("не удалось получить клиента для редис")
	}
	val, err := c.Get("ping").Result()
	if err != nil {
		log.Println("can not get value: " + err.Error())
		return "", err
	}
	return val, nil
}

func getClient() *redis.Client {
	cfg := config.GetConf()

	addr := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return client
}
