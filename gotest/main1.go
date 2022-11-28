package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//有三个函数 分别打印cat dog fish
//要求每个函数都起一个goroutine，请按照cat dog fish的顺序打印在屏幕上，每个100次
// func main() {
// 	var wg sync.WaitGroup
// 	var dogCount uint64
// 	var catCount uint64
// 	var fishCount uint64

// 	dogch := make(chan struct{}, 1)
// 	catch := make(chan struct{}, 1)
// 	fishch := make(chan struct{}, 1)

// 	wg.Add(3)
// 	go cat(&wg, catCount, catch, dogch)
// 	go dog(&wg, dogCount, dogch, fishch)
// 	go fish(&wg, fishCount, fishch, catch)

// 	catch <- struct{}{}

// 	wg.Wait()
// }

func main() {
	fmt.Println(time.Now().AddDate(0, 0, 7).Format("2006-01-02"))

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
