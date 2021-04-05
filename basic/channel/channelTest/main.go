package main

import (
	"fmt"
	"time"
)

//定义channel
var chennel = make(chan int)

func printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(300 * time.Millisecond)
	}
}

func person1() {
	printer("hello")
	chennel <- 8 //写入channel中
}

func person2() {
	<-chennel  //读出channel
	printer("world")
}

func main() {
	go person1()
	go person2()
	for {

	}
}
