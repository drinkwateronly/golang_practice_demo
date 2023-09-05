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
	result, err := rdb.ZAdd(ctx, "sset1",
		&redis.Z{Score: 1.5, Member: "a"},
		&redis.Z{Score: 2, Member: "b"},
		&redis.Z{Score: 2.5, Member: "c"},
		&redis.Z{Score: 3, Member: "d"}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	count, err := rdb.ZCard(ctx, "sset1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

	zcount, err := rdb.ZCount(ctx, "sset1", "1.5", "2.5").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(zcount)

	newVal, err := rdb.ZIncrBy(ctx, "sset1", 2, "a").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(newVal)

	zRange, err := rdb.ZRange(ctx, "sset1", 0, 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(zRange)

	op := redis.ZRangeBy{
		Min:    "2.5", // 最小分数
		Max:    "3.5", // 最大分数
		Offset: 0,     // 偏移量
		Count:  5,     // 从偏移量取多少数据
	}
	zList, err := rdb.ZRangeByScoreWithScores(ctx, "sset1", &op).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(zList)

	zRange, err = rdb.ZRange(ctx, "sset1", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(zRange)

	//result, err = rdb.ZRemRangeByRank(ctx, "sset1", -2, -2).Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(result)

	result, err = rdb.ZRemRangeByScore(ctx, "sset1", "2.5", "(3.5").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	zRange, err = rdb.ZRange(ctx, "sset1", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(zRange)

}
