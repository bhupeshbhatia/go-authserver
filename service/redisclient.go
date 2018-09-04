package service

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisStore interface {
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(key string) *redis.StringCmd
}

// type Redis struct {
// 	r RedisStore
// }

// //RedisClient connects to Redis server to store tokens
// func (rc *Redis) RedisClient(add string, password string, db int) (*redis.Client, error) {

// 	client := redis.NewClient(&redis.Options{
// 		//"localhost:6379"
// 		//""
// 		//0
// 		Addr:     add,
// 		Password: password, // no password set
// 		DB:       db,       // use default DB
// 	})

// 	if client == nil {
// 		return nil, errors.New("No connection")
// 	}
// 	return client, nil
// }

// //SetToken sets the token for redis client
// func (rc *Redis) SetToken(key string, val *models.RefreshToken, client *redis.Client) error {

// 	err := client.Set(key, val, 0).Err()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// //GetToken retrieves token from redis db
// func (rc *Redis) GetToken(key string, client *redis.Client) (string, error) {
// 	// client, err := rc.r.RedisClient()

// 	// if err != nil {
// 	// 	return "", err
// 	// }

// 	val, err := client.Get(key).Result()
// 	if err == redis.Nil {
// 		err = errors.Wrap(err, "Key does not exist")
// 		log.Println(err)
// 		return "", errors.New("Key not found")
// 	}
// 	return val, nil
// }
