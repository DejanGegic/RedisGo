package controller

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func ConnectToRedisClient(db int) *redis.Client {

	host := "localhost"
	log.Println("Connecting to redis at", host)
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":6379",
		Password: "",
		DB:       db,
	})
	return client
}

func SetEvictionPolicy(client *redis.Client, key string, policy string) error {
	err := client.ConfigSet("maxmemory-policy", policy).Err()
	if err != nil {
		return err
	}
	return err
}

func WriteToRedisTimeout(client *redis.Client, key string, value string, timeFrame time.Duration) error {
	err := client.Set(key, value, timeFrame).Err()
	if err != nil {
		return err
	}
	return err
}

func IncrementValue(client *redis.Client, key string) error {
	incr := client.Incr(key)
	return incr.Err()
}
func WriteToRedisHash(client *redis.Client, key string, field string, value string) error {
	err := client.HSet(key, field, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func ReadFromRedis(client *redis.Client, key string) (int, error) {

	val, err := client.Get(key).Result()
	if err != nil {
		val = "0"
	}
	// convert the string to an int
	intVal, err := strconv.Atoi(val)
	return intVal, err
}

func ReadFromRedisWithTTL(client *redis.Client, key string) (int, time.Duration, error) {
	val, err := client.Get(key).Result()
	if err != nil {
		val = "0"
	}
	// convert the string to an int
	ttl, err := client.TTL(key).Result()
	intVal, err := strconv.Atoi(val)
	return intVal, ttl, err
}

func WriteJSONToRedis(client *redis.Client, key string, value interface{}, timeFrame time.Duration) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = client.Set(key, string(jsonData), timeFrame).Err()
	if err != nil {
		return err
	}
	return nil
}

// read json as a struct from redis
// set timeFrame to 0 if you don't want to update the ttl
func ReadJSONFromRedis(client *redis.Client, key string, timeFrame time.Duration) (string, error) {
	val, err := client.Get(key).Result()

	if err != nil {
		return "", err
	}
	// set the ttl if timeFrame not nil
	if timeFrame != 0 {
		go client.Expire(key, timeFrame)
	}
	return val, nil
}

func FlushAll(client *redis.Client) error {
	err := client.FlushDB().Err()
	if err != nil {
		return err
	}
	return err
}
