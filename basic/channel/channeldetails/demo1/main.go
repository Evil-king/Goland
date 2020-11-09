package main

import "fmt"

func main() {
	//管道可以声明为只读或只写

	//1、在默认情况下，管道是双向
	//var chan1 chan int //可读可写

	//2、声明为只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	//num := <-chan2  error
	fmt.Println("chan2=", chan2)

	//3、声明为只读
	var chan3 <-chan int
	num2 := <-chan3
	//chan3<- 30 error
	fmt.Println("num2=", num2)
}
