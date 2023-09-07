package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		读取数据：
			Reader接口
				Read(p []byte) (n int, err error)
	*/
	filename := "io-demo/file-for-io"
	file, err := os.Open(filename) // 只读
	if err != nil {
		return
	}
	defer file.Close()
	// 读取数据
	bs := make([]byte, 3, 3)
	n, err := file.Read(bs)
	fmt.Println(err, n, string(bs)) // <nil> 3 123
	n, err = file.Read(bs)
	fmt.Println(err, n, string(bs)) //<nil> 2 453
	n, err = file.Read(bs)
	fmt.Println(err, n, string(bs)) // EOF 0 453
}
