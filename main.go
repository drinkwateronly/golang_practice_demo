package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/robertkrimen/otto"
)

func getPwdDefaultEncryptSalt(text string) string {
	regex := regexp.MustCompile(`pwdDefaultEncryptSalt\s*=\s*"([^"]+)"`)
	match := regex.FindStringSubmatch(text)
	log.Println("match: ", match)
	if len(match) >= 2 {
		log.Println("pwdDefaultEncryptSalt:", match[1])
		return match[1]
	} else {
		log.Println("未找到pwdDefaultEncryptSalt的值")
		return "error"
	}
}

func getlt(text string) string {
	// 使用正则表达式提取lt值
	re := regexp.MustCompile(`<input\s+type="hidden"\s+name="lt"\s+value="([^"]+)"`)
	match := re.FindStringSubmatch(text)
	if len(match) >= 2 {
		ltValue := match[1]
		log.Println("lt:", ltValue)
		return ltValue

	} else {
		fmt.Println("lt value not found")
		return "error"
	}
}

func getExecution(text string) string {
	// 使用正则表达式提取lt值
	re := regexp.MustCompile(`<input\s+type="hidden"\s+name="execution"\s+value="([^"]+)"`)
	match := re.FindStringSubmatch(text)
	if len(match) >= 2 {
		ltValue := match[1]
		log.Println("execution:", ltValue)
		return ltValue

	} else {
		fmt.Println("lt value not found")
		return "error"
	}
}

func getLoginMessage() (string, string, string, string) {

	targetUrl := "https://authserver.szu.edu.cn/authserver/login"

	params := url.Values{}
	params.Add("service", "https://ehall.szu.edu.cn:443/qljfwapp/sys/lwSzuCgyy/index.do#/sportVenue")
	encodedParams := params.Encode()

	targetUrl = targetUrl + "?" + encodedParams

	response, err := http.Get(targetUrl)
	if err != nil {
		log.Println("GET request failed:", err)
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Failed to read response:", err)
	}
	log.Println(response.Header)

	//读取响应内容成功
	// log.Println(string(body))
	log.Println("状态码: ", response.StatusCode) // 获取状态码
	log.Println("状态: ", response.Status)      // 获取状态码对应的文案

	if response.StatusCode == 401 {
		log.Println("cookie 错误或失效")
	}
	if response.StatusCode != 200 {
		log.Println("状态码错误，请联系软件作者解决")
	}

	// 获取密钥
	pwdDefaultEncryptSalt := getPwdDefaultEncryptSalt(string(body))
	// 获取lt
	lt := getlt(string(body))
	execution := getExecution(string(body))
	// 获取cookie
	newLoginCookie := strings.Split(response.Header["Set-Cookie"][0], ";")[0] + "; " + strings.Split(response.Header["Set-Cookie"][1], ";")[0] + "; org.springframework.web.servlet.i18n.CookieLocaleResolver.LOCALE=zh_CN"
	log.Println("new LoginCookie: ", newLoginCookie)
	return pwdDefaultEncryptSalt, lt, newLoginCookie, execution
}

func getEncryptPWD(pwd string, pwdDefaultEncryptSalt string) string {
	// 创建一个新的JavaScript虚拟机实例
	vm := otto.New()

	// 从JavaScript文件中读取代码
	jsCode, err := os.ReadFile("rsa.js")
	if err != nil {
		log.Fatal("Error reading JavaScript file:", err)
	}

	// 在JavaScript虚拟机中执行代码
	_, err = vm.Run(jsCode)
	if err != nil {
		log.Fatal("JavaScript execution error:", err)
	}

	// 调用JavaScript函数
	result, err := vm.Call("encryptAES", nil, pwd, pwdDefaultEncryptSalt)
	log.Println("result: ", result)
	if err != nil {
		log.Fatal("JavaScript function call error:", err)
	}

	return result.String()
}

func login(userid string, encryptPWD string, lt string, loginCookie string, execution string) {
	// 发送请求
	targetUrl := "https://authserver.szu.edu.cn/authserver/login?service=https://ehall.szu.edu.cn:443/qljfwapp/sys/lwSzuCgyy/index.do#/sportVenue"
	payloadstr := fmt.Sprintf("username=%s&password=%s&lt=%s&dllt=userNamePasswordLogin&execution=%s&_eventId=submit&rmShown=1",
		userid, encryptPWD, lt, execution)
	log.Println("payloadstr: ", payloadstr)
	payload := strings.NewReader(payloadstr)
	req, _ := http.NewRequest("POST", targetUrl, payload)
	// 设置请求头
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Cache-Control", "max-age=0")

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", loginCookie)
	req.Header.Set("Host", "authserver.szu.edu.cn")
	req.Header.Set("Origin", "https://authserver.szu.edu.cn")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://ehall.szu.edu.cn/qljfwapp/sys/lwSzuCgyy/index.do#/sportVenue")
	req.Header.Set("Sec-Ch-Ua", "\"Microsoft Edge\";v=\"123\", \"Not:A-Brand\";v=\"8\", \"Chromium\";v=\"123\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Requests", "")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("redirect!! ")
			return nil
		},
	}
	response, err := client.Do(req)
	if err != nil {
		log.Println("Error: connect failed!")
	}

	defer response.Body.Close() // 这步是必要的，防止以后的内存泄漏

	log.Println("状态码: ", response.StatusCode) // 获取状态码
	log.Println("状态: ", response.Status)      // 获取状态码对应的文案

	if response.StatusCode == 401 {
		log.Println("cookie 错误或失效")
	}
	if response.StatusCode != 200 {
		log.Println("状态码错误，请联系软件作者解决")
	}

	body, _ := io.ReadAll(response.Body) // 读取响应 body, 返回为 []byte
	log.Printf("响应为：\n %s \n", string(body))
	header := response.Header
	log.Printf("响应头为：\n %s \n", header)
}

func main() {
	// getLoginMessage()
	pwdDefaultEncryptSalt, lt, loginCookie, execution := getLoginMessage()
	userid := "2110276065"
	password := "04141512"
	encryptPWD := getEncryptPWD(password, pwdDefaultEncryptSalt)
	log.Println(encryptPWD)
	time.Sleep(time.Second * 5)
	login(userid, encryptPWD, lt, loginCookie, execution)

}

// yPkUGfflAX6YQfdE
// hXJMTYRrrbksAuM9
// HGYsRUs1tGjnDSDA
