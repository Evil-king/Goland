package main

import (
	"fmt"
	"time"
)

//向 intChan放入1-8000个数
func punNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan<- i
	}
	//关闭intChan
	close(intChan)
}

//开启4个协程 从intChan取出数据 并判断是否是素数，如果是就放入到primeChan
func putPrimNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	//使用for循环
	var flag bool
	for {
		time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		if !ok { //intChan 取不到数据
			break
		}
		num = <-intChan
		flag = true
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num % i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}
		if flag {
			//将这个数放入到primeChan
			primeChan<- num
		}
	}
	//这里我们还不能关闭 primeChan 向退出的管道exitChan 写入true
	fmt.Println("有一个primeNum 协程因为取不到数据，退出")
	//向退出的管道exitChan 写入true
	exitChan<- true
}

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 8000) //放入结果
	//标识退出的管道
	exitChan := make(chan bool, 4) //4个

	//开启一个协程 向 intChan 放入 1-8000
	go punNum(intChan)
	//开启4个协程 从intChan取出数据 并判断是否是素数，如果是就放入到primeChan
	for i := 0; i < 4; i++ {
		go putPrimNum(intChan, primeChan, exitChan)
	}
	//这里我们主线程 进行处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		//当我们从exitChan 取出了四个结果 就可以放心的关闭primeChan
		close(primeChan)
	}()
	//遍历我们的primeChan
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		//输出结果
		fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("主线程退出")

}
