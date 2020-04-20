package main

import "fmt"

func main() {
	//演示一下管道的使用
	//创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int,3)

	//看看intChan是什么
	fmt.Printf("intChan 的值=%v \n",intChan)

	//向管道写入数据
	intChan<- 10
	num := 211
	intChan<- num
	intChan<- 50


	//注意点当我们给管道写入数据时，不能超过容量

	//输出管道的长度和cap(容量)
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))

	//从管道中读取数据
	var num2 int
	num2 = <- intChan
	fmt.Println("num2=",num2)
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))
}
