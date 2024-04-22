package main

import "fmt"

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

func DeferFunc4() (t int) {

	defer func(i int) {
		fmt.Println(i)
		fmt.Println(t)
	}(t)
	t = 1
	return 2
}

func panicFunc() {
	panic("panic")
}

func main() {
	//fmt.Println(DeferFunc1(1))
	//fmt.Println(DeferFunc2(1))
	//fmt.Println(DeferFunc3(1))
	//DeferFunc4()

	//defer func() {
	//	panic("defer panic")
	//}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		} else {
			fmt.Println("fatal")
		}
	}()

	panicFunc()

	fmt.Println("123")
}
