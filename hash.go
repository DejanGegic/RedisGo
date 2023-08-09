package RedisGo

import "time"

func (r RedisInstance) WriteToRedisHash(key string, field string, value string) error {
	err := r.HSet(key, field, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r RedisInstance) WriteToRedisHashWithTTL(key string, field string, value string, ttl time.Duration) error {
	// make hash set with ttl
	err := r.HSet(key, field, value).Err()
	if err != nil {
		return err
	}
	// set the ttl
	go r.Expire(key, ttl)
	return nil
}

func (r RedisInstance) ReadFromRedisHash(key string, field string) (string, error) {
	val, err := r.HGet(key, field).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r RedisInstance) ReadFromRedisHashWithTTL(key string, field string) (string, time.Duration, error) {
	val, err := r.HGet(key, field).Result()
	if err != nil {
		return "", 0, err
	}
	// get the ttl
	ttlVal, err := r.TTL(key).Result()
	if err != nil {
		return "", 0, err
	}
	return val, ttlVal, nil
}
