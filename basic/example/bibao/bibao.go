package main

import (
	"fmt"
	"strings"
)

//累加器
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		//如果 name 没有指定后缀，则加上，否则就返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	//f := AddUpper()
	//fmt.Println(f(1))
	//fmt.Println(f(2))
	//fmt.Println(f(3))

	//测试makeSuffix
	f := makeSuffix(".jpg")
	fmt.Println("文件名处理后=", f("winter"))
	fmt.Println("文件名处理后=", f("bird.jpg"))
}
