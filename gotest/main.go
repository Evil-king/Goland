package main

import (
	"fmt"
	"github.com/spf13/cast"
	"sync"
	"sync/atomic"
	"time"
)

// 有三个函数 分别打印cat dog fish
// 要求每个函数都起一个goroutine，请按照cat dog fish的顺序打印在屏幕上，每个100次

func main() {
	//referenceList := []string{"a", "b", "c"}
	//
	//flagChannel := make(chan string, len(referenceList))
	//resultChannel := make(chan bool)
	//
	//ticker := time.NewTicker(5 * time.Second)
	//
	//// 启动一个 goroutine 来执行 BatchCartTransStateController
	//go func() {
	//	for range ticker.C {
	//		BatchCartTransStateController(flagChannel, resultChannel, referenceList)
	//	}
	//}()
	//flag := <-resultChannel
	//if flag {
	//	ticker.Stop()
	//}
	//fmt.Println("通道中剩余的数据量:", len(flagChannel))
	//PrintOddAndEven()
	//PrintNGoroutine()
}

func BatchCartTransStateController(flagChannel chan string, resultChannel chan bool, referenceList []string) {
	// 执行 BatchCartTransStateController 的逻辑
	for key, str := range referenceList {
		fmt.Println(str)
		flagChannel <- cast.ToString(key)
	}
	//fmt.Println("通道中剩余的数据量:", len(flagChannel))
	//if len(flagChannel) == len(referenceList) {
	//	resultChannel <- "stop"
	//}
	resultChannel <- true

}

func doSomething() {
	fmt.Println("执行第一个方法")
}

func doSomethingElse() {
	fmt.Println("执行第二个方法")
}

func GetWeek(datetime string) (y, w int) {
	timeLayout := "2006-01-02"
	loc, _ := time.LoadLocation("Local")
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	return tmp.ISOWeek()
}

func cat(wg *sync.WaitGroup, count uint64, catch, doghch chan struct{}) {
	for {
		if count >= uint64(100) {
			wg.Done()
		}
		<-catch
		fmt.Println("cat")
		atomic.AddUint64(&count, 1)
		doghch <- struct{}{}
	}
}

func dog(wg *sync.WaitGroup, count uint64, dogch, fishch chan struct{}) {
	for {
		if count >= uint64(100) {
			wg.Done()
		}
		<-dogch
		fmt.Println("dog")
		atomic.AddUint64(&count, 1)
		fishch <- struct{}{}
	}

}

func fish(wg *sync.WaitGroup, count uint64, fishch, catch chan struct{}) {
	for {
		if count >= uint64(100) {
			wg.Done()
		}
		<-fishch
		fmt.Println("fish")
		atomic.AddUint64(&count, 1)
		catch <- struct{}{}
	}
}

// 打印奇数和偶数
func PrintOddAndEven() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	oddTurn := true

	wg.Add(2)
	// 奇数
	printOdd := func() {
		defer wg.Done()
		for i := 1; i <= 100; i += 2 {
			mu.Lock()
			if !oddTurn {
				cond.Wait()
			}
			fmt.Println(i)
			oddTurn = false
			cond.Signal()
			mu.Unlock()
		}
	}

	// 偶数
	printEven := func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			mu.Lock()
			if oddTurn {
				cond.Wait()
			}
			fmt.Println(i)
			oddTurn = true
			cond.Signal()
			mu.Unlock()
		}
	}
	go printEven()
	go printOdd()
	wg.Wait()
}

// n 个 goroutine 顺序输出 1-100
func PrintNGoroutine() {
	n := 5 //假如有5个
	count := 100
	var wg sync.WaitGroup
	wg.Add(n)

	// 创建 n 个 channel 用于同步控制
	channels := make([]chan struct{}, n)
	for i := range channels {
		channels[i] = make(chan struct{})
	}
	// 启动 n 个 goroutine
	for i := 0; i < n; i++ {
		go func(id int) {
			defer wg.Done()
			// 控制打印1-100
			for j := id; j < count; j += n {
				<-channels[id] // 等待当前channel被打开
				fmt.Println(j + 1)
				// 打开下一个 channel
				next := (id + 1) % n
				channels[next] <- struct{}{}
			}
		}(i)
	}
	// 初始启动第一个 channel
	channels[0] <- struct{}{}

	// 等待所有 goroutine 完成
	wg.Wait()
}
