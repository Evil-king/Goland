package main

import (
	"fmt"
	"github.com/spf13/cast"
	"math/rand"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

// 有三个函数 分别打印cat dog fish
// 要求每个函数都起一个goroutine，请按照cat dog fish的顺序打印在屏幕上，每个100次

type RechargeListExcel struct {
	ID   int    `json:"id" excel:"userId"`
	Name string `json:"name" excel:"nickname"`
}

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

	//cards := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println("Before shuffle:", cards)
	//shuffle(cards)
	//fmt.Println("After shuffle:", cards)
	//removeElementByIndex(cards, 3)
	//for _, v := range cards {
	//	fmt.Println(v)
	//}]

	// Example slice of pointers to structs
	var slice []RechargeListExcel
	slice = append(slice, RechargeListExcel{
		ID:   1,
		Name: "test",
	}, RechargeListExcel{
		ID:   2,
		Name: "fox",
	})
	// Get the type of the slice
	t := reflect.TypeOf(slice)           // 获取类型
	sliceValue := reflect.ValueOf(slice) // 获取值

	// Check if it's a slice
	if t.Kind() == reflect.Slice {
		// Get the type of the elements in the slice
		elemType := t.Elem()

		for i := 0; i < sliceValue.Len(); i++ {
			elemValue := sliceValue.Index(i) // get slice elemValue
			// Check if the element is a pointer
			if elemType.Kind() == reflect.Ptr {
				// Dereference the pointer to get the underlying struct type
				elemType = elemType.Elem()
			}

			// Check if the element is a struct
			if elemType.Kind() == reflect.Struct {
				// Now you can call NumField() on the struct type
				fmt.Printf("The struct has %d fields\n", elemType.NumField())

				// Print all field names
				for j := 0; j < elemType.NumField(); j++ {
					field := elemType.Field(j)
					tag := field.Tag.Get("excel")
					value := elemValue.Field(j) // 获取字段值
					fmt.Printf("Field %d: %s (%s) tag:%s,value:%s\n", j, field.Name, field.Type, tag, value)
				}
			} else {
				fmt.Println("The slice does not contain structs")
			}
		}
	} else {
		fmt.Println("The provided value is not a slice")
	}
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
		for i := 1; i <= 26; i += 2 {
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
		for i := 2; i <= 26; i += 2 {
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

// 洗牌算法
func shuffle(slice []int) {

	rand.Seed(time.Now().UnixNano())
	n := len(slice)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func solution(ranks map[int]int) int {
	tempMap := make(map[int]int, 0)
	for k, v := range ranks {
		// 如果用户的排名与荣誉值相等
		if k != v {
			// 排名最接近荣誉值:排名>=荣誉值
			if k >= v {
				tempMap[k] = v
			}
		} else {
			return k
		}
	}
	// 如果tempMap不为空,寻找排名最低
	if len(tempMap) > 0 {
		tempNum := 0
		flag := true
		for k := range tempMap {
			if flag {
				tempNum = k
				flag = false
			}
			if k < tempNum {
				tempNum = k
			}
		}
		return tempNum
	}
	return 0
}

func removeElementByIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
