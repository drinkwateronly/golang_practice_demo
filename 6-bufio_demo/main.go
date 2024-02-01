package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func bufioReaderDemo() {
	str := "01234567890123456789012345678901234567890123" // len(str) = 44
	strReader := strings.NewReader(str)                   // io.Reader
	reader := bufio.NewReaderSize(strReader, 25)          // bufio.Reader

	res := make([]byte, 5)
	for i := 0; i < 10; i++ {
		n, err := reader.Read(res)
		if err != nil {
			fmt.Printf("the %d-th read fail: %s", i+1, err)
			return
		}
		fmt.Printf("%d) read content：%s; res: %s ;buffer remains：%2d B\n",
			i+1, string(res[:n]), res, reader.Buffered())
	}
}

func bufioWriterDemo() {
	// 打开可写文件
	file, _ := os.OpenFile("./6-bufio_demo/file.txt", os.O_WRONLY, 0777)
	defer file.Close()
	writer := bufio.NewWriterSize(file, 5) // io.Writer
	// 第一次写
	_, err := writer.Write([]byte("123456")) // bufio.Writer
	if err != nil {
		fmt.Println("bufio writer failed:", err)
		return
	}
	fmt.Println("Buffered:", writer.Buffered(), "; Available:", writer.Available())

	// 第二次写
	_, err = writer.Write([]byte("12345")) // bufio.Writer
	if err != nil {
		fmt.Println("bufio writer failed:", err)
		return
	}
	fmt.Println("Buffered:", writer.Buffered(), "; Available:", writer.Available())

	// 第三次写
	_, err = writer.Write([]byte("6789")) // bufio.Writer
	if err != nil {
		fmt.Println("bufio writer failed:", err)
		return
	}
	fmt.Println("Buffered:", writer.Buffered(), "; Available:", writer.Available())

	// 由于缓冲未满，如果不flush，则不会写入文件中
	//err = writer.Flush()
	//if err != nil {
	//	fmt.Println("flush failed:", err)
	//	return
	//}
	//fmt.Println("Buffered:", writer.Buffered(), "; Available:", writer.Available())
}

func main() {
	bufioWriterDemo()
}
