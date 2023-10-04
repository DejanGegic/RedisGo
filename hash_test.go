package RedisGo_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/DejanGegic/RedisGo"
)

var R = RedisGo.ConnectToRedisClient("localhost", "6379", "", 0)

func TestWriteToRedisHash(t *testing.T) {

	R.FlushDB()

	err := R.WriteToRedisHash(strconv.Itoa(1), strconv.Itoa(1), "test")

	if err != nil {
		t.Fail()
	}

}

func TestWriteToRedisHashWithTTL(t *testing.T) {

	R.FlushDB()

	err := R.WriteToRedisHashWithTTL(strconv.Itoa(1), strconv.Itoa(1), "test", time.Second*2)

	if err != nil {
		t.Fail()
	}

}
func TestReadFromRedisHash(t *testing.T) {

	R.FlushDB()

	err := R.WriteToRedisHash(strconv.Itoa(1), strconv.Itoa(1), "test")

	if err != nil {
		t.Fail()
	}

	val, err := R.ReadFromRedisHash(strconv.Itoa(1), strconv.Itoa(1))

	if err != nil {
		t.Fail()
	}

	if val != "test" {
		t.Fail()
	}

}

func TestReadFromRedisHashWithTTL(t *testing.T) {

	R.FlushDB()

	err := R.WriteToRedisHashWithTTL(strconv.Itoa(1), strconv.Itoa(1), "test", time.Second*2)

	if err != nil {
		t.Fail()
	}

	val, _, err := R.ReadFromRedisHashWithTTL(strconv.Itoa(1), strconv.Itoa(1))

	if err != nil {
		t.Fail()
	}

	if val != "test" {
		t.Fail()
	}

}
