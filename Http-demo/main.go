package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// 创建http.Client对象
	client := &http.Client{}

	// 创建HTTP请求
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		panic(err)
	}

	// 发送HTTP请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// 处理HTTP响应
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	params := []string{"xjx", "zzz"}
	req, err := http.NewRequest("POST", "https://www.example.com/api/v1/posts", nil)
	if err != nil {
		log.Fatal(err)
	}
	data := url.Values{}

	for _, tag := range params {
		data.Add("tags", tag)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Body = io.NopCloser(strings.NewReader(data.Encode()))

	http.Server{}
}
