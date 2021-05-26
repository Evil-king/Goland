package main

import (
	"fmt"
	"time"
)

func main() {
	//output1 := make(chan string)
	//output2 := make(chan string)
	//
	//go test1(output1)
	//go test2(output2)
	//
	//select {
	//case s1:=<-output1:
	//	fmt.Println("s1=",s1)
	//case s2:=<-output2:
	//	fmt.Println("s2=", s2)
	//}
	//fmt.Println("main结束")

	//创建管道
	output1 := make(chan string, 10)
	go write(output1)

	for s := range output1 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}

func write(output1 chan string) {
	for {
		select {
		//写数据
		case output1 <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}

//func test2(ch chan string) {
//	//time.Sleep(time.Second * 2)
//	ch <- "test2"
//}
//
//func test1(ch chan string) {
//	//time.Sleep(time.Second * 5)
//	ch <- "test1"
//}
