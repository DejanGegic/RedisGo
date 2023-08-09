package RedisGo_test

import (
	"testing"
	"time"
)

func TestWriteToRedisString(t *testing.T) {

	R.FlushDB()

	err := R.WriteToRedisString("test", "1")

	if err != nil {
		t.Fail()
	}

}

func TestWriteToRedisStringWithTTL(t *testing.T) {

	R.FlushDB()

	err := R.WriteToRedisStringWithTTL("test", "1", time.Second*10)

	if err != nil {
		t.Fail()
	}

}

func TestReadFromRedisString(t *testing.T) {

	R.FlushDB()

	R.WriteToRedisString("test", "1")

	val, err := R.ReadFromRedisString("test")

	if err != nil {
		t.Fail()
	}

	if val != 1 {
		t.Fail()
	}

}

func TestReadFromRedisStringWithTTL(t *testing.T) {

	R.FlushDB()

	R.WriteToRedisStringWithTTL("test", "1", time.Second*10)

	val, ttl, err := R.ReadFromRedisStringWithTTL("test")

	if err != nil {
		t.Fail()
	}

	if val != 1 {
		t.Fail()
	}

	if ttl > time.Second*10 {
		t.Fail()
	}

}
