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
	rows, err := rdb.HSet(ctx, "user1", "user_name", "cj").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)

	rows, err = rdb.HSet(ctx, "user1", "age", 18).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)

	// GET，也有MGET
	res, err := rdb.HGet(ctx, "user1", "age").Result()
	if err != nil {
		return
	}
	fmt.Println(res)

	newVal, err := rdb.HIncrBy(ctx, "user1", "age", 2).Result()
	if err != nil {
		return
	}
	fmt.Println(newVal)

	getAllResult, err := rdb.HGetAll(ctx, "user1").Result()
	if err != nil {
		return
	}
	for k, v := range getAllResult {
		fmt.Println(k, v)
	}

	keysResult, err := rdb.HKeys(ctx, "user1").Result()
	if err != nil {
		return
	}
	for _, s := range keysResult {
		fmt.Println(s)
	}

	hlen, err := rdb.HLen(ctx, "user1").Result()
	if err != nil {
		return
	}
	fmt.Println(hlen)

	rows, err = rdb.HDel(ctx, "user1", "age", "user_name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)
}
