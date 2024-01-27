package main

import (
	"fmt"
	"time"
)

func timeDemo() {
	now := time.Now()
	fmt.Println(now)

	// 都返回int
	year := now.Year()                 // 年
	month := now.Month()               // 月
	day := now.Day()                   // 日
	hour := now.Hour()                 // 时
	minute := now.Minute()             // 分
	second := now.Second()             // 秒
	hour, minute, second = now.Clock() // 同时获得时分秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	yearDay := now.YearDay() // 当年的第几天
	fmt.Println(yearDay)

	year, week := now.ISOWeek() // 第几年的第几周
	fmt.Println(year, week)
}

func timeStampDemo() {
	now := time.Now() //获取当前时间
	// -------Time对象转化为时间戳，即int64
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1: %v\n", timestamp1)
	fmt.Printf("current timestamp2: %v\n", timestamp2)

	// ------Unix()将时间戳转化为Time对象
	timeObj := time.Unix(timestamp1, 0) // 0表示纳秒部分为0
	fmt.Println(timeObj)

	// 但纳秒时间戳不能被Unix()转化为Time对象，因为也是int64可以调用，但结果有错
	timeObj2 := time.Unix(timestamp2, 0) // 0表示纳秒部分为0
	fmt.Println(timeObj2)                // 有错
}

func timeOperation() {
	// 创建时间
	t1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)

	// 时间加减
	t2 := t1.Add(time.Hour)
	fmt.Println(t2.Format("2006-01-02 15:04:05"))
	subDuration := t2.Sub(t1)
	fmt.Println(subDuration)

	// 时间比较
	fmt.Println("equal?: ", t1.Equal(t2))
	fmt.Println("before?: ", t1.Before(t2))
	fmt.Println("after?: ", t1.After(t2))
}

func formatDemo() {
	now := time.Now()
	// Time->字符串
	fmt.Println(now.Format("2006-01-02 15:04:05.000000 Mon PM Jan"))
	fmt.Println(now.Format("2006-1-2 3:04:05.000000"))
	fmt.Println(now.Format("03:04 06/01/02"))

	// 字符串->Time
	// 使用默认Parse
	timeObj, err := time.Parse("2006/01/02 15:04:05", "2023/09/12 15:04:05")
	if err != nil {
		panic(err)
	}
	fmt.Println(timeObj, timeObj.Location())

	// 加载时区
	// loc := time.Local
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err = time.ParseInLocation("2006/01/02 15:04:05", "2023/09/12 15:04:05", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj, timeObj.Location())
}

func tickerDemo() {
	go func() {
		ticker := time.NewTicker(time.Second)
		for i := range ticker.C { // 每秒往通道里放当前Time对象
			fmt.Println("NewTicker: ", i)

		}
	}()

	go func() {
		tickerChan := time.Tick(time.Second) // 返回通道
		for i := range tickerChan {          // 每秒往通道里放当前Time对象
			fmt.Println("Tick: ", i)
		}
	}()
	time.Sleep(time.Second * 3)
}

func delayDemo() {
	fmt.Println(time.Now())
	go func() {
		timer := time.NewTimer(time.Second * 3)
		for i := range timer.C {
			fmt.Println("after 3 second: ", i)
		}
	}()

	go func() {
		C := time.After(time.Second * 2)
		for i := range C {
			fmt.Println("after 2 second: ", i)
		}
	}()

	time.Sleep(time.Second * 4)
}

func main() {
	//timeDemo()
	//timeOperation()
	//timeStampDemo()
	//formatDemo()
	delayDemo()
}
