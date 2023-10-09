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
	rows, err := rdb.LPush(ctx, "list1", 1, 2, 3, "str1", 1, "str2", 2).Result()
	if err != nil {
		return
	}
	fmt.Println(rows)

	val, err := rdb.RPop(ctx, "list1").Result() // rows为插入的数量
	fmt.Println(val)

	list, err := rdb.LRange(ctx, "list1", 0, 2).Result() // rows为插入的数量
	fmt.Println(list)

	inserts, err := rdb.LInsert(ctx, "list1", "after", 2, "insert").Result()
	fmt.Println(inserts)

}
