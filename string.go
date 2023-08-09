package RedisGo

import (
	"strconv"
	"time"
)

func (r RedisInstance) WriteToRedisString(key string, value string) error {
	err := r.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return err
}

func (r RedisInstance) WriteToRedisStringWithTTL(key string, value string, timeFrame time.Duration) error {
	err := r.Set(key, value, timeFrame).Err()
	if err != nil {
		return err
	}
	return err
}

func (r RedisInstance) ReadFromRedisString(key string) (int, error) {

	val, err := r.Get(key).Result()
	if err != nil {
		val = "0"
	}
	// convert the string to an int
	intVal, err := strconv.Atoi(val)
	return intVal, err
}

func (r RedisInstance) ReadFromRedisStringWithTTL(key string) (int, time.Duration, error) {
	val, err := r.Get(key).Result()
	if err != nil {
		val = "0"
	}
	// convert the string to an int
	ttl, err := r.TTL(key).Result()
	intVal, err := strconv.Atoi(val)
	return intVal, ttl, err
}
