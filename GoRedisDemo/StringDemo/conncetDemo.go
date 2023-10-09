package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
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
	// context
	ctx := context.Background()

	fmt.Println("####### Set & Get ######")
	val, err := rdb.Set(ctx, "strKey", "this is a string", time.Second*10).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Set val: ", val)
	val, err = rdb.Get(ctx, "strKey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("GET val: ", val)

	fmt.Println("####### Do ######")
	result, err := rdb.Do(ctx, "GET", "strKey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result.(string))

	fmt.Println("####### GetSet ######")
	oldVal, err := rdb.GetSet(ctx, "strKey", "this is another str").Result()
	if err != nil {
		return
	}
	fmt.Println("GetSet oldVal: ", oldVal)
	newVal, err := rdb.Get(ctx, "strKey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("GetSet newVal: ", newVal)

	fmt.Println("####### SetNX ######")
	ok, err := rdb.SetNX(ctx, "strKey1", "value2", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("SetNX success? if false means key is already existed:", ok)

	fmt.Println("####### MSet ######")
	err = rdb.MSet(ctx, "iKey", 1, "fKey", 1.1, "strKey2", "str2").Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("####### MGet ######")
	valList, err := rdb.MGet(ctx, "strKey", "strKey2").Result()
	if err != nil {
		panic(err)
	}
	for _, val := range valList {
		fmt.Println(val.(string))
	}

	fmt.Println("####### Incr IncrBy IncrByFloat######")
	newVal1, err := rdb.Incr(ctx, "iKey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(newVal1)
	newVal2, err := rdb.IncrByFloat(ctx, "fKey", 1.1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(newVal2)
	newVal3, err := rdb.IncrBy(ctx, "iKey", 3).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(newVal3)

	fmt.Println("####### Incr IncrBy IncrByFloat######")
	newVal1, err = rdb.Decr(ctx, "iKey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(newVal1)
	newVal3, err = rdb.DecrBy(ctx, "iKey", 10).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(newVal3)

	fmt.Println("####### Del ######")
	line, err := rdb.Del(ctx, "iKey", "fKey", "strKey1", "strKey2").Result()
	fmt.Println("affect rows: ", line)
	if err != nil {
		panic(err)
	}

	fmt.Println("####### Expire ######")
	ok, err = rdb.Expire(ctx, "strKey", time.Second*10).Result()
	fmt.Println("set expire success?: ", ok)
	if err != nil {
		panic(err)
	}
}
