package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func main() {
	ctx := context.Background()
	rdb.Publish(ctx, "channel_1", "msg")
	chs, _ := rdb.PubSubNumSub(ctx, "channel_1").Result()
	for ch, count := range chs {
		fmt.Println(ch)    // channel名字
		fmt.Println(count) // channel的订阅者数量
	}
}
