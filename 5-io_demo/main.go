package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func readFileDemo() {
	filename := "5-io_demo/file-for-io.txt"
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

func readAtDemo() {
	byteReader := bytes.NewReader([]byte("0123456789"))
	store := make([]byte, 5)
	n, err := byteReader.ReadAt(store, 4)
	if err != nil {
		fmt.Println(err, n, string(store))
	}
	fmt.Println(err, n, string(store))
	// <nil> 5 jieda
}

func writeAtDemo1() {
	f, _ := os.OpenFile("./5-io_demo/file-for-io.txt", os.O_RDWR, 0775)
	defer f.Close()
	//f.WriteString("0123456789")
	n, err := f.WriteAt([]byte("从第3个位置写入"), 100)
	if err != nil {
		fmt.Println("WriteAt error: ", err)
		return
	}
	fmt.Println(n)
}

func writeAtDemo2() {
	f, _ := os.OpenFile("./5-io_demo/file-for-io.txt", os.O_RDWR, 0775)
	defer f.Close()
	n, err := f.WriteAt([]byte("从第3个位置写入"), 30)
	if err != nil {
		fmt.Println("WriteAt error: ", err)
		return
	}
	fmt.Println(n)
}

func writeToDemo() {
	stdReader := bufio.NewReader(os.Stdin)
	stdWriter := bufio.NewWriter(os.Stdout)
	// 将从os.Stdin读到数据直接写到os.Stdout中
	n, err := stdReader.WriteTo(stdWriter)
	fmt.Println(n)
	if err != nil {
		return
	}
}

func readFromDemo() {
	stdReader := bufio.NewReader(os.Stdin)
	stdWriter := bufio.NewWriter(os.Stdout)
	// 将从os.Stdin读到数据直接写到os.Stdout中
	n, err := stdWriter.ReadFrom(stdReader)
	fmt.Println(n)
	if err != nil {
		return
	}
}

func readFromFileDemo() {
	f, _ := os.Open("./5-io_demo/file-for-io.txt")
	defer f.Close()
	stdWriter := bufio.NewWriter(os.Stdout)
	// 将从os.Stdin读到数据直接写到os.Stdout中
	n, err := stdWriter.ReadFrom(f)
	fmt.Println("\n", n)
	if err != nil {
		return
	}
}

func copyFromFileDemo() {
	f, _ := os.Open("./5-io_demo/file-for-io.txt")
	n, err := io.Copy(os.Stdout, f)
	if err != nil {
		fmt.Println("Copy fail: ", err)
		return
	}
	fmt.Println("\nn: ", n)
	f.Close()
}

func readDemo() {
	strReader := strings.NewReader("01234567890")
	b1 := make([]byte, 10)
	n, err := io.ReadAtLeast(strReader, b1, 5)
	if err != nil {
		fmt.Println("ReadAtLeast fail: ", err)
		return
	}
	fmt.Println("ReadAtLeast: ", string(b1), n)

	strReader = strings.NewReader("01234567890")
	b2 := make([]byte, 6)
	n, err = io.ReadFull(strReader, b2)
	if err != nil {
		fmt.Println("ReadFull fail: ", err)
		return
	}
	fmt.Println("ReadFull: ", string(b2), n)
}

func main() {
	copyFromFileDemo()
}
