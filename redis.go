package RedisGo

import (
	"strconv"

	"github.com/go-redis/redis"
)

type RedisInstance struct {
	*redis.Client
}

func (r RedisInstance) IncrementValue(key int) error {
	err := r.Incr(strconv.Itoa(key))
	return err.Err()
}
