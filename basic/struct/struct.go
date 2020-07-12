package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              //指针
	slice  []int             //切片
	map1   map[string]string //切片
}

type Person1 struct {
	Name string
	Age  int
}

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	var p1 Person

	//使用slice 一定先要make
	p1.slice = make([]int, 10)

	//使用map也是一定要先make
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "Tom"
	fmt.Println(p1)

	p2 := Person1{"mary", 20}
	fmt.Println("p2=", p2)

	var p3 *Person1 = new(Person1)
	(*p3).Age = 10
	(*p3).Name = "Jim"
	fmt.Println("p3=", *p3)

	monster := Monster{"牛魔王", 500, "芭蕉扇"}
	monsterStr, _ := json.Marshal(monster)
	fmt.Println("monsterStr=", string(monsterStr))

}
