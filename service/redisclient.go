package service

import (
	"log"

	"github.com/bhupeshbhatia/go-authserver/models"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

//RedisClient connects to Redis server to store tokens
func RedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if client == nil {
		return nil, errors.New("No connection")
	}
	return client, nil
}

//SetToken sets the token for redis client
func SetToken(key string, val *models.RefreshToken) (string, error) {
	client, err := RedisClient()
	if err != nil {
		return "", err
	}

	err = client.Set(key, val, 0).Err()
	if err != nil {
		return "", err
	}
	return "Token added", nil
}

//GetToken retrieves token from redis db
func GetToken(key string) (string, error) {
	client, err := RedisClient()

	if err != nil {
		return "", err
	}

	val, err := client.Get(key).Result()
	if err == redis.Nil {
		err = errors.Wrap(err, "Key does not exist")
		log.Println(err)
		return "", errors.New("Key not found")
	}
	return val, nil
}
