package main

import (
	"encoding/json"
	"fmt"
)

type A int

const (
	jsonBlob     = `{"name": "Aaron"}`
	fulljsonBlob = `{"name":"Aaron", "age":0}`
)

// Example结构体包含age和name字段
type Example struct {
	Age  int    `json:"age,omitempty"`
	Name string `json:"name"`
}

func main() {

	//开平方示例
	//i := "25"
	//
	//// i 是整型，所以需要转型
	//
	//// 25开方结果是 5
	//
	//a, err := strconv.ParseFloat(i, 64)
	//if err != nil {
	//}
	//
	//fmt.Print(a)

	test()
}

func test() error {

	e := Example{}

	// 注意jsonBlob没有age字段
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}

	//由于age被设置为omitempty(为空则不输出) 所以显示 Regular Marshal, with no age: {"name":"Aaron"}
	fmt.Println("Regular Marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, with age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	//Regular Marshal, with age = 0: {"name":"Aaron"}
	fmt.Println("Regular Marshal, with age = 0:", string(value))

	return nil
}
