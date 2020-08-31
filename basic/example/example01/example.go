package main

import "fmt"

func main() {
	var num = 9
	fmt.Printf("num address=%v\n", num)
	var ptr *int
	ptr = &num
	*ptr = 10
	fmt.Printf("num=%v", num)
}
