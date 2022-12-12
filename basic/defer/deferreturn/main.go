package main

import "fmt"

//return之后的语句先执行，defer后的语句后执行
func deferFunc() int {
	fmt.Println("defer func called")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called")
	return 0
}

func returnAndDefer() int {

	defer deferFunc()

	return returnFunc()
}

func main() {
	returnAndDefer()
}
