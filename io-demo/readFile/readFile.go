package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 打开一个文本文件
	file, err := os.Open("./io-demo/readFile/example.txt")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	// 创建一个 Scanner 对象，将文件作为输入源
	scanner := bufio.NewScanner(file)

	// 逐行读取文本文件并输出
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("行内容:", line)
	}

	// 检查扫描是否发生错误
	if err := scanner.Err(); err != nil {
		fmt.Println("扫描错误:", err)
	}

}
