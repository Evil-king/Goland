package main

import "fmt"

func fpnx(ch <-chan int, quit <-chan bool) {
	for {
		select {
		case num := <-ch:
			fmt.Print(num, " ")
		case <-quit:
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go fpnx(ch, quit)

	x, y := 1, 1
	for i := 0; i < 20; i++ {
		//1,1,2,3,5
		ch <- x
		x, y = y, x+y
	}

	quit <- true
}
