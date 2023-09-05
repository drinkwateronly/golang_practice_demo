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
	subsc := rdb.Subscribe(ctx, "channel_1")
	//for ch := range subsc.Channel() {
	//	fmt.Println(ch.Channel)
	//	fmt.Println(ch.Payload)
	//}
	for {
		message, err := subsc.ReceiveMessage(ctx)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(message.Channel)
		fmt.Println(message.Payload)
	}
}
