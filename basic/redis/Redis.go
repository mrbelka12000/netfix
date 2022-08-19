package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/mrbelka12000/netfix/basic/config"
	"github.com/mrbelka12000/netfix/basic/models"
)

func SetValue(key, value string) error {
	c := getClient()
	if c == nil {
		log.Println("не удалось получить клиента для редис")
		return errors.New("не удалось получить клиента для редис")
	}

	err := c.Set(key, value, 0).Err()
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
	val, err := c.Get(key).Result()
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

func GetUserType(session string) (*models.Role, error) {
	ut := &models.Role{}

	jsonB, err := GetValue(session)
	if err != nil {
		log.Println("no value in redis: " + err.Error())
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonB), &ut)
	if err != nil {
		log.Println("unmarshall error: " + err.Error())
		return nil, err
	}

	return ut, nil
}

func CloseRedis() error {
	c := getClient()
	return c.Close()
}
