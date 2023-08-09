package RedisGo

import (
	"encoding/json"
	"time"
)

func (r RedisInstance) WriteJSONToRedisWithTTL(key string, value interface{}, timeFrame time.Duration) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.Set(key, string(jsonData), timeFrame).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r RedisInstance) WriteJSONToRedis(key string, value interface{}) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.Set(key, string(jsonData), 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// read json as a struct from redis
// set timeFrame to 0 if you don't want to update the ttl
func (r RedisInstance) ReadJSONFromRedisWithTTL(key string) (string, time.Duration, error) {
	val, err := r.Get(key).Result()

	if err != nil {
		return "", 0, err
	}

	ttl, err := r.TTL(key).Result()
	if err != nil {
		return "", 0, err
	}

	return val, ttl, nil

}

func (r RedisInstance) ReadJSONFromRedis(key string) (string, error) {
	val, err := r.Get(key).Result()

	if err != nil {
		return "", err
	}
	return val, nil
}
