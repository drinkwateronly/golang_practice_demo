package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(goodsChan chan<- int, goodsNum int) {
	for i := 0; i < goodsNum; i++ {
		goods := rand.Intn(100)
		goodsChan <- goods
		fmt.Println("生产产品：", goods)
		time.Sleep(time.Second)
	}
}

func deliver(goodsChan <-chan int, shopChan chan<- int) {
	for {
		goods := <-goodsChan
		fmt.Println("运输产品：", goods)
		shopChan <- goods
	}
}

func consumer(shopChan <-chan int, exitChan chan<- bool) {
	for i := 0; i < 10; i++ {
		goods := <-shopChan
		fmt.Println("消费次数序号：", i, " 商品：", goods)
		time.Sleep(time.Second * 2)
	}
	exitChan <- true
}

/*
生产者：每一秒产生一个商品
运输者：一有商品就运输到商店
消费者：每两秒消费一个商品，消费十次后，程序结束
*/

func main() {
	goodsChan := make(chan int, 5)
	shopChan := make(chan int, 5)
	exitChan := make(chan bool)
	// 如果goodsNum<10，那么消费者在消费9次后，会进入死锁
	// 如果goodsNum>10，那么消费者消费完后，商店还会有剩余商品
	goodsNum := 10
	go producer(goodsChan, goodsNum)
	go deliver(goodsChan, shopChan)
	go consumer(shopChan, exitChan)
	select {
	case <-exitChan:
		return
	}
}
