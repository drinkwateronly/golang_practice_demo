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
	ctx := context.Background()
	// 创建事务
	pipeline := rdb.TxPipeline()
	// 执行事务
	pipeline.Set(ctx, "key", "1", 0)
	pipeline.Expire(ctx, "key", time.Second*60)
	pipeIncr := pipeline.Incr(ctx, "key")
	pipeline.Get(ctx, "key")
	// 提交事务
	_, err := pipeline.Exec(ctx)
	if err != nil {
		panic(err)
	}
	// 查看部分运行结果
	newVal, err := pipeIncr.Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(newVal)

}
