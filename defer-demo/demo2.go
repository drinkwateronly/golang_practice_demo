package main

import "fmt"

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
	}
}

func performTask() int {
	defer handlePanic()
	fmt.Println("task")
	panic("panic")
	// 后续不会执行
	fmt.Println("finish")
	return 1
}

func main() {

	num := performTask()

	// 从panic的函数恢复，继续执行
	fmt.Println(num)
	fmt.Println("Main function continues.")
}
