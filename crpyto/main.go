package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func md5Demo() {
	// 加密方式1：使用Sum方法，加密为字符串
	hash1 := md5.New()
	_, err := io.WriteString(hash1, "md5 demo")
	if err != nil {
		fmt.Println("io error")
		return
	}
	fmt.Printf("%x\n", hash1.Sum(nil))
	fmt.Printf("%X\n", hash1.Sum(nil))
	fmt.Printf("%x\n", hash1.Sum([]byte{'1', '4'}))

	// 加密方式2：使用hex.EncodeToString加密为字符串
	hash1.Reset() // 重置hash1
	hash1.Write([]byte("md5 demo"))
	fmt.Println("hash1.Size(): ", hash1.Size())
	fmt.Printf("%s", hex.EncodeToString(hash1.Sum(nil)))
}

func sha256Demo() {
	hash := sha256.New()
	hash.Write([]byte("sha256 demo"))

	fmt.Printf("%x\n", hash.Sum(nil))
	fmt.Printf("%X\n", hash.Sum(nil))
	fmt.Printf("%x\n", hash.Sum([]byte{'1', '4'}))
}

func main() {
	//md5Demo()
	sha256Demo()
	type UnionFind []int

	a := &UnionFind{}
	*a = append(*a, 0)
	fmt.Println((*a)[0])
}
