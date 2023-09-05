package main

import (
	"fmt"
	"time"
)

func TimeDemo() {
	now := time.Now()
	fmt.Println(now)

	// 都返回int
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func TimeStampDemo() {
	now := time.Now() //获取当前时间
	// 均返回int64
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)

	// Unix将时间戳转化为时间 得到的和time.now()一致的时间对象。
	timeObj := time.Unix(timestamp1, 0) // 0表示纳秒部分为0
	fmt.Println(timeObj)
}

func tickerDemo() {
	ticker := time.Tick(time.Second) // 返回通道
	for i := range ticker {          // 每秒往通道里放当前Time对象
		fmt.Println(i)
	}
}

func formatDemo() {
	now := time.Now().Add(time.Hour * 3)
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 3:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("06/01/02"))
}

func formatDemo2() {
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2023/09/12 15:04:05", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
}

func main() {

	formatDemo2()
}
