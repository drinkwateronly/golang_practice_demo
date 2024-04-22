package main

import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 捕捉不到panic
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 只捕捉到了后panic
		}
	}()

	defer func() {
		panic("defer panic") // 后panic
	}()

	panic("panic") // 先panic
}
