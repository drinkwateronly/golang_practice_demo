package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// 方式1：依赖struct
type MyClaims struct {
	Username             string
	jwt.RegisteredClaims // 不要写成RegisteredClaims jwt.RegisteredClaims
}

// 方式2：依赖map，使用方式更复杂一些
type MyClaims2 struct {
	Username string
	jwt.MapClaims
}

func main() {
	jwtMapClaims()
}

func jwtMapClaims() {
	// 加密
	mc := MyClaims2{
		Username: "chen jie",
		MapClaims: jwt.MapClaims{
			"iss": "chen jie",
			"nbf": time.Now().Unix() - 5,
			"exp": time.Now().Unix() + 10,
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)
	fmt.Println(t)

	// 一致
	strToSigned := []byte("123")
	signedString, err := t.SignedString(strToSigned)
	if err != nil {
		panic(err)
	}
	fmt.Println(signedString)

	//time.Sleep(time.Second * 7)

	// 解密
	strToValid := []byte("123")
	//strToValid := []byte("321")
	fmt.Println("解密")
	// jwt.ParseWithClaims 输入 需要解析的JWT字符串、一个实现了jwt.Claims接口的结构体、用于提供验证签名所需的密钥的回调函数
	t, err = jwt.ParseWithClaims(signedString, &MyClaims2{}, func(token *jwt.Token) (interface{}, error) {
		return strToValid, nil
	})
	// 如果解析失败
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Claims.GetIssuer())
	fmt.Println(t.Claims.GetExpirationTime())

}

func jwtRegisteredClaims() {
	// 加密
	mc := MyClaims{
		Username: "chen jie",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "chen jie",                                              // 签发人
			NotBefore: &jwt.NumericDate{Time: time.Now().Add(-time.Minute)},    // 一分钟前生效
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Second * 5)}, // 一小时内过期
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)
	fmt.Println(t)

	strToSigned := []byte("123")
	signedString, err := t.SignedString(strToSigned)
	if err != nil {
		panic(err)
	}
	fmt.Println(signedString)

	time.Sleep(time.Second * 7)

	// 解密
	strToValid := []byte("123")
	//strToValid := []byte("321")
	fmt.Println("解密")
	t, err = jwt.ParseWithClaims(signedString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return strToValid, nil
	})
	// 如果解析失败
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Claims.GetIssuer())
	fmt.Println(t.Claims.GetExpirationTime())
}
