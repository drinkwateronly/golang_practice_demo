package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

//https://www.liwenzhou.com/posts/Go/json-tricks/#c-0-0-0

func demo1() {
	type Partner struct {
		Name string `json:"firstName"`
		Age  uint   `json:"age"`
	}

	type Person struct {
		FirstName string  `json:"firstName"`       // 指定字段名
		LastName  string  `json:"-"`               // 序列化/反序列化时，忽略该字段
		Email     string  `json:"email,omitempty"` // "omitempty" 忽略空值字段，需要和指定字段名搭配使用
		Age       uint    `json:"age"`             // 反序列化时，若为负值，解析失败。
		Height    float64 `json:"height,string"`   // 序列化/反序列化时，该数字以字符串形式解析。若不是字符串形式，解析失败。
		Sex       bool    `json:"sex"`             // 反序列化时，只能是true或false
		IsNative  bool    `json:"isNative,string"` // 反序列化时，只能是"true"或"false"

		PartnerA Partner  `json:"partnerA"`           // 嵌套结构体
		PartnerB *Partner `json:"partnerB,omitempty"` // 只有嵌套结构体是指针时，才能被忽略空值
	}

	// 序列化
	p := Person{
		FirstName: "John",
		LastName:  "Smith",
		Email:     "john@example.com",
		Age:       11,
		Height:    180.01,
		Sex:       true,
		IsNative:  false,

		PartnerA: Partner{
			Name: "",
			Age:  0,
		},
		PartnerB: nil,
	}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Marshal fail: ", err)
	} else {
		fmt.Printf("Marshal success: %v\n", string(b))
	}

	// 反序列化
	s := `{
		"firstName":"John",
		"lastName":"Smith",
		"email":"john@example.com",
		"age":11,
		"height":"180.01",
		"sex":true,
		"isNative":"false",
		"partnerA":{
			"firstName":"",
			"age":0
		}
	}`
	var p2 Person
	err = json.Unmarshal([]byte(s), &p2)
	if err != nil {
		fmt.Println("Unmarshal fail: ", err)
	} else {
		fmt.Printf("Unmarshal success: %+v\n", p2)
	}
}

func Demo2() {
	type User struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	type UserPublic struct {
		*User
		Password string `json:"-"`
	}
}

func demo3() {
	//处理数字和字符类型的数字
	// 在该json中，num是数字，numString是字符串数字
	s := `{"num": 1, "numString": "12.34"}`

	type Nums struct {
		Num       int     `json:"num"`
		NumString float64 `json:"numString,string"`
	}
	var nums Nums

	err := json.Unmarshal([]byte(s), &nums)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("%+v\n", nums)

}

func parseTimeInDefaultDemo() {
	type Times struct {
		CreatedAt time.Time `json:"createdAt"`
	}
	// 反序列化
	str := `{"createdAt": "2024-01-25 01:23:45"}`
	var times Times
	err := json.Unmarshal([]byte(str), &times)
	if err != nil {
		// parsing time "2024-01-25 01:23:45" as "2006-01-02T15:04:05Z07:00": cannot parse " 01:23:45" as "T"
		fmt.Println("unmarshal error:", err)
	}

	// 序列化
	b, err := json.Marshal(Times{CreatedAt: time.Now()})
	if err != nil {
		// 不会err
		return
	}
	fmt.Println(string(b))
}

type Order struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

const layout = "2006-01-02 15:04:05"

func (o *Order) MarshalJSON() ([]byte, error) {
	type TempOrder Order // 定义与Order字段一致的新类型，避免直接嵌套Order进入死循环
	ot := struct {
		CreatedAt  string `json:"createdAt"`
		*TempOrder        // 避免直接嵌套Order进入死循环
	}{
		CreatedAt: o.CreatedAt.Format(layout),
		TempOrder: (*TempOrder)(o), // 类型转换，虽然TempOrder也有CreatedAt，但编译器会直接使用当前作用域已有的CreatedAt
	}
	fmt.Println("MarshalJSON: ", ot.CreatedAt)
	return json.Marshal(ot)
}

func (o *Order) UnmarshalJSON(data []byte) error {
	type TempOrder Order // 定义与Order字段一致的新类型，
	ot := struct {
		CreatedAt  string `json:"createdAt"`
		*TempOrder        // 避免直接嵌套Order进入死循环
	}{
		TempOrder: (*TempOrder)(o), // 类型转换，虽然TempOrder也有CreatedAt，但编译器会直接使用当前作用域已有的CreatedAt，也就是空值str
	}
	fmt.Println("ot created: ", ot.CreatedAt)

	if err := json.Unmarshal(data, &ot); err != nil {
		return err
	}

	var err error
	o.CreatedAt, err = time.Parse(layout, ot.CreatedAt)
	if err != nil {
		return err
	}
	fmt.Println("UnmarshalJSON: ", ot.CreatedAt)
	return nil
}

func parseTimeInCustomDemo() {
	// 反序列化
	str := `{"id":"123","createdAt": "2024-01-25 01:23:45"}`
	var order Order
	err := json.Unmarshal([]byte(str), &order)
	if err != nil {
		// parsing time "2024-01-25 01:23:45" as "2006-01-02T15:04:05Z07:00": cannot parse " 01:23:45" as "T"
		fmt.Println("unmarshal error: ", err)
	} else {
		fmt.Printf("unmarshal success: %+v\n", order)
	}

	// 序列化
	b, err := json.Marshal(&Order{CreatedAt: time.Now(), Id: "123"})
	if err != nil {
		fmt.Println("marshal error:", err)
	} else {
		fmt.Println("marshal success: " + string(b))
	}
}

func parseNumbersDemo() {
	s := `{"name":"John", "age":123, "height": 1234567890}`
	var m1 map[string]interface{}
	// 常用的Unmarshal反序列化
	err := json.Unmarshal([]byte(s), &m1)
	if err != nil {
		fmt.Println("Unmarshal fail: ", err)
		return
	}
	fmt.Printf("value: %#v, type:%T\n", m1["name"], m1["name"])
	fmt.Printf("value: %#v, type:%T\n", m1["age"], m1["age"])
	fmt.Printf("value: %#v, type:%T\n", m1["height"], m1["height"])

	// 使用decoder反序列化，指定使用number类型，此时会解析为json.Number类型
	var m2 map[string]interface{}
	decoder := json.NewDecoder(bytes.NewReader([]byte(s)))
	decoder.UseNumber()

	err = decoder.Decode(&m2)
	if err != nil {
		fmt.Println("Decode fail: ", err)
		return
	}
	fmt.Printf("value: %#v, type:%T\n", m2["name"], m2["name"])
	fmt.Printf("value: %#v, type:%T\n", m2["age"], m2["age"])
	fmt.Printf("value: %#v, type:%T\n", m2["height"], m2["height"])

	age, err := m2["age"].(json.Number).Int64()
	if err != nil {
		fmt.Println("parse to int64 fail: ", err)
		return
	}
	fmt.Printf("value: %#v, type:%T\n", age, age)
}

func main() {
	parseNumbersDemo()
}
