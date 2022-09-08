package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// WaitForResult 等待结果模式
func WaitForResult() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "data" //放入数据到channel中
		fmt.Println("child : sent signal")
	}()
	d := <-ch //从channel中取出数据
	fmt.Println("parent : recv'd signal :", d)
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// FanOut 扇出/扇入模式
func FanOut() {
	children := 2000
	ch := make(chan string, children) //带缓冲的channel
	for c := 0; c < children; c++ {
		go func(child int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "data" //放入数据到channel中
			fmt.Println("child : sent signal :", child)
		}(c)
	}

	for children > 0 {
		d := <-ch
		children--
		fmt.Println(d)
		fmt.Println("parent : recv'd signal :", children)
	}
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// WaitForTask 等待任务模式
func WaitForTask() {
	ch := make(chan string)
	go func() {
		d := <-ch //放入数据到channel中
		fmt.Println("child : sent signal :", d)
	}()
	ch <- "data"
	fmt.Println("parent : sent signal")
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// Pooling Goroutine池
func Pooling() {
	ch := make(chan string)
	g := runtime.GOMAXPROCS(0)
	for c := 0; c < g; c++ {
		go func(child int) {
			for d := range ch {
				fmt.Printf("child %d : recv'd signal : %s\n", child, d)
			}
			fmt.Printf("child %d : recv'd shutdown signal\n", child)
		}(c)
	}
	const work = 100
	for w := 0; w < work; w++ {
		ch <- "data"
		fmt.Println("parent : sent signal :", w)
	}
	close(ch)
	fmt.Println("parent : sent shutdown signal")
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")

}

// Drop 模式
func Drop() {
	const cap = 100
	ch := make(chan string, cap)
	go func() {
		for p := range ch {
			fmt.Println("child : recv'd signal :", p)
		}
	}()
	const work = 2000
	for w := 0; w < work; w++ {
		select {
		case ch <- "data":
			fmt.Println("parent : sent signal :", w)
		default:
			fmt.Println("parent : dropped data :", w)
		}
	}
	close(ch)
	fmt.Println("parent : sent shutdown signal")
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// Cancellation 取消模式
func Cancellation() {
	duration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	ch := make(chan string, 1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "data"
	}()
	select {
	case d := <-ch:
		fmt.Println("work complete", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// RetryTimeout 重试超时模式
func RetryTimeout(ctx context.Context, retryInterval time.Duration,
	check func(ctx context.Context) error) {
	for {
		fmt.Println("perform user check call")
		if err := check(ctx); err == nil {
			fmt.Println("work finished successfully")
			return
		}
		fmt.Println("check if timeout has expired")
		if ctx.Err() != nil {
			fmt.Println("time expired 1 :", ctx.Err())
			return
		}
		fmt.Printf("wait %s before trying again\n", retryInterval)
		t := time.NewTimer(retryInterval)
		select {
		case <-ctx.Done():
			fmt.Println("timed expired 2 :", ctx.Err())
			t.Stop()
			return
		case <-t.C:
			fmt.Println("retry again")
		}
	}
}

func main() {
	//RetryTimeout()
}
