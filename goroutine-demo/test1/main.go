package main

import "fmt"

func writeData(intChan chan int) {
	for i := 1; i <= 150; i++ {
		fmt.Println("第", i, "次写入", i)

		intChan <- i
	}
	fmt.Println("intChan关闭")
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		i, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("读取到", i)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	intChan := make(chan int)
	exitChan := make(chan bool)
	go writeData(intChan)

	go readData(intChan, exitChan)

	if <-exitChan {

		return
	}

}
