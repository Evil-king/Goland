package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
}

//给person类型绑定一个方法
func (p Person) test() {
	p.Name = "Mary"
	fmt.Println("test()", p.Name)
}

//给Person结构体添加一个speak 方法
func (p Person) speak() {
	fmt.Println(p.Name, "是一个goodMan")
}

//给Person结构体添加计算的方法，，可以计算从1+...+1000的结果
func (p Person) jisuan(num int) int {
	var result int
	for num = 1; num <= 1000; num++ {
		result += num
	}
	return result
}

//给person结构体jisuan2方法 该方法可以接受一个数n 计算从1+....+n的结果
func (p Person) jisuan2(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}

//给person结构体添加getSum方法，可以计算两个数的和，并返回结果
func (p Person) getSum(a, b int) int {
	return a + b
}

//匿名函数赋值给变量
var currentTime = func() {
	fmt.Println(time.Now())
}

//匿名函数当作回掉函数使用
func proc(input string, processor func(str string)) {
	processor(input)
}

//闭包作为函数返回值
func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {
	//p := Person{"Tom"}
	//p.test()
	//fmt.Println("main() p.Name = ", p.Name)
	//
	//p.speak()
	//
	//result := p.jisuan(1)
	//fmt.Println("result = ", result)
	//
	//result2 := p.jisuan2(1001)
	//fmt.Println("result2 = ", result2)
	//
	//result3 := p.getSum(50,80)
	//fmt.Println("result3 = ", result3)

	//---------------------------------------

	//proc("王小明", func(str string) {
	//	for _, v := range str {
	//		fmt.Printf("%c\n", v)
	//	}
	//})
	//
	////匿名函数
	//func(name string) {
	//	fmt.Println("My Name", name)
	//}("王小二")
	//
	//currentTime()

	//---------------------------------------

	//in := Increase()
	//fmt.Println(in())

	n := 0
	f := func() int {
		n += 1
		return n
	}
	fmt.Println(f())
}
