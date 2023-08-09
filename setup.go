package RedisGo

import (
	"log"

	"github.com/go-redis/redis"
)

func ConnectToRedisClient(db int) RedisInstance {

	host := "localhost"
	log.Println("Connecting to redis at", host)
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":6379",
		Password: "",
		DB:       db,
	})

	instance := RedisInstance{
		client,
	}
	return instance
}

func (r RedisInstance) SetEvictionPolicy(policy string) error {
	err := r.ConfigSet("maxmemory-policy", policy).Err()
	if err != nil {
		return err
	}
	return err
}
