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
	result, err := rdb.SAdd(ctx, "set1", 200, "a", "300", "c", 100).Result()
	if err != nil {
		return
	}
	fmt.Println(result)

	popNResult, _ := rdb.SPopN(ctx, "set1", 3).Result()
	fmt.Println(popNResult)
}
