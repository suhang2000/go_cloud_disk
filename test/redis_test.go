package test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go_cloud_disk/core/define"
	"testing"
	"time"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: define.RedisPassword, // password set
	//DB:       0,        // use default DB
})

func TestSetValue(t *testing.T) {
	err := rdb.Set(ctx, "key", "value", time.Second*30).Err()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetValue(t *testing.T) {
	val, err := rdb.Get(ctx, "key").Result()
	if err == redis.Nil {
		t.Log("key does not exist")
	} else if err != nil {
		t.Fatal(err)
	}
	fmt.Println(val)
}
