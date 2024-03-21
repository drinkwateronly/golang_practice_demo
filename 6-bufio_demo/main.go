package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func bufioReaderDemo() {
	str := "0123456789012345678901234"           // len(str) = 44
	strReader := strings.NewReader(str)          // io.Reader
	reader := bufio.NewReaderSize(strReader, 20) // bufio.Reader

	res := make([]byte, 5)
	for i := 0; i < 10; i++ {
		peekRes, err := reader.Peek(5)
		fmt.Printf("%d) peek content：%s; peek error: %v\n", i, peekRes, err)

		n, err := reader.Read(res)
		if err != nil {
			fmt.Printf("the %d-th read fail: %s\n", i+1, err)
			return
		}
		fmt.Printf("%d) read content：%s; res: %s ;buffer remains：%2d B\n",
			i+1, string(res[:n]), res, reader.Buffered())
	}
}

func readSliceDemo() {
	str := "abcdefghijklmnopqrstuvwxyz1234567890" // len(str) = 44
	strReader := strings.NewReader(str)           // io.Reader
	reader := bufio.NewReaderSize(strReader, 16)  // bufio.Reader

	b, err := reader.ReadSlice('q') // 第17个字母
	fmt.Println("b:", string(b), " error:", err)

	b, err = reader.ReadSlice('5') //
	fmt.Println("b:", string(b), " error:", err)

	b, err = reader.ReadSlice('a')
	fmt.Println("b:", string(b), " error:", err)

	reader.ReadLine()
}

func readBytesDemo() {
	str := "abcdefghijklmnopqrstuvwxyz1234567890" // len(str) = 44
	strReader := strings.NewReader(str)           // io.Reader
	reader := bufio.NewReaderSize(strReader, 16)  // bufio.Reader

	b, err := reader.ReadBytes('q') // 第17个字母
	fmt.Println("b:", string(b), " error:", err)

	b, err = reader.ReadBytes('5') //
	fmt.Println("b:", string(b), " error:", err)

	b, err = reader.ReadBytes('a')
	fmt.Println("b:", string(b), " error:", err)
}

//func demo() {
//	str := "abcdefghijklmnopqrstuvwxyz1234567890" // len(str) = 44
//	strReader := strings.NewReader(str)           // io.Reader
//	reader := bufio.NewReaderSize(strReader, 16)  // bufio.Reader
//
//	reader.Discard()
//}

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

func ScanDemo1() {
	scanner := bufio.NewScanner(os.Stdin) // 带缓冲区的scanner
	writer := bufio.NewWriter(os.Stdout)  // 带缓冲区的writer
	defer writer.Flush()                  // 必要操作
	for {
		boo := scanner.Scan() // 从换行符分隔的文本读取一行，换行符已经被去掉
		// 出现错误或读到了文件末尾
		if boo == false {
			break
		}
		/* Note: 没有必要的错误处理
		// 当读到io.EOF时，err为nil
		if err := scanner.Err(); err != nil {
			break
		}*/
		// 从scanner中以string获取当前读到的一行
		input := scanner.Text()
		// 数据处理
		list := strings.Fields(input)
		// 输出
		fmt.Fprintf(writer, "%q\n", list)
		writer.Flush() // 必要操作
	}
}

func ScanDemo2() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	// 由于n一定是整数，不做错误处理
	fmt.Fscanf(os.Stdin, "%d\n", &n)
	for i := 0; i < n; i++ {
		boo := scanner.Scan()
		if boo == false {
			break
		}
		input := scanner.Text()
		list := strings.Fields(input)
		fmt.Fprintf(writer, "%q\n", list)
		writer.Flush()
	}
}

func main() {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 0.2223), 64)
	fmt.Println(value)

	value = math.Trunc(0.125321*1e2+0.5) * 1e-2
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	fmt.Println(value)
}
