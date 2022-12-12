package main

import "fmt"

func main() {
	defer_call()

	fmt.Println("main 正常结束")
}

// defer遇见panic，但是并不捕获异常的情况
func defer_call() {
	defer func() { fmt.Println("defer: panic 之前1") }()
	defer func() { fmt.Println("defer: panic 之前2") }()

	panic("异常内容") //触发defer出栈

	defer func() { fmt.Println("defer: panic 之后，永远执行不到") }()
}
