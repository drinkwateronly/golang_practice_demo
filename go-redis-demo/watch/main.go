package main

import (
	"context"
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

	rdb.Set(ctx, "key", 3, 0)

	// 定义一个回调函数，用于处理事务逻辑
	fn := func(tx *redis.Tx) error {
		// 先查询下当前watch监听的key的值
		v, err := tx.Get(ctx, "key").Int()
		if err != nil && err != redis.Nil {
			return err
		}
		// 这里可以处理业务
		v++

		// 如果key的值没有改变的话，Pipelined函数才会调用成功
		_, err = tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			// 在这里给key设置最新值
			pipe.Set(ctx, "key", v, 0)
			return nil
		})
		return err
	}

	// 使用Watch监听一些Key, 同时绑定一个回调函数fn, 监听Key后的逻辑写在fn这个回调函数里面
	// 如果想监听多个key，可以这么写：client.Watch(ctx,fn, "key1", "key2", "key3")
	retryCount := 3 // 重试次数
	for i := 0; i < retryCount; i++ {
		err := rdb.Watch(ctx, fn, "key")
		if err == nil { // 成功
			break
		}
		if err == redis.TxFailedErr {
			continue
		}
	}

}
