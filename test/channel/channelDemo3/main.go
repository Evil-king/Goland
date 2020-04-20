package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) //关闭channel 就不能在往channel里面新增数据了  但是仍然可以读取数据
	n1 := <-intChan
	fmt.Println("n1=", n1)

	//遍历channel 遍历channel不能使用普通的for循环
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 //放入100个数据到channel中
	}

	close(intChan2)

	//在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历
	for v := range intChan2{
		fmt.Println("v=", v)
	}

}
