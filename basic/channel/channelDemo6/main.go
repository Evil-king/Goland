package main

import "fmt"

// channel 上的发送操作总在对应的接收操作完成前发生。
// 从无缓冲的 channel 中进行的接收，要发生在对该 channel 进行的发送完成前。
// 如果是无缓冲的 channel， write 还没写入结束，read 就已经开始接收了，所以可以保证 read 执行，
// 但是反过来如果有缓冲，那么 read 可能还没开始 write 就结束了，所以就有可能什么都不输出就结束了

// 对于带缓冲的Channel，对于 Channel 的第 K 个接收完成操作发生在第 K+C 个发送操作完成之前，(接收操作 在 发送操作之前)
// 其中 C 是 Channel 的缓存大小。 如果将 C 设置为 0 自然就对应无缓存的 Channel，
// 也即使第 K 个接收完成在第 K 个发送完成之前

func main() {
	done := make(chan int, 1)
	go func() {
		fmt.Println("你好, 世界")
		done <- 1 //发送
	}()
	review := <-done //接收
	fmt.Println(review)
}
