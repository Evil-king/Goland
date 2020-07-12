package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	testStruct()
}

//定义一个结构体
type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {

	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}

	//将结构体序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("monster序列化后=%v", string(data))
}
