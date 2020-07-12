package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) //close
	//这时不能够写入数据到channel
	//intChan<- 300
	//fmt.Println("okook~")
	//当管道关闭后，读取数据是可以的
	n1 := <-intChan
	fmt.Println("n1=", n1)

	//遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 //放入100个数据到管道
	}
	//遍历管道不能使用普通的for循环
	//for i :=0;i<len(intChan2);i++{
	//
	//}

	//如果遍历时 不关闭管道 数据会全部取出 但是会出现死锁 fatal error: all goroutines are asleep - deadlock!
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}
