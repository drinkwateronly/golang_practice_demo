package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func WriteToDemo() {
	stdReader := bufio.NewReader(os.Stdin)
	stdWriter := bufio.NewWriter(os.Stdout)
	stdReader.WriteTo(stdWriter)
}

func ReadAtDemo() {
	byteReader := bytes.NewReader([]byte("chenjiedashuaige"))
	store := make([]byte, 5)
	n, err := byteReader.ReadAt(store, 4)
	if err != nil {
		fmt.Println(err, n, string(store))
	}
	fmt.Println(err, n, string(store))
}

func WriteAtDemo() {
	f, _ := os.OpenFile("./io-demo/file-for-io", os.O_RDWR, 0775)
	defer f.Close()
	f.WriteString("01234567890")
	f.WriteAt([]byte("从第3个位置写入"), 3)
}

func main() {

	WriteAtDemo()
}
