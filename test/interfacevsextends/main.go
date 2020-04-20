package main

import "fmt"

//结构体
type Monkey struct {
	Name string
}
//结构体
type LittleMonkey struct {
	Monkey //继承
}

//声明一个接口
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swiming()
}

func (monkey *Monkey) climbing() {
	fmt.Println(monkey.Name, "生来就会爬树")
}

func (LittleMonkey *LittleMonkey) Flying() {
	fmt.Println(LittleMonkey.Name, "通过学习会飞翔")
}

func (LittleMonkey *LittleMonkey) Swiming() {
	fmt.Println(LittleMonkey.Name, "通过学习会游泳")
}

func main() {
	monkey := LittleMonkey{
		Monkey{Name: "悟空"},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swiming()
}
