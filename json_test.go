package RedisGo_test

import (
	"testing"
	"time"
)

func TestWriteJSONToRedisWithTTL(t *testing.T) {

	R.FlushDB()
	jsonString := `{"test": "1"}`
	err := R.WriteJSONToRedisWithTTL("test", jsonString, time.Second*10)

	if err != nil {
		t.Fail()
	}

}

func TestReadJSONFromRedisWithTTL(t *testing.T) {

	R.FlushDB()

	jsonString := `{"test": "1"}`
	R.WriteJSONToRedisWithTTL("test", jsonString, time.Second*10)

	_, ttl, err := R.ReadJSONFromRedisWithTTL("test")
	if err != nil {
		t.FailNow()
	}

	if ttl > time.Second*10 {
		t.FailNow()
	}

}
